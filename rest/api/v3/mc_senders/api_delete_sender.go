/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Senders API
* The Twilio SendGrid Marketing Campaigns Senders API allows you to create a verified sender from which your marketing emails will be sent. To ensure our customers maintain the best possible sender reputations and to uphold legitimate sending behavior, we require customers to verify their Senders. A Sender represents your “From” email address—the address your recipients will see as the sender of your emails.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

type DeleteSenderParam struct {
	// The unique identifier of the Sender.
	Id *int32 `json:"id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *DeleteSenderParam) SetId(Id int32) *DeleteSenderParam {
	params.Id = &Id
	return params
}
func (params *DeleteSenderParam) SetOnbehalfof(Onbehalfof string) *DeleteSenderParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to delete an existing Sender.**
func (c *ApiService) DeleteSender(params *DeleteSenderParam) (interface{}, error) {
	path := "/v3/marketing/senders/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", fmt.Sprint(*params.Id), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
