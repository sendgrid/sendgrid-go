/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Account Provisioning API
* The Twilio SendGrid Account Provisioning API provides a platform for Twilio SendGrid resellers to manage their customer accounts. This API is for companies that have a formal reseller partnership with Twilio SendGrid.  You can access Twilio SendGrid sub-account functionality without becoming a reseller. If you require sub-account functionality, see the Twilio [SendGrid Subusers](https://docs.sendgrid.com/ui/account-and-settings/subusers) feature, which is available with [Pro and Premier plans](https://sendgrid.com/pricing/).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type CreateAccountParam struct {
	//
	CreateAccountParams *CreateAccountParams `json:"CreateAccountParams"`
	// **OPTIONAL** Custom request header provided ONLY for a test account
	TTestAccount *string `json:"T-Test-Account,omitempty"`
}

func (params *CreateAccountParam) SetCreateAccountParams(CreateAccountParams CreateAccountParams) *CreateAccountParam {
	params.CreateAccountParams = &CreateAccountParams
	return params
}
func (params *CreateAccountParam) SetTTestAccount(TTestAccount string) *CreateAccountParam {
	params.TTestAccount = &TTestAccount
	return params
}

// Creates a new account, with specified offering, under the organization.
func (c *ApiService) CreateAccount(params *CreateAccountParam) (interface{}, error) {
	path := "/v3/partners/accounts"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.CreateAccountParams != nil {
		b, err := json.Marshal(*params.CreateAccountParams)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.TTestAccount != nil {
		headers["T-Test-Account"] = *params.TTestAccount
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &AccountProvisioningAccountId{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}