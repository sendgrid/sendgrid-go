/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Address Management API
* The Twilio SendGrid IP Address Management API combines functionality that was previously split between the Twilio SendGrid [IP Address API](https://docs.sendgrid.com/api-reference/ip-address) and [IP Pools API](https://docs.sendgrid.com/api-reference/ip-pools). This functionality includes adding IP addresses to your account, assigning IP addresses to IP Pools and Subusers, among other tasks.
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

type AddIpParam struct {
	//
	AddIpRequest *AddIpRequest `json:"AddIpRequest,omitempty"`
}

func (params *AddIpParam) SetAddIpRequest(AddIpRequest AddIpRequest) *AddIpParam {
	params.AddIpRequest = &AddIpRequest
	return params
}

// This operation adds a Twilio SendGrid IP address to your account. You can also assign up to 100 Subusers to the IP address at creation.
func (c *ApiService) AddIp(params *AddIpParam) (interface{}, error) {
	path := "/v3/send_ips/ips"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.AddIpRequest != nil {
		b, err := json.Marshal(*params.AddIpRequest)
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
	if resp.StatusCode == 201 {
		ps := &AddIp201Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &IpAddressManagementErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 401 {
		ps := &IpAddressManagementErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 500 {
		ps := &IpAddressManagementErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
