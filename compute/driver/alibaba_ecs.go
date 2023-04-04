package drivers

import (
	"github.com/open-tdp/libcloud/compute"
)

type AlibabaEcsDriver struct {
	// add any necessary fields here
}

func (p *AlibabaEcsDriver) ListLocations() ([]*compute.Location, error) {
	// TDDO
	return nil, nil
}
