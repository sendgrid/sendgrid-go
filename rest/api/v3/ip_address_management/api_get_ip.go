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
	"fmt"
	"net/http"
	"net/url"
)

type GetIpParam struct {
	// Boolean indicating whether or not to return the IP Pool's region information in the response.
	IncludeRegion *bool `json:"include_region,omitempty"`
}

func (params *GetIpParam) SetIncludeRegion(IncludeRegion bool) *GetIpParam {
	params.IncludeRegion = &IncludeRegion
	return params
}

// This operation returns details for a specified IP address. Details include whether the IP is assigned to a parent account, set to warm up automatically, which Pools the IP is associated with, when the IP was added and modified, whether the IP is leased, and whether the IP is enabled. Note that this operation will not return Subuser information associated with the IP. To retrieve Subuser information, use the \"Get a List of Subusers Assigned to an IP\" endpoint.
func (c *ApiService) GetIp(params *GetIpParam) (interface{}, error) {
	path := "/v3/send_ips/ips/{Ip}"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.IncludeRegion != nil {
		data.Set("include_region", fmt.Sprint(*params.IncludeRegion))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &GetIp200Response{}
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