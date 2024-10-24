/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Email Address Validation API
* The Twilio SendGrid Email Address Validation API provides real-time detailed information on the validity of email addresses. You can integrate this validation process into your platform's signup form and customize the best use of email address validation for your use case.  Email Address Validation is available to [Email API Pro and Premier level accounts](https://sendgrid.com/pricing) only. Prior to upgrading your account to Pro or Premier, you will not see the option to create an Email Validation API key. An Email Validation API key is separate from and in addition to your other keys, including Full Access API keys.  You can use this API to: - Indicate to users that the address they have entered into a form is invalid. - Drop invalid email addresses from your database. - Suppress invalid email addresses from your sending to decrease your bounce rate.  See [**Email Address Validation**](https://docs.sendgrid.com/ui/managing-contacts/email-address-validation) for more information.  You can also view your Email Validation results and metrics in the Validation section of the [Twilio SendGrid application user interface](https://docs.sendgrid.com/ui/managing-contacts/email-address-validation).
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

type ValidateEmailParam struct {
	//
	ValidateEmailRequest *ValidateEmailRequest `json:"ValidateEmailRequest,omitempty"`
}

func (params *ValidateEmailParam) SetValidateEmailRequest(ValidateEmailRequest ValidateEmailRequest) *ValidateEmailParam {
	params.ValidateEmailRequest = &ValidateEmailRequest
	return params
}

// **This endpoint allows you to validate an email address.**
func (c *ApiService) ValidateEmail(params *ValidateEmailParam) (interface{}, error) {
	path := "/v3/validations/email"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.ValidateEmailRequest != nil {
		b, err := json.Marshal(*params.ValidateEmailRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &ValidateEmail200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}