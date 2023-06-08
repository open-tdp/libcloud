package dns

// Define a Dns Provider interface that matches the one in Apache Libcloud

type DnsProvider interface {

	// List all zones
	ListZones() ([]*Zone, error)

	// Detail a zones
	DetailZone(zone *Zone) (*Zone, error)

	// Create a new zone
	CreateZone(zone *Zone) error

	// Update an existing zone
	UpdateZone(zone *Zone, record *Record) error

	// Delete an existing zone
	DeleteZone(zone *Zone) error

	// List all records in a zone
	ListRecords(zone *Zone) ([]*Record, error)

	// Detail a record in a zone
	DetailRecord(zone *Zone, record *Record) error

	// Create a new record in a zone
	CreateRecord(zone *Zone, record *Record) error

	// Update an existing record in a zone
	UpdateRecord(zone *Zone, record *Record) error

	// Delete an existing record in a zone
	DeleteRecord(zone *Zone, record *Record) error

	// List Record Types
	ListRecordTypes() ([]RecordType, error)
}

// Zone represents a Dns zone

type Zone struct {
	Id          string
	Domain      string
	PunyCode    string
	Type        ZoneType
	CreateTime  int
	DnsServers  []string
	MinTTL      int
	Description string
	Extra       map[string]interface{}
}

// Record represents a Dns record

type Record struct {
	Id          string
	Name        string
	Type        RecordType
	Value       string
	Line        string
	TTL         int
	Priority    int
	Description string
	Extra       map[string]interface{}
}
