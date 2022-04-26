package webdav

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type driver struct {
	csiDriver  *csicommon.CSIDriver
	endpoint   string
	configFile string

	//ids *identityServer
	ns    *nodeServer
	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	driverName          = "csi-webdav"
	parameterURL        = "url"
	parameterDir        = "dir"
	parameterUser       = "user"
	parameterPassword   = "password"
	parameterConfigName = "configName"
)

var (
	Version   = "latest"
	BuildTime = "1970-01-01 00:00:00"
)

func NewDriver(nodeID, endpoint string, configFile string) *driver {
	glog.Infof("Starting new %s driver in version %s built %s", driverName, Version, BuildTime)

	d := &driver{}

	d.endpoint = endpoint
	d.configFile = configFile

	csiDriver := csicommon.NewCSIDriver(driverName, Version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	// SSHFS plugin does not support ControllerServiceCapability now.
	// If support is added, it should set to appropriate
	// ControllerServiceCapability RPC types.
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{csi.ControllerServiceCapability_RPC_UNKNOWN})

	d.csiDriver = csiDriver

	return d
}

func GetK8SClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func NewNodeServer(d *driver, k8sClient *kubernetes.Clientset) *nodeServer {
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.csiDriver),
		configFile:        d.configFile,
		socketPath:        "",
		K8sClient:         k8sClient,
	}
}

func NewControllerServer(d *driver, k8sClient *kubernetes.Clientset) *WebdavControllerServer {
	return &WebdavControllerServer{
		Driver:    d,
		K8sClient: k8sClient,
	}
}

func (d *driver) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	k8sClient := GetK8SClient()
	d.ns = NewNodeServer(d, k8sClient)
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.csiDriver),
		// SSHFS plugin has not implemented ControllerServer
		// using default controllerserver.
		NewControllerServer(d, k8sClient),
		d.ns)
	s.Wait()
}
