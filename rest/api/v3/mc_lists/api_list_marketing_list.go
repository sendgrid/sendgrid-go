/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Lists API
* The Twilio SendGrid Marketing Campaigns Lists API allows you to manage your contacts lists programmatically. Lists are static collections of Marketing Campaigns contacts. You can use this API to interact with the list objects themselves. To add contacts to a list, you must use the [Contacts API](https://docs.sendgrid.com/api-reference/contacts/).  You can also manage your lists using the Contacts menu in the [Marketing Campaigns application user interface](https://mc.sendgrid.com/contacts). For more information about lists and best practices for building them, see [**Building your Contact Lists**](https://sendgrid.com/docs/ui/managing-contacts/building-your-contact-list/).
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

type ListMarketingListParam struct {
	// Maximum number of elements to return. Defaults to 100, returns 1000 max
	PageSize *float32 `json:"page_size,omitempty"`
	//
	PageToken *string `json:"page_token,omitempty"`
}

func (params *ListMarketingListParam) SetPageSize(PageSize float32) *ListMarketingListParam {
	params.PageSize = &PageSize
	return params
}
func (params *ListMarketingListParam) SetPageToken(PageToken string) *ListMarketingListParam {
	params.PageToken = &PageToken
	return params
}

// **This endpoint returns an array of all of your contact lists.**
func (c *ApiService) ListMarketingList(params *ListMarketingListParam) (interface{}, error) {
	path := "/v3/marketing/lists"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if params != nil && params.PageToken != nil {
		data.Set("page_token", *params.PageToken)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &ListMarketingList200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
