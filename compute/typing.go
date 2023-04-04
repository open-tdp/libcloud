package compute

import (
	"context"
)

// Define a Compute Provider interface that matches the one in Apache Libcloud

type Provider interface {
	// List all nodes/servers
	ListNodes(ctx context.Context) ([]*Node, error)

	// Get a node/server by its ID
	GetNode(ctx context.Context, id string) (*Node, error)

	// Create a new node/server instance
	CreateNode(ctx context.Context, opts *NodeCreateOpts) (*Node, error)

	// Destroy an existing node/server instance
	DestroyNode(ctx context.Context, node *Node) error

	// Reboot a node/server instance
	RebootNode(ctx context.Context, node *Node) error

	// Start a node/server instance
	StartNode(ctx context.Context, node *Node) error

	// Stop a node/server instance
	StopNode(ctx context.Context, node *Node) error

	// Resize a node/server instance
	ResizeNode(ctx context.Context, node *Node, opts *NodeResizeOpts) error

	// List all available sizes for a node/server instance
	ListSizes(ctx context.Context, node *Node) ([]*Size, error)

	// List all available images for a node/server instance
	ListImages(ctx context.Context, node *Node) ([]*Image, error)

	// Get the current status of a node/server instance
	GetNodeStatus(ctx context.Context, node *Node) (Status, error)

	// Get the public IP address of a node/server instance
	GetPublicIPAddress(ctx context.Context, node *Node) (string, error)

	// Get the private IP address of a node/server instance
	GetPrivateIPAddress(ctx context.Context, node *Node) (string, error)

	// Get the SSH connection information for a node/server instance
	GetSSHConnection(ctx context.Context, node *Node) (*SSHConnection, error)
}

// Node represents a compute node/server instance

type Node struct {
	ID               string
	Name             string
	Size             *Size
	Image            *Image
	Location         *Location
	PublicIPAddress  string
	PrivateIPAddress string
	Extra            map[string]interface{}
}

// NodeCreateOpts represents options for creating a new node/server instance

type NodeCreateOpts struct {
	Name     string
	Size     *Size
	Image    *Image
	Location *Location
	Extra    map[string]interface{}
}

// NodeResizeOpts represents options for resizing a node/server instance

type NodeResizeOpts struct {
	Size  *Size
	Extra map[string]interface{}
}

// Size represents a compute node/server size

type Size struct {
	ID        string
	Name      string
	RAM       int
	Disk      int
	Bandwidth int
	Price     float64
	Extra     map[string]interface{}
}

// Image represents a compute node/server image

type Image struct {
	ID       string
	Name     string
	Location *Location
	Extra    map[string]interface{}
}

// Location represents a compute node/server location

type Location struct {
	ID      string
	Name    string
	Country string
	Extra   map[string]interface{}
}

// SSHConnection represents SSH connection information for a node/server instance

type SSHConnection struct {
	Hostname string
	Port     int
	Username string
	Password string
	Key      string
}

// Status represents a compute node/server status

type Status string

const (
	StatusRunning Status = "running"
	StatusStopped Status = "stopped"
	StatusError   Status = "error"
)
