package compute

import (
	"time"
)

// Define Compute Provider interface

type ComputeProvider interface {
	// List all instance
	ListNodes() ([]*Node, error)

	// Detail instance by Id
	DetailNode(id string) (*Node, error)

	// Create new instance
	CreateNode(opts *NodeCreateOpts) (*Node, error)

	// Destroy an existing instance
	DestroyNode(node *Node) error

	// Reboot instance
	RebootNode(node *Node) error

	// Start instance
	StartNode(node *Node) error

	// Stop instance
	StopNode(node *Node) error

	// Get the current state of instance
	GetNodeState(node *Node) (NodeState, error)

	// Get the Console url for instance
	GetNodeConsole(node *Node) (string, error)

	// Get the public IP address of instance
	GetNodePublicIp(node *Node) (string, error)

	// Get the private IP address of instance
	GetNodePrivateIp(node *Node) (string, error)

	// List all available storage volumes for instance
	ListVolumes(node *Node) ([]*StorageVolume, error)

	// Attach volume to instance
	AttachVolume(node *Node, snapshot *StorageVolume) error

	// Detach volume from instance
	DetachVolume(node *Node, snapshot *StorageVolume) error

	// List all snapshots for instance
	ListSnapshots(node *Node) ([]*VolumeSnapshot, error)

	// Create snapshot for instance
	CreateSnapshot(node *Node, name string) (*VolumeSnapshot, error)

	// Destroy snapshot for instance
	DestroySnapshot(node *Node, snapshot *VolumeSnapshot) error

	// Apply snapshot to instance
	ApplySnapshot(node *Node, snapshot *VolumeSnapshot) error

	// List all available images for instance
	ListImages() ([]*NodeImage, error)

	// Apply Image to instance
	ApplyImage(node *Node, image *NodeImage) error

	// List all available sizes for instance
	ListSizes() ([]*NodeSize, error)

	// Resize instance
	ResizeNode(node *Node, opts *NodeResizeOpts) error

	// List all available locations for instance
	ListLocations() ([]*Location, error)
}

// compute instance

type Node struct {
	Id        string
	Name      string
	State     NodeState
	Size      *NodeSize
	Image     *NodeImage
	PublicIp  string
	PrivateIp string
	Location  *Location
	CreatedAt time.Time
	Extra     map[string]interface{}
}

// options for creating new compute

type NodeCreateOpts struct {
	Name     string
	Size     *NodeSize
	Image    *NodeImage
	Location *Location
	Extra    map[string]interface{}
}

// options for resizing compute

type NodeResizeOpts struct {
	Size  *NodeSize
	Extra map[string]interface{}
}

// compute snapshot

type StorageVolume struct {
	Id    string
	Name  string
	Type  string
	Size  int
	State StorageVolumeState
}

// compute snapshot

type VolumeSnapshot struct {
	Id        string
	Name      string
	Size      int
	State     StorageVolumeState
	CreatedAt time.Time
}

// compute image

type NodeImage struct {
	Id     string
	Name   string
	OSType OSType
	State  NodeImageState
	Extra  map[string]interface{}
}

// compute size

type NodeSize struct {
	Id           string
	Name         string
	Architecture Architecture
	Gpu          int
	Cpu          int
	Ram          int
	Disk         int
	Bandwidth    int
	Price        float64
	Extra        map[string]interface{}
}

// compute location

type Location struct {
	Id      string
	Name    string
	Country string
	Extra   map[string]interface{}
}
