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
	publishParams[parameterPassword], err = cs.GetPVCPassword(pvcname, namespace)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "The PVC does not have a 'password' annotation")
	}

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
	if val, ok := ns.Annotations["url"]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "Given namespace does not belong to any project")
}

func (cs *WebdavControllerServer) GetPVCDir(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations["dir"]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "Given namespace does not belong to any project")
}

func (cs *WebdavControllerServer) GetPVCUser(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations["user"]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "Given namespace does not belong to any project")
}

func (cs *WebdavControllerServer) GetPVCPassword(pvc string, namespace string) (string, error) {
	ns, err := cs.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvc, v1.GetOptions{})
	if err != nil {
		return "", err
	}
	if val, ok := ns.Annotations["password"]; ok {
		return val, nil
	}

	return "", status.Error(codes.NotFound, "Given namespace does not belong to any project")
}
