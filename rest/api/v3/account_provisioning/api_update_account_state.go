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

	"strings"
)

type UpdateAccountStateParam struct {
	// Twilio SendGrid account ID
	AccountID *string `json:"accountID"`
	//
	AccountProvisioningStateWrite *AccountProvisioningStateWrite `json:"AccountProvisioningStateWrite"`
}

func (params *UpdateAccountStateParam) SetAccountID(AccountID string) *UpdateAccountStateParam {
	params.AccountID = &AccountID
	return params
}
func (params *UpdateAccountStateParam) SetAccountProvisioningStateWrite(AccountProvisioningStateWrite AccountProvisioningStateWrite) *UpdateAccountStateParam {
	params.AccountProvisioningStateWrite = &AccountProvisioningStateWrite
	return params
}

// Update the state of the specified account.
func (c *ApiService) UpdateAccountState(params *UpdateAccountStateParam) (interface{}, error) {
	path := "/v3/partners/accounts/{AccountID}/state"
	if params != nil && params.AccountID != nil {
		path = strings.Replace(path, "{"+"AccountID"+"}", *params.AccountID, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.AccountProvisioningStateWrite != nil {
		b, err := json.Marshal(*params.AccountProvisioningStateWrite)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Put(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}