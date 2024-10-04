/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Link Branding API
* The Twilio SendGrid Link Branding API allows you to configure your domain settings so that all of the click-tracked links, opens, and images in your emails are served from your domain rather than `sendgrid.net`. Spam filters and recipient servers look at the links within emails to determine whether the email appear trustworthy. They use the reputation of the root domain to determine whether the links can be trusted.  You can also manage Link Branding in the **Sender Authentication** section of the Twilio SendGrid application user interface.   See [**How to Set Up Link Branding**](https: //sendgrid.com/docs/ui/account-and-settings/how-to-set-up-link-branding/) for more information.
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

	"strings"
)

type UpdateBrandedLinkParam struct {
	// The ID of the branded link you want to retrieve.
	Id *int32 `json:"id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateBrandedLinkRequest *UpdateBrandedLinkRequest `json:"UpdateBrandedLinkRequest,omitempty"`
}

func (params *UpdateBrandedLinkParam) SetId(Id int32) *UpdateBrandedLinkParam {
	params.Id = &Id
	return params
}
func (params *UpdateBrandedLinkParam) SetOnbehalfof(Onbehalfof string) *UpdateBrandedLinkParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateBrandedLinkParam) SetUpdateBrandedLinkRequest(UpdateBrandedLinkRequest UpdateBrandedLinkRequest) *UpdateBrandedLinkParam {
	params.UpdateBrandedLinkRequest = &UpdateBrandedLinkRequest
	return params
}

// **This endpoint allows you to update a specific branded link. You can use this endpoint to change a branded link's default status.**  You can submit this request as one of your subusers if you include their ID in the `on-behalf-of` header in the request.
func (c *ApiService) UpdateBrandedLink(params *UpdateBrandedLinkParam) (interface{}, error) {
	path := "/v3/whitelabel/links/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", fmt.Sprint(*params.Id), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateBrandedLinkRequest != nil {
		b, err := json.Marshal(*params.UpdateBrandedLinkRequest)
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
	if resp.StatusCode == 200 {
		ps := &LinkBranding200{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
