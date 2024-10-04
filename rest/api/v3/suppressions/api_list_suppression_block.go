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
	"fmt"
	"net/http"
	"net/url"
)

type ListSuppressionBlockParam struct {
	// The start of the time range when a blocked email was created (inclusive). This is a unix timestamp.
	StartTime *int32 `json:"start_time,omitempty"`
	// The end of the time range when a blocked email was created (inclusive). This is a unix timestamp.
	EndTime *int32 `json:"end_time,omitempty"`
	// `limit` sets the page size, i.e. maximum number of items from the list to be returned for a single API request. If omitted, the default page size is used. The maximum page size for this endpoint is 500 items per page.
	Limit *int32 `json:"limit,omitempty"`
	// The number of items in the list to skip over before starting to retrieve the items for the requested page. The default `offset` of `0` represents the beginning of the list, i.e. the start of the first page. To request the second page of the list, set the `offset` to the page size as determined by `limit`. Use multiples of the page size as your `offset` to request further consecutive pages. E.g. assume your page size is set to `10`. An `offset` of `10` requests the second page, an `offset` of `20` requests the third page and so on, provided there are sufficiently many items in your list.
	Offset *int32 `json:"offset,omitempty"`
	// Specifies which records to return based on the records' associated email addresses. For example, `sales` returns records with email addresses that start with 'sales', such as `salesdepartment@example.com` or `sales@example.com`.  You can also use `%25` as a wildcard. For example, `%25market` returns records containing email addresses with the string 'market' anywhere in the email address, and `%25market%25tree` returns records containing email addresses with the string 'market' followed by the string 'tree'. Any reserved characters should be [percent-encoded](https://en.wikipedia.org/wiki/Percent-encoding#Reserved_characters), e.g., the `@` symbol should be encoded as `%40`.
	Email *string `json:"email,omitempty"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *ListSuppressionBlockParam) SetStartTime(StartTime int32) *ListSuppressionBlockParam {
	params.StartTime = &StartTime
	return params
}
func (params *ListSuppressionBlockParam) SetEndTime(EndTime int32) *ListSuppressionBlockParam {
	params.EndTime = &EndTime
	return params
}
func (params *ListSuppressionBlockParam) SetLimit(Limit int32) *ListSuppressionBlockParam {
	params.Limit = &Limit
	return params
}
func (params *ListSuppressionBlockParam) SetOffset(Offset int32) *ListSuppressionBlockParam {
	params.Offset = &Offset
	return params
}
func (params *ListSuppressionBlockParam) SetEmail(Email string) *ListSuppressionBlockParam {
	params.Email = &Email
	return params
}
func (params *ListSuppressionBlockParam) SetOnbehalfof(Onbehalfof string) *ListSuppressionBlockParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve a paginated list of all email addresses that are currently on your blocks list.**  You can use the `limit` query parameter to set the page size. If your list contains more items than the page size permits, you can make multiple requests. Use the `offset` query parameter to control the position in the list from which to start retrieving additional items.
func (c *ApiService) ListSuppressionBlock(params *ListSuppressionBlockParam) (interface{}, error) {
	path := "/v3/suppression/blocks"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.StartTime != nil {
		data.Set("start_time", fmt.Sprint(*params.StartTime))
	}
	if params != nil && params.EndTime != nil {
		data.Set("end_time", fmt.Sprint(*params.EndTime))
	}
	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.Offset != nil {
		data.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params != nil && params.Email != nil {
		data.Set("email", *params.Email)
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
		ps := &[]BlocksResponseInner{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
