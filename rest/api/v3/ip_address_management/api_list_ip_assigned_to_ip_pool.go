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

type ListIpAssignedToIpPoolParam struct {
	// Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
	Limit *int32 `json:"limit,omitempty"`
	// Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
	AfterKey *int32 `json:"after_key,omitempty"`
	// Boolean indicating whether or not to return the IP Pool's region information in the response.
	IncludeRegion *bool `json:"include_region,omitempty"`
}

func (params *ListIpAssignedToIpPoolParam) SetLimit(Limit int32) *ListIpAssignedToIpPoolParam {
	params.Limit = &Limit
	return params
}
func (params *ListIpAssignedToIpPoolParam) SetAfterKey(AfterKey int32) *ListIpAssignedToIpPoolParam {
	params.AfterKey = &AfterKey
	return params
}
func (params *ListIpAssignedToIpPoolParam) SetIncludeRegion(IncludeRegion bool) *ListIpAssignedToIpPoolParam {
	params.IncludeRegion = &IncludeRegion
	return params
}

// This operation returns the IP addresses that are assigned to the specified IP pool.
func (c *ApiService) ListIpAssignedToIpPool(params *ListIpAssignedToIpPoolParam) (interface{}, error) {
	path := "/v3/send_ips/pools/{Poolid}/ips"

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
	if params != nil && params.IncludeRegion != nil {
		data.Set("include_region", fmt.Sprint(*params.IncludeRegion))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &ListIpAssignedToIpPool200Response{}
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