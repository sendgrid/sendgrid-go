/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid User API
* The Twilio SendGrid User API allows you to modify the settings of a SendGrid user account such as the user's email address or username. Keeping your user profile up to date helps SendGrid verify who you are and share important communications with you.  See [**Account Details**](https://docs.sendgrid.com/ui/account-and-settings/account) for more information. You can also manage your user settings in the [SendGrid application user interface](https://app.sendgrid.com/account/details).
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

type UpdateEmailParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateEmailRequest *UpdateEmailRequest `json:"UpdateEmailRequest,omitempty"`
}

func (params *UpdateEmailParam) SetOnbehalfof(Onbehalfof string) *UpdateEmailParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateEmailParam) SetUpdateEmailRequest(UpdateEmailRequest UpdateEmailRequest) *UpdateEmailParam {
	params.UpdateEmailRequest = &UpdateEmailRequest
	return params
}

// **This endpoint allows you to update the email address currently on file for your account.**
func (c *ApiService) UpdateEmail(params *UpdateEmailParam) (interface{}, error) {
	path := "/v3/user/email"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateEmailRequest != nil {
		b, err := json.Marshal(*params.UpdateEmailRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Put(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &UpdateEmail200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
