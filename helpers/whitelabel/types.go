package whitelabel

type SubdomainInfo struct {
	Domain    string `json:"domain"`
	Subdomain string `json:"subdomain,omitempty"`
}

func NewSubdomainInfo(domain, subdomain string) SubdomainInfo {
	return SubdomainInfo{domain, subdomain}
}

type DNSEntry struct {
	Valid bool   `json:"valid"`
	Type  string `json:"type"`
	Host  string `json:"host"`
	Data  string `json:"data"`
}
