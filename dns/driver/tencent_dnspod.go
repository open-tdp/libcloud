package drivers

import (
	"github.com/open-tdp/libcloud/dns"
)

type TecentDnspodDriver struct {
	// add any necessary fields here
}

func (p *TecentDnspodDriver) ListZones() ([]*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) DetailZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) CreateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) UpdateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) DeleteZone(zone *dns.Zone) error {
	// TODO
	return nil
}

func (p *TecentDnspodDriver) ListRecords(zone *dns.Zone) ([]*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) DetailRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) CreateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) UpdateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *TecentDnspodDriver) DeleteRecord(zone *dns.Zone, record *dns.Record) error {
	// TODO
	return nil
}
