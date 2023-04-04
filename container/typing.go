package container

import (
	"time"
)

// Define the Container interface

type ContainerProvider interface {
	// List all containers
	ListContainers() ([]ContainerInfo, error)

	// Get a container
	DetailContainer(name string) (ContainerInfo, error)

	// Create a container
	CreateContainer(name string) error

	// Delete a container
	DestroyContainer(name string) error

	// Start a container
	StartContainer(name string) error

	// Stop a container
	StopContainer(name string) error

	// Restart a container
	RestartContainer(name string) error

	// List all images
	ListImages() ([]ContainerImage, error)

	// Get a image
	DetailImage(name string) (ContainerImage, error)

	// Pull a image
	PullImage(name string) error

	// Delete a image
	DeleteImage(name string) error
}

// Define the ContainerInfo

type ContainerInfo struct {
	Id        string
	Name      string
	Image     string
	IpAddress string
	State     ContainerState
	CreatedAt time.Time
	Extra     map[string]string
}

// Define the ContainerImage

type ContainerImage struct {
	Id      string
	Name    string
	Path    string
	Version string
	Extra   map[string]string
}
