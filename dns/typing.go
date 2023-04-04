package dns

// Define a DNS Provider interface that matches the one in Apache Libcloud

type DNSProvider interface {
	// Create a new zone
	CreateZone(zone *Zone) error

	// Delete an existing zone
	DeleteZone(zone *Zone) error

	// List all zones
	ListZones() ([]*Zone, error)

	// Create a new record in a zone
	CreateRecord(zone *Zone, record *Record) error

	// Update an existing record in a zone
	UpdateRecord(zone *Zone, record *Record) error

	// Delete an existing record in a zone
	DeleteRecord(zone *Zone, record *Record) error

	// List all records in a zone
	ListRecords(zone *Zone) ([]*Record, error)
}

// Zone represents a DNS zone

type Zone struct {
	ID     string
	Domain string
	Type   string
	TTL    int
	Extra  map[string]interface{}
}

// Zone Type constants

type ZoneType string

const (
	ZoneTypePrimary   ZoneType = "PRIMARY"
	ZoneTypeSecondary ZoneType = "SECONDARY"
)

// Record represents a DNS record

type Record struct {
	ID    string
	Name  string
	Type  string
	Data  string
	Zone  *Zone
	TTL   int
	Extra map[string]interface{}
}

// Record Type constants

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypeNS    RecordType = "NS"
	RecordTypePTR   RecordType = "PTR"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
)
