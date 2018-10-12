package whitelabel

import (
	"encoding/json"
	"github.com/sendgrid/sendgrid-go"
)

type brandedLinkRequest struct {
	SubdomainInfo
	IsDefault bool `json:"default"`
}

type BrandedLinkResponse struct {
	SubdomainInfo
	Id int64 `json:"id"`

	Username string `json:"username"`
	UserId   int64  `json:"user_id"`

	IsDefault bool                `json:"default"`
	IsValid   bool                `json:"valid"`
	isLegacy  bool                `json:"legacy"`
	DNS       map[string]DNSEntry `json:"dns"`
}

// CreateBrandedLink adds a new domain for branded links. For more info, see: https://sendgrid.com/docs/API_Reference/api_v3.html
func CreateBrandedLink(key string, sub SubdomainInfo, isDefault bool) (BrandedLinkResponse, error) {
	cl := sendgrid.NewClientForEndpoint(key, "/v3/whitelabel/links")
	var err error
	var resp = BrandedLinkResponse{}

	cl.Body, err = json.Marshal(brandedLinkRequest{sub, isDefault})
	if err != nil {
		return resp, err
	}

	cl.Method = "POST"
	apiResp, err := sendgrid.MakeRequestRetry(cl.Request)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal([]byte(apiResp.Body), &resp)
	return resp, err
}
