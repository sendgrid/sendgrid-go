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

type ListIpPoolParam struct {
	// Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
	Limit *int32 `json:"limit,omitempty"`
	// Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
	AfterKey *int32 `json:"after_key,omitempty"`
	// Specifies an IP address. The `ip` query parameter can be used to filter results by IP address.
	Ip *string `json:"ip,omitempty"`
	// Allowed values are `all`, `eu`, and `us`. If you provide a specific region, results will include all pools that have at least one IP in the filtered region. If `all`, pools with at least one IP (regardless of region) will be returned. If the `region` filter is not provided, the query returns all pools, including empty ones.
	Region *Region7 `json:"region,omitempty"`
	// Boolean indicating whether or not to return the IP Pool's region information in the response.
	IncludeRegion *bool `json:"include_region,omitempty"`
}

func (params *ListIpPoolParam) SetLimit(Limit int32) *ListIpPoolParam {
	params.Limit = &Limit
	return params
}
func (params *ListIpPoolParam) SetAfterKey(AfterKey int32) *ListIpPoolParam {
	params.AfterKey = &AfterKey
	return params
}
func (params *ListIpPoolParam) SetIp(Ip string) *ListIpPoolParam {
	params.Ip = &Ip
	return params
}
func (params *ListIpPoolParam) SetRegion(Region Region7) *ListIpPoolParam {
	params.Region = &Region
	return params
}
func (params *ListIpPoolParam) SetIncludeRegion(IncludeRegion bool) *ListIpPoolParam {
	params.IncludeRegion = &IncludeRegion
	return params
}

// This operation returns a list of your IP Pools and a sample of each Pools' associated IP addresses.  A maximum of 10 IPs will be returned per IP Pool by default. To retrieve additional IP addresses associated with a Pool, use the \"Get IPs Assigned to an IP Pool\" operation. Each user may have a maximum of 100 IP Pools.
func (c *ApiService) ListIpPool(params *ListIpPoolParam) (interface{}, error) {
	path := "/v3/send_ips/pools"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.AfterKey != nil {
		data.Set("after_key", fmt.Sprint(*params.AfterKey))
	}
	if params != nil && params.Ip != nil {
		data.Set("ip", *params.Ip)
	}
	if params != nil && params.Region != nil {
		data.Set("region", fmt.Sprint(*params.Region))
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
		ps := &ListIpPool200Response{}
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
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
