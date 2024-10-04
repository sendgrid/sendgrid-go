/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Suppressions API
* The Twilio SendGrid Suppressions API allows you to manage your Suppressions or Unsubscribes and Suppression or Unsubscribe groups. With SendGrid, an unsubscribe is the action an email recipient takes when they opt-out of receiving your messages. A suppression is the action you take as a sender to filter or suppress an unsubscribed address from your email send. From the perspective of the recipient, your suppression is the result of their unsubscribe.  You can have global suppressions, which represent addresses that have been unsubscribed from all of your emails. You can also have suppression groups, also known as ASM groups, which represent categories or groups of emails that your recipients can unsubscribe from, rather than unsubscribing from all of your messages.  SendGrid automatically suppresses emails sent to users for a variety of reasons, including blocks, bounces, invalid email addresses, spam reports, and unsubscribes. SendGrid suppresses these messages to help you maintain the best possible sender reputation by attempting to prevent unwanted mail. You may also add addresses to your suppressions.  See [**Suppressions**](https://docs.sendgrid.com/for-developers/sending-email/suppressions) for more information.
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

type GetGlobalSuppressionParam struct {
	// The email address of the global suppression you want to retrieve. Or, if you want to check if an email address is on the global suppressions list, enter that email address here.
	Email *string `json:"email"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *GetGlobalSuppressionParam) SetEmail(Email string) *GetGlobalSuppressionParam {
	params.Email = &Email
	return params
}
func (params *GetGlobalSuppressionParam) SetOnbehalfof(Onbehalfof string) *GetGlobalSuppressionParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve a global suppression. You can also use this endpoint to confirm if an email address is already globally suppresed.**  If the email address you include in the URL path parameter `{email}` is already globally suppressed, the response will include that email address. If the address you enter for `{email}` is not globally suppressed, an empty JSON object `{}` will be returned.
func (c *ApiService) GetGlobalSuppression(params *GetGlobalSuppressionParam) (interface{}, error) {
	path := "/v3/asm/suppressions/global/{Email}"
	if params != nil && params.Email != nil {
		path = strings.Replace(path, "{"+"Email"+"}", *params.Email, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &RetrieveAGlobalSuppressionResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
