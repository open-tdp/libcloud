package drivers

import (
	"github.com/opentdp/go-libcloud/container"
)

type DockerDriver struct {
	// add any necessary fields here
}

func (p *DockerDriver) ListContainers() ([]*container.ContainerInfo, error) {
	// TODO
	return nil, nil
}
