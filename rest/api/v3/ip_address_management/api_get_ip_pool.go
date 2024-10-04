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
	"net/url"
)

type GetIpPoolParam struct {
	// Boolean indicating whether or not to return the IP Pool's region information in the response.
	IncludeRegion *bool `json:"include_region,omitempty"`
}

func (params *GetIpPoolParam) SetIncludeRegion(IncludeRegion bool) *GetIpPoolParam {
	params.IncludeRegion = &IncludeRegion
	return params
}

// This operation will return the details for a specified IP Pool, including the Pool's name, ID, a sample list of the IPs associated with the Pool, and the total number of IPs belonging to the Pool.  A maximum of 10 IPs will be returned per IP Pool by default. To retrieve additional IP addresses associated with a Pool, use the \"Get IPs Assigned to an IP Pool\" operation.
func (c *ApiService) GetIpPool(params *GetIpPoolParam) (interface{}, error) {
	path := "/v3/send_ips/pools/{Poolid}"

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
		ps := &GetIpPool200Response{}
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
	return nil, nil
}
