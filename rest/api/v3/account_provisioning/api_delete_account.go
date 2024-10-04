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
	"net/http"
	"net/url"

	"strings"
)

type DeleteAccountParam struct {
	// Twilio SendGrid account ID
	AccountID *string `json:"accountID"`
}

func (params *DeleteAccountParam) SetAccountID(AccountID string) *DeleteAccountParam {
	params.AccountID = &AccountID
	return params
}

// Delete a specific account under your organization by account ID. Note that this is an **irreversible** action that does the following:   - Revokes API Keys and SSO so that the account user cannot log in or access SendGrid data.  - Removes all offerings and configured SendGrid resources such as dedicated IPs.  - Cancels billing effective immediately.
func (c *ApiService) DeleteAccount(params *DeleteAccountParam) (interface{}, error) {
	path := "/v3/partners/accounts/{AccountID}"
	if params != nil && params.AccountID != nil {
		path = strings.Replace(path, "{"+"AccountID"+"}", *params.AccountID, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
