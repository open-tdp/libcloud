package drivers

import (
	"github.com/open-tdp/libcloud/dns"
)

type CloudflareDnsDriver struct {
	// add any necessary fields here
}

func (p *CloudflareDnsDriver) ListZones() ([]*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) DetailZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) CreateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) UpdateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) DeleteZone(zone *dns.Zone) error {
	// TODO
	return nil
}

func (p *CloudflareDnsDriver) ListRecords(zone *dns.Zone) ([]*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) DetailRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) CreateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) UpdateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *CloudflareDnsDriver) DeleteRecord(zone *dns.Zone, record *dns.Record) error {
	// TODO
	return nil
}
