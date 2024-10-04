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

	"strings"
)

type DeleteSubUsersFromIpParam struct {
	// The `ip` path parameter specifies an IP address to make the request against.
	Ip *string `json:"ip"`
	//
	DeleteSubUsersFromIpRequest *DeleteSubUsersFromIpRequest `json:"DeleteSubUsersFromIpRequest,omitempty"`
}

func (params *DeleteSubUsersFromIpParam) SetIp(Ip string) *DeleteSubUsersFromIpParam {
	params.Ip = &Ip
	return params
}
func (params *DeleteSubUsersFromIpParam) SetDeleteSubUsersFromIpRequest(DeleteSubUsersFromIpRequest DeleteSubUsersFromIpRequest) *DeleteSubUsersFromIpParam {
	params.DeleteSubUsersFromIpRequest = &DeleteSubUsersFromIpRequest
	return params
}

// This operation removes a batch of Subusers from a specified IP address.
func (c *ApiService) DeleteSubUsersFromIp(params *DeleteSubUsersFromIpParam) (interface{}, error) {
	path := "/v3/send_ips/ips/{Ip}/subusers:batchDelete"
	if params != nil && params.Ip != nil {
		path = strings.Replace(path, "{"+"Ip"+"}", *params.Ip, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.DeleteSubUsersFromIpRequest != nil {
		b, err := json.Marshal(*params.DeleteSubUsersFromIpRequest)
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
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
