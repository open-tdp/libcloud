package drivers

import (
	"github.com/open-tdp/libcloud/dns"
)

type AlibabaAlidnsDriver struct {
	// add any necessary fields here
}

func (p *AlibabaAlidnsDriver) ListZones() ([]*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) DetailZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) CreateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) UpdateZone(zone *dns.Zone) (*dns.Zone, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) DeleteZone(zone *dns.Zone) error {
	// TODO
	return nil
}

func (p *AlibabaAlidnsDriver) ListRecords(zone *dns.Zone) ([]*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) DetailRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) CreateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) UpdateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {
	// TODO
	return nil, nil
}

func (p *AlibabaAlidnsDriver) DeleteRecord(zone *dns.Zone, record *dns.Record) error {
	// TODO
	return nil
}
