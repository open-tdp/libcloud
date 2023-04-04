package dns

// Zone Type constants

type ZoneType string

const (
	ZoneTypePrimary   ZoneType = "PRIMARY"
	ZoneTypeSecondary ZoneType = "SECONDARY"
)

// Record Type constants

type RecordType string

const (
	RecordTypeA          RecordType = "A"
	RecordTypeAAAA       RecordType = "AAAA"
	RecordTypeAFSDB      RecordType = "A"
	RecordTypeALIAS      RecordType = "ALIAS"
	RecordTypeCERT       RecordType = "CERT"
	RecordTypeCNAME      RecordType = "CNAME"
	RecordTypeDNAME      RecordType = "DNAME"
	RecordTypeDNSKEY     RecordType = "DNSKEY"
	RecordTypeDS         RecordType = "DS"
	RecordTypeGEO        RecordType = "GEO"
	RecordTypeHINFO      RecordType = "HINFO"
	RecordTypeKEY        RecordType = "KEY"
	RecordTypeLOC        RecordType = "LOC"
	RecordTypeMX         RecordType = "MX"
	RecordTypeNAPTR      RecordType = "NAPTR"
	RecordTypeNS         RecordType = "NS"
	RecordTypeNSEC       RecordType = "NSEC"
	RecordTypeOPENPGPKEY RecordType = "OPENPGPKEY"
	RecordTypePTR        RecordType = "PTR"
	RecordTypeREDIRECT   RecordType = "REDIRECT"
	RecordRP             RecordType = "RP"
	RecordTypeRRSIG      RecordType = "RRSIG"
	RecordTypeSOA        RecordType = "SOA"
	RecordTypeSPF        RecordType = "SPF"
	RecordTypeSRV        RecordType = "SRV"
	RecordTypeSSHFP      RecordType = "SSHFP"
	RecordTypeTLSA       RecordType = "TLSA"
	RecordTypeTXT        RecordType = "TXT"
	RecordTypeURL        RecordType = "URL"
	RecordTypeWKS        RecordType = "WKS"
	RecordTypeCAA        RecordType = "CAA"
)

// Dns Error

type DnsError string

const (
	ZoneError                = "ZoneError"
	ZoneDoesNotExistError    = "ZoneDoesNotExistError"
	ZoneAlreadyExistsError   = "ZoneAlreadyExistsError"
	RecordError              = "RecordError"
	RecordDoesNotExistError  = "RecordDoesNotExistError"
	RecordAlreadyExistsError = "RecordAlreadyExistsError"
)
