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

type UpdateAsmGroupParam struct {
	// The ID of the suppression group you would like to retrieve.
	GroupId *string `json:"group_id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateAsmGroupRequest *UpdateAsmGroupRequest `json:"UpdateAsmGroupRequest,omitempty"`
}

func (params *UpdateAsmGroupParam) SetGroupId(GroupId string) *UpdateAsmGroupParam {
	params.GroupId = &GroupId
	return params
}
func (params *UpdateAsmGroupParam) SetOnbehalfof(Onbehalfof string) *UpdateAsmGroupParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateAsmGroupParam) SetUpdateAsmGroupRequest(UpdateAsmGroupRequest UpdateAsmGroupRequest) *UpdateAsmGroupParam {
	params.UpdateAsmGroupRequest = &UpdateAsmGroupRequest
	return params
}

// **This endpoint allows you to update or change a suppression group.**
func (c *ApiService) UpdateAsmGroup(params *UpdateAsmGroupParam) (interface{}, error) {
	path := "/v3/asm/groups/{GroupId}"
	if params != nil && params.GroupId != nil {
		path = strings.Replace(path, "{"+"GroupId"+"}", *params.GroupId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateAsmGroupRequest != nil {
		b, err := json.Marshal(*params.UpdateAsmGroupRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &SuppressionGroup{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
