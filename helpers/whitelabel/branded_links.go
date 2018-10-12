package whitelabel

import (
	"encoding/json"
	"github.com/sendgrid/sendgrid-go"
)

type brandedLinkRequest struct {
	SubdomainInfo
	IsDefault bool `json:"default"`
}

type BrandedLink struct {
	SubdomainInfo
	Id int64 `json:"id"`

	Username string `json:"username"`
	UserId   int64  `json:"user_id"`

	IsDefault bool                `json:"default"`
	IsValid   bool                `json:"valid"`
	isLegacy  bool                `json:"legacy"`
	DNS       map[string]DNSEntry `json:"dns"`
}

const linksEndpoint = "/v3/whitelabel/links"

// CreateBrandedLink adds a new domain for branded links. For more info, see: https://sendgrid.com/docs/API_Reference/api_v3.html
func CreateBrandedLink(key string, sub SubdomainInfo, isDefault bool) (*BrandedLink, error) {
	cl := sendgrid.NewClientForEndpoint(key, linksEndpoint)
	var err error
	var link = BrandedLink{}

	cl.Body, err = json.Marshal(brandedLinkRequest{sub, isDefault})
	if err != nil {
		return &link, err
	}

	cl.Method = "POST"
	resp, err := sendgrid.MakeRequestRetry(cl.Request)
	if err != nil {
		return &link, err
	}

	err = json.Unmarshal([]byte(resp.Body), &link)
	return &link, err
}

// GetBrandedLinks fetches all branded link domains.
func GetBrandedLinks(key string) ([]BrandedLink, error) {
	cl := sendgrid.NewClientForEndpoint(key, linksEndpoint)
	cl.Method = "GET"

	resp, err := sendgrid.MakeRequestRetry(cl.Request)
	if err != nil {
		return nil, err
	}

	var links []BrandedLink
	err = json.Unmarshal([]byte(resp.Body), &links)
	return links, err
}
