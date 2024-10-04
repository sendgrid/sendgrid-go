/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Access Management API
* IP Twilio SendGrid IP Access Management API allows you to control which IP addresses can be used to access your account, either through the SendGrid application user interface or the API.  There is no limit to the number of IP addresses that you can allow.  It is possible to remove your own IP address from your list of allowed addresses, thus blocking your own access to your account. While we are able to restore your access, we do require thorough proof of your identify and ownership of your account. We take the security of your account very seriously and wish to prevent any 'bad actors' from maliciously gaining access to your account. Your current IP is clearly displayed to help prevent you from accidentally removing it from the allowed addresses.  See [**IP Access Management**](https://docs.sendgrid.com/ui/account-and-settings/ip-access-management) for more information.
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

type AddIpToAllowListParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	AddIpToAllowListRequest *AddIpToAllowListRequest `json:"AddIpToAllowListRequest,omitempty"`
}

func (params *AddIpToAllowListParam) SetOnbehalfof(Onbehalfof string) *AddIpToAllowListParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *AddIpToAllowListParam) SetAddIpToAllowListRequest(AddIpToAllowListRequest AddIpToAllowListRequest) *AddIpToAllowListParam {
	params.AddIpToAllowListRequest = &AddIpToAllowListRequest
	return params
}

// **This endpoint allows you to add one or more allowed IP addresses.**  To allow one or more IP addresses, pass them to this endpoint in an array. Once an IP address is added to your allow list, it will be assigned an `id` that can be used to remove the address. You can retrieve the ID associated with an IP using the \"Retrieve a list of currently allowed IPs\" endpoint.
func (c *ApiService) AddIpToAllowList(params *AddIpToAllowListParam) (interface{}, error) {
	path := "/v3/access_settings/whitelist"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.AddIpToAllowListRequest != nil {
		b, err := json.Marshal(*params.AddIpToAllowListRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &IpAccessManagement2xx{}
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
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
