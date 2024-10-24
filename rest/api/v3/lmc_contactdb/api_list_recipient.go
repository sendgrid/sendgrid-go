/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Legacy Marketing Campaigns Contacts API
* The Twilio SendGrid Legacy Marketing Campaigns Contacts API allows you to manage your marketing contacts programmatically. This API is operational, but we recommend using the current version of Marketing Campaigns' [Contacts API](https://docs.sendgrid.com/api-reference/contacts/), [Lists API](https://docs.sendgrid.com/api-reference/lists/), and [Segments API](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2/) to manage your contacts.  See [**Migrating from Legacy Marketing Campaigns**](https://docs.sendgrid.com/ui/sending-email/migrating-from-legacy-marketing-campaigns) for more information.
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

type ListRecipientParam struct {
	// Page index of first recipients to return (must be a positive integer)
	Page *int32 `json:"page,omitempty"`
	// Number of recipients to return at a time (must be a positive integer between 1 and 1000)
	PageSize *int32 `json:"page_size,omitempty"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *ListRecipientParam) SetPage(Page int32) *ListRecipientParam {
	params.Page = &Page
	return params
}
func (params *ListRecipientParam) SetPageSize(PageSize int32) *ListRecipientParam {
	params.PageSize = &PageSize
	return params
}
func (params *ListRecipientParam) SetOnbehalfof(Onbehalfof string) *ListRecipientParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve all of your Marketing Campaigns recipients.**  Batch deletion of a page makes it possible to receive an empty page of recipients before reaching the end of the list of recipients. To avoid this issue; iterate over pages until a 404 is retrieved.
func (c *ApiService) ListRecipient(params *ListRecipientParam) (interface{}, error) {
	path := "/v3/contactdb/recipients"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Page != nil {
		data.Set("page", fmt.Sprint(*params.Page))
	}
	if params != nil && params.PageSize != nil {
		data.Set("page_size", fmt.Sprint(*params.PageSize))
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
		ps := &ListRecipientsResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 401 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}