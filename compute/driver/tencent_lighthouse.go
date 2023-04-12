package drivers

import (
	"github.com/open-tdp/go-libcloud/compute"
)

type TencentLighthouseDriver struct {
	// add any necessary fields here
}

func (p *TencentLighthouseDriver) ListLocations() ([]*compute.Location, error) {
	// TDDO
	return nil, nil
}
