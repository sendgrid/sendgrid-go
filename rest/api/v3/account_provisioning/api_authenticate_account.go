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

type AuthenticateAccountParam struct {
	// Twilio SendGrid account ID
	AccountID *string `json:"accountID"`
}

func (params *AuthenticateAccountParam) SetAccountID(AccountID string) *AuthenticateAccountParam {
	params.AccountID = &AccountID
	return params
}

// Authenticates and logs in a user to Twilio Sendgrid as a specific admin identity configured for SSO by partner. Any additional teammates or subusers will need to log in directly via app.sendgrid.com
func (c *ApiService) AuthenticateAccount(params *AuthenticateAccountParam) (interface{}, error) {
	path := "/v3/partners/accounts/{AccountID}/sso"
	if params != nil && params.AccountID != nil {
		path = strings.Replace(path, "{"+"AccountID"+"}", *params.AccountID, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}