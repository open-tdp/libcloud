package drivers

import (
	"github.com/open-tdp/go-libcloud/dns"
	"github.com/open-tdp/go-libcloud/provider"
	"github.com/open-tdp/go-libcloud/provider/alibaba"

	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AlibabaAlidnsDriver struct {
	client *alibaba.Client
	alidns *alidns.Client
	rq     *provider.ReqeustParam
}

func (p *AlibabaAlidnsDriver) ListZones() ([]*dns.Zone, error) {

	resp, err := p.alidns.DescribeDomains(&alidns.DescribeDomainsRequest{})

	if err != nil {
		return nil, err
	}

	zones := make([]*dns.Zone, 0)

	for _, domain := range resp.Body.Domains.Domain {
		zones = append(zones, &dns.Zone{
			Id:         *domain.DomainId,
			Domain:     *domain.DomainName,
			CreateTime: int(*domain.CreateTimestamp),
		})
	}

	return zones, nil

}

func (p *AlibabaAlidnsDriver) DetailZone(domain string) (*dns.Zone, error) {

	resp, err := p.alidns.DescribeDomainInfo(&alidns.DescribeDomainInfoRequest{
		DomainName: tea.String(domain),
	})

	if err != nil {
		return nil, err
	}

	dnsServers := make([]string, 0)
	for _, dnsServer := range resp.Body.DnsServers.DnsServer {
		dnsServers = append(dnsServers, *dnsServer)
	}

	zone := &dns.Zone{
		Id:          *resp.Body.DomainId,
		Domain:      *resp.Body.DomainName,
		DnsServers:  dnsServers,
		MinTTL:      int(*resp.Body.MinTtl),
		Description: *resp.Body.Remark,
	}

	return zone, nil

}

func (p *AlibabaAlidnsDriver) CreateZone(domain string) (*dns.Zone, error) {

	resp, err := p.alidns.AddDomain(&alidns.AddDomainRequest{
		DomainName: tea.String(domain),
	})

	if err != nil {
		return nil, err
	}

	zone := &dns.Zone{
		Id:     *resp.Body.DomainId,
		Domain: *resp.Body.DomainName,
	}

	return zone, nil

}

func (p *AlibabaAlidnsDriver) UpdateZone(zone *dns.Zone) (*dns.Zone, error) {

	_, err := p.alidns.UpdateDomainRemark(&alidns.UpdateDomainRemarkRequest{
		DomainName: tea.String(zone.Domain),
		Remark:     tea.String(zone.Description),
	})

	if err != nil {
		return nil, err
	}

	return zone, nil

}

func (p *AlibabaAlidnsDriver) DeleteZone(zone *dns.Zone) error {

	_, err := p.alidns.DeleteDomain(&alidns.DeleteDomainRequest{
		DomainName: tea.String(zone.Domain),
	})

	return err

}

func (p *AlibabaAlidnsDriver) ListRecords(zone *dns.Zone) ([]*dns.Record, error) {

	resp, err := p.alidns.DescribeDomainRecords(&alidns.DescribeDomainRecordsRequest{
		DomainName: tea.String(zone.Domain),
	})

	if err != nil {
		return nil, err
	}

	records := make([]*dns.Record, 0)

	for _, record := range resp.Body.DomainRecords.Record {
		recordType := dns.RecordType(*record.Type)

		records = append(records, &dns.Record{
			Id:    *record.RecordId,
			Name:  *record.RR,
			Type:  recordType,
			Value: *record.Value,
			TTL:   int(*record.TTL),
		})
	}

	return records, nil

}

func (p *AlibabaAlidnsDriver) DetailRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {

	resp, err := p.alidns.DescribeDomainRecordInfo(&alidns.DescribeDomainRecordInfoRequest{
		RecordId: tea.String(record.Id),
	})

	if err != nil {
		return nil, err
	}

	recordType := dns.RecordType(*resp.Body.Type)

	record = &dns.Record{
		Id:    *resp.Body.RecordId,
		Name:  *resp.Body.RR,
		Type:  recordType,
		Value: *resp.Body.Value,
		TTL:   int(*resp.Body.TTL),
		Zone:  zone,
	}

	return record, nil

}

func (p *AlibabaAlidnsDriver) CreateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {

	resp, err := p.alidns.AddDomainRecord(&alidns.AddDomainRecordRequest{
		DomainName: tea.String(zone.Domain),
		RR:         tea.String(record.Name),
		Type:       tea.String(string(record.Type)),
		Value:      tea.String(record.Value),
		TTL:        tea.Int64(int64(record.TTL)),
	})

	if err != nil {
		return nil, err
	}

	record = &dns.Record{
		Id:   *resp.Body.RecordId,
		Zone: zone,
	}

	return record, nil

}

func (p *AlibabaAlidnsDriver) UpdateRecord(zone *dns.Zone, record *dns.Record) (*dns.Record, error) {

	_, err := p.alidns.UpdateDomainRecord(&alidns.UpdateDomainRecordRequest{
		RecordId: tea.String(record.Id),
		RR:       tea.String(record.Name),
		Type:     tea.String(string(record.Type)),
		Value:    tea.String(record.Value),
		TTL:      tea.Int64(int64(record.TTL)),
	})

	if err != nil {
		return nil, err
	}

	return record, nil

}

func (p *AlibabaAlidnsDriver) DeleteRecord(zone *dns.Zone, record *dns.Record) error {

	_, err := p.alidns.DeleteDomainRecord(&alidns.DeleteDomainRecordRequest{
		RecordId: tea.String(record.Id),
	})

	return err

}
