package webdav

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type WebdavControllerServer struct {
	Driver    *driver
	K8sClient *kubernetes.Clientset
}

type WebdavVolume struct {
}

func (cs *WebdavControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	name := req.GetName()
	if len(name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "CreateVolume name must be provided")
	}

	parameters := req.GetParameters()
	if parameters == nil {
		parameters = make(map[string]string)
	}

	pvname, ok := parameters["csi.storage.k8s.io/pv/name"]
	if !ok {
		return nil, status.Error(codes.NotFound, "The PV does not have a name, this is weird and indicative of a larger issue")
	}
	pvcname, ok := parameters["csi.storage.k8s.io/pvc/name"]
	if !ok {
		return nil, status.Error(codes.NotFound, "The PVC does not have a name, this is weird and indicative of a larger issue")
	}
	namespace, ok := parameters["csi.storage.k8s.io/pvc/namespace"]
	if !ok {
		return nil, status.Error(codes.NotFound, "The PVC does not have a namespace, this is weird and indicative of a larger issue")
	}

	publishParams := make(map[string]string)
	var err error = nil
	publishParams[parameterURL], err = cs.GetPVCURL(pvcname, namespace)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "The PVC does not have a 'url' annotation")
	}
	publishParams[parameterDir], err = cs.GetPVCDir(pvcname, namespace)
	if err != nil {
		publishParams[parameterDir] = ""
	}
	publishParams[parameterUser], err = cs.GetPVCUser(pvcname, namespace)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "The PVC does not have a 'user' annotation")
	}
	// TODO password secret instead of plaintext
	password, err := cs.GetPVCPassword(pvcname, namespace)
	if err != nil {
		password, err = cs.GetPVCPasswordSecret(pvcname, namespace)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "The PVC does not have a valid 'password' or 'passwordSecret' annotation: "+err.Error())
		}
	}
	publishParams[parameterPassword] = password
	publishParams[parameterConfigName] = pvname

	id := pvname + ":" + namespace

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      id,
			CapacityBytes: 0, // by setting it to zero, Provisioner will use PVC requested size as PV size
			VolumeContext: publishParams,
		},
	}, nil
}

// DeleteVolume delete a volume
func (cs *WebdavControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// We always retain all files, so just return delete response
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *WebdavControllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if req.GetVolumeCapabilities() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities missing in request")
	}

	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.GetVolumeCapabilities(),
		},
		Message: "",
	}, nil
}

func (cs *WebdavControllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerGetCapabilities implements the default GRPC callout.
// Default supports all capabilities
func (cs *WebdavControllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: []*csi.ControllerServiceCapability{
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (cs *WebdavControllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *WebdavControllerServer) validateVolumeCapabilities(caps []*csi.VolumeCapability) error {
	return nil
}

func (cs *WebdavControllerServer) validateVolumeCapability(c *csi.VolumeCapability) error {
	return nil
}

func (cs *WebdavControllerServer) GetPVCURL(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations[parameterURL]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "PVC does not have URL")
}

func (cs *WebdavControllerServer) GetPVCDir(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations[parameterDir]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "PVC does not have dir")
}

func (cs *WebdavControllerServer) GetPVCUser(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations[parameterUser]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "PVC does not have user")
}

func (cs *WebdavControllerServer) GetPVCPassword(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations[parameterPassword]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "PVC does not have password")
}

func (cs *WebdavControllerServer) GetPVCPasswordSecret(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations[parameterPasswordSecret]; ok {
		secret, err := cs.K8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), val, v1.GetOptions{})
		if err != nil {
			return "", status.Error(codes.NotFound, "Secret could not be found: "+err.Error())
		}
		if password, ok := secret.Data["password"]; ok {
			return string(password), nil
		}
		return "", status.Error(codes.InvalidArgument, "Secret does not have 'password' entry")
	}

	return "", status.Error(codes.NotFound, "PVC does not have passwordSecret: "+err.Error())
}
