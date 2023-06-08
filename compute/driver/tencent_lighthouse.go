package drivers

import (
	"github.com/open-tdp/go-libcloud/compute"
	"github.com/open-tdp/go-libcloud/provider"
	"github.com/open-tdp/go-libcloud/provider/tencent"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

type TencentLighthouseDriver struct {
	client     *tencent.Client
	lighthouse *lighthouse.Client
	rq         *provider.ReqeustParam
}

func NewTencentLighthouseDriver(rq *provider.ReqeustParam) *TencentLighthouseDriver {

	client := tencent.NewClient(rq)
	lighthouse, _ := client.Lighthouse()

	return &TencentLighthouseDriver{client, lighthouse, rq}

}

// List all instance
func (p *TencentLighthouseDriver) ListNodes() ([]*compute.Node, error) {
	return nil, nil
}

// Detail instance by Id
func (p *TencentLighthouseDriver) DetailNode(id string) (*compute.Node, error) {
	return nil, nil
}

// Create new instance
func (p *TencentLighthouseDriver) CreateNode(opts *compute.NodeCreateOpts) (*compute.Node, error) {
	return nil, nil
}

// Destroy an existing instance
func (p *TencentLighthouseDriver) DestroyNode(node *compute.Node) error {
	return nil
}

// Reboot instance
func (p *TencentLighthouseDriver) RebootNode(node *compute.Node) error {
	return nil
}

// Start instance
func (p *TencentLighthouseDriver) StartNode(node *compute.Node) error {
	return nil
}

// Stop instance
func (p *TencentLighthouseDriver) StopNode(node *compute.Node) error {
	return nil
}

// Get the current state of instance
func (p *TencentLighthouseDriver) GetNodeState(node *compute.Node) (compute.NodeState, error) {
	return compute.NodeState(""), nil
}

// Get the Console url for instance
func (p *TencentLighthouseDriver) GetNodeConsole(node *compute.Node) (string, error) {
	return "", nil
}

// Get the public IP address of instance
func (p *TencentLighthouseDriver) GetNodePublicIp(node *compute.Node) (string, error) {
	return "", nil
}

// Get the private IP address of instance
func (p *TencentLighthouseDriver) GetNodePrivateIp(node *compute.Node) (string, error) {
	return "", nil
}

// List all available storage volumes for instance
func (p *TencentLighthouseDriver) ListVolumes(node *compute.Node) ([]*compute.StorageVolume, error) {
	return nil, nil
}

// Attach volume to instance
func (p *TencentLighthouseDriver) AttachVolume(node *compute.Node, volume *compute.StorageVolume) error {
	return nil
}

// Detach volume from instance
func (p *TencentLighthouseDriver) DetachVolume(node *compute.Node, volume *compute.StorageVolume) error {
	return nil
}

// List all snapshots for instance
func (p *TencentLighthouseDriver) ListSnapshots(node *compute.Node) ([]*compute.VolumeSnapshot, error) {
	return nil, nil
}

// Create snapshot for instance
func (p *TencentLighthouseDriver) CreateSnapshot(node *compute.Node, name string) (*compute.VolumeSnapshot, error) {
	return nil, nil
}

// Destroy snapshot for instance
func (p *TencentLighthouseDriver) DestroySnapshot(node *compute.Node, snapshot *compute.VolumeSnapshot) error {
	return nil
}

// Apply snapshot to instance
func (p *TencentLighthouseDriver) ApplySnapshot(node *compute.Node, snapshot *compute.VolumeSnapshot) error {
	return nil
}

// List all available images for instance
func (p *TencentLighthouseDriver) ListImages() ([]*compute.NodeImage, error) {
	return nil, nil
}

// Apply Image to instance
func (p *TencentLighthouseDriver) ApplyImage(node *compute.Node, image *compute.NodeImage) error {
	return nil
}

// List all available sizes for instance
func (p *TencentLighthouseDriver) ListSizes() ([]*compute.NodeSize, error) {
	return nil, nil
}

// Resize instance
func (p *TencentLighthouseDriver) ResizeNode(node *compute.Node, opts *compute.NodeResizeOpts) error {
	return nil
}

// List all available locations for instance
func (p *TencentLighthouseDriver) ListLocations() ([]*compute.Location, error) {
	return nil, nil
}
