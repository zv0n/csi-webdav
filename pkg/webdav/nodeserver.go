package webdav

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/tidwall/gjson"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/util/mount"

	"github.com/zv0n/csi-webdav/webdavrpc"

	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	K8sClient  *kubernetes.Clientset
	configFile string
	socketPath string
}

func getSocketPath(configPath string) string {
	defaultPath := "/var/lib/webdav/webdav-proxy.sock"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("Could not read file \""+configPath+"\" - %v", err)
		return defaultPath
	}
	return gjson.Get(string(data), "socket").Str
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	notMnt, e := mount.New("").IsLikelyNotMountPoint(targetPath)
	if e != nil {
		if os.IsNotExist(e) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, e.Error())
		}
	}

	if !notMnt {
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if e := validateVolumeContext(req); e != nil {
		return nil, e
	}

	url := req.GetVolumeContext()[parameterURL]
	dir := req.GetVolumeContext()[parameterDir]
	user := req.GetVolumeContext()[parameterUser]
	password := req.GetVolumeContext()[parameterPassword]
	configname := req.GetVolumeContext()[parameterConfigName]
	podName := req.GetVolumeContext()["csi.storage.k8s.io/pod.name"]
	podNamespace := req.GetVolumeContext()["csi.storage.k8s.io/pod.namespace"]

	uid, gid, err := ns.getPodUIDGID(podName, podNamespace)

	if err != nil {
		return nil, fmt.Errorf("Could not get pod information: %v", err)
	}

	if ns.socketPath == "" {
		ns.socketPath = fmt.Sprintf("unix://%s", getSocketPath(ns.configFile))
	}
	conn, err := grpc.Dial(ns.socketPath, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rpcClient := webdavrpc.NewMountServiceClient(conn)
	input := webdavrpc.MountWebdavRequest{
		Url:        url,
		Dir:        dir,
		User:       user,
		Password:   password,
		ConfigName: configname,
		Uid:        strconv.FormatInt(uid, 10),
		Gid:        strconv.FormatInt(gid, 10),
		Target:     targetPath,
	}

	response, err := rpcClient.MountWebdav(context.Background(), &input)
	if err != nil {
		return nil, err
	}

	if response.Output != "Success" {
		return nil, status.Error(codes.Internal, response.Output)
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	id := req.VolumeId
	pvname := strings.Split(id, ":")[0]

	if ns.socketPath == "" {
		ns.socketPath = fmt.Sprintf("unix://%s", getSocketPath(ns.configFile))
	}
	conn, err := grpc.Dial(ns.socketPath, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rpcClient := webdavrpc.NewMountServiceClient(conn)

	input := webdavrpc.UmountWebdavRequest{
		MountTarget: targetPath,
		ConfigName:  pvname,
	}

	response, err := rpcClient.UmountWebdav(context.Background(), &input)
	if err != nil {
		return nil, err
	}

	if response.Output != "Success" {
		return nil, status.Error(codes.Internal, response.Output)
	}

	glog.Infof("successfully unmount volume: %s", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func validateVolumeContext(req *csi.NodePublishVolumeRequest) error {
	if _, ok := req.GetVolumeContext()[parameterURL]; !ok {
		return status.Errorf(codes.InvalidArgument, "missing volume context value: "+parameterURL)
	}
	if _, ok := req.GetVolumeContext()[parameterDir]; !ok {
		return status.Errorf(codes.InvalidArgument, "missing volume context value: "+parameterDir)
	}
	if _, ok := req.GetVolumeContext()[parameterUser]; !ok {
		return status.Errorf(codes.InvalidArgument, "missing volume context value: "+parameterUser)
	}
	if _, ok := req.GetVolumeContext()[parameterPassword]; !ok {
		return status.Errorf(codes.InvalidArgument, "missing volume context value: "+parameterPassword)
	}
	if _, ok := req.GetVolumeContext()[parameterConfigName]; !ok {
		return status.Errorf(codes.InvalidArgument, "missing volume context value: "+parameterConfigName)
	}
	return nil
}

func (ns *nodeServer) getPodUIDGID(podName string, namespace string) (int64, int64, error) {
	pod, err := ns.K8sClient.CoreV1().Pods(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		return -1, -1, err
	}
	var uid int64 = 0
	var gid int64 = 0
	if pod.Spec.SecurityContext.RunAsUser != nil {
		uid = *pod.Spec.SecurityContext.RunAsUser
	}
	if pod.Spec.SecurityContext.RunAsGroup != nil {
		gid = *pod.Spec.SecurityContext.RunAsGroup
	}
	return uid, gid, nil
}
