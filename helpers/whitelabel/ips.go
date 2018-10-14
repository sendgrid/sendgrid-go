package whitelabel

import (
	"github.com/sendgrid/sendgrid-go"
	"encoding/json"
)

type ReverseDNSRequest struct {
	IP	string	`json:"ip"`
	SubdomainInfo
}

type ReverseDNSIP struct {
	ID			int64	`json:"id"`

	ReverseDNSRequest

	Rdns		string	`json:"rdns"`
	Users		[]User	`json:"users"`

	Valid		bool	`json:"valid"`
	Legacy		bool	`json:"legacy"`

	ARecord		DNSEntry	`json:"a_record"`
}

const ipsEndpoint = "/v3/whitelabel/ips"

// CreateReverseDNS sets up a reverse DNS ip.
func CreateReverseDNS(key, ip string, sub SubdomainInfo) (ReverseDNSIP, error) {
	var err error
	var rdip ReverseDNSIP
	cl := sendgrid.NewClientForEndpoint(key, ipsEndpoint)
	cl.Method = "POST"
	cl.Body, err = json.Marshal(ReverseDNSRequest{ip, sub})
	if err != nil {
		return rdip, err
	}

	resp, err := sendgrid.MakeRequestRetry(cl.Request)
	if err != nil {
		return rdip, err
	}

	err = json.Unmarshal([]byte(resp.Body), &rdip)
	return rdip, err
}

