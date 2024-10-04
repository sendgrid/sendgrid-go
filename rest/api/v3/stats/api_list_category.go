/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Statistics API
* The Twilio SendGrid Statistics API allows you to retrieve the various statistics related to your email program.  Tracking your emails is an important part of being a good sender and learning about how your users interact with your email. This includes everything from clicks and opens to looking at which browsers and mailbox providers your customers use.  SendGrid has broken up statistics in specific ways so that you can get at-a-glance data, as well as the details of how your email is being used.  Category statistics are available for the previous thirteen months only.  See [**Statistics Overview**](https://docs.sendgrid.com/ui/analytics-and-reporting/stats-overview) for more information.
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

type ListCategoryParam struct {
	// `limit` sets the page size, i.e. maximum number of items from the list to be returned for a single API request. If omitted, the default page size is used.
	Limit *int32 `json:"limit,omitempty"`
	// Allows you to perform a prefix search on this particular category.
	Category *string `json:"category,omitempty"`
	// The number of items in the list to skip over before starting to retrieve the items for the requested page. The default `offset` of `0` represents the beginning of the list, i.e. the start of the first page. To request the second page of the list, set the `offset` to the page size as determined by `limit`. Use multiples of the page size as your `offset` to request further consecutive pages. E.g. assume your page size is set to `10`. An `offset` of `10` requests the second page, an `offset` of `20` requests the third page and so on, provided there are sufficiently many items in your list.
	Offset *int32 `json:"offset,omitempty"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *ListCategoryParam) SetLimit(Limit int32) *ListCategoryParam {
	params.Limit = &Limit
	return params
}
func (params *ListCategoryParam) SetCategory(Category string) *ListCategoryParam {
	params.Category = &Category
	return params
}
func (params *ListCategoryParam) SetOffset(Offset int32) *ListCategoryParam {
	params.Offset = &Offset
	return params
}
func (params *ListCategoryParam) SetOnbehalfof(Onbehalfof string) *ListCategoryParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve a paginated list of all of your categories.**  You can use the `limit` query parameter to set the page size. If your list contains more items than the page size permits, you can make multiple requests. Use the `offset` query parameter to control the position in the list from which to start retrieving additional items.
func (c *ApiService) ListCategory(params *ListCategoryParam) (interface{}, error) {
	path := "/v3/categories"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.Category != nil {
		data.Set("category", *params.Category)
	}
	if params != nil && params.Offset != nil {
		data.Set("offset", fmt.Sprint(*params.Offset))
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
		ps := &[]ListCategory200ResponseInner{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ListCategory400Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
