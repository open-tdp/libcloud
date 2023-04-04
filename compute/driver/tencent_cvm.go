package drivers

import (
	"github.com/open-tdp/libcloud/compute"
)

type TencentCvmDriver struct {
	// add any necessary fields here
}

func (p *TencentCvmDriver) ListLocations() ([]*compute.Location, error) {
	// TDDO
	return nil, nil
}
