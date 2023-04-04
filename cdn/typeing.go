package cdn

// Define a CDN Provider interface that matches the one in Apache Libcloud

type CDNProvider interface {
    // CreateCDN creates a new CDN with the given endpoint and name
    CreateCDN(endpoint string, name string) (*CDN, error)

    // ListCDNs lists all CDNs
    ListCDNs() ([]*CDN, error)

    // GetCDN gets a CDN by ID
    GetCDN(id string) (*CDN, error)

    // UpdateCDN updates a CDN's name by ID
    UpdateCDN(id string, name string) (*CDN, error)

    // DeleteCDN deletes a CDN by ID
    DeleteCDN(id string) (bool, error)
}

// Define a CDN struct

type CDN struct {
    ID      string
    Name    string
    Enabled bool
}
