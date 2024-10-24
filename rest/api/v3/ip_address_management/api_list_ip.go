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

type ListIpParam struct {
	// Specifies an IP address. The `ip` query parameter can be used to filter results by IP address.
	Ip *string `json:"ip,omitempty"`
	// Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
	Limit *int32 `json:"limit,omitempty"`
	// Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
	AfterKey *int32 `json:"after_key,omitempty"`
	// Specifies which items to be returned by the API. When the `before_key` is specified, the API will return items beginning from the first item before the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
	BeforeKey *string `json:"before_key,omitempty"`
	// Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account.
	IsLeased *bool `json:"is_leased,omitempty"`
	// Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// A parent must be assigned to an IP address before the parent can send mail from the IP and before the address can be assigned to an IP pool. Set this parameter value to true to allow the parent to send mail from the IP and make the IP eligible for IP pool assignment using the IP pool endpoints.
	IsParentAssigned *bool `json:"is_parent_assigned,omitempty"`
	// Specifies the unique ID for an IP Pool. When included, only IP addresses belonging to the specified Pool will be returned.
	Pool *string `json:"pool,omitempty"`
	// The `start_added_at` and `end_added_at` parameters are used to set a time window. IP addresses that were added to your account in the specified time window will be returned. The `start_added_at` parameter sets the beginning of the time window.
	StartAddedAt *int32 `json:"start_added_at,omitempty"`
	// The `start_added_at` and `end_added_at` parameters are used to set a time window. IP addresses that were added to your account in the specified time window will be returned. The `end_added_at` parameter sets the end of the time window.
	EndAddedAt *int32 `json:"end_added_at,omitempty"`
	// Allowed values are `all`, `eu`, and `us`. If you provide a specific region, results will include all pools that have at least one IP in the filtered region. If `all`, pools with at least one IP (regardless of region) will be returned. If the `region` filter is not provided, the query returns all pools, including empty ones.
	Region *Region7 `json:"region,omitempty"`
	// Boolean indicating whether or not to return the IP Pool's region information in the response.
	IncludeRegion *bool `json:"include_region,omitempty"`
}

func (params *ListIpParam) SetIp(Ip string) *ListIpParam {
	params.Ip = &Ip
	return params
}
func (params *ListIpParam) SetLimit(Limit int32) *ListIpParam {
	params.Limit = &Limit
	return params
}
func (params *ListIpParam) SetAfterKey(AfterKey int32) *ListIpParam {
	params.AfterKey = &AfterKey
	return params
}
func (params *ListIpParam) SetBeforeKey(BeforeKey string) *ListIpParam {
	params.BeforeKey = &BeforeKey
	return params
}
func (params *ListIpParam) SetIsLeased(IsLeased bool) *ListIpParam {
	params.IsLeased = &IsLeased
	return params
}
func (params *ListIpParam) SetIsEnabled(IsEnabled bool) *ListIpParam {
	params.IsEnabled = &IsEnabled
	return params
}
func (params *ListIpParam) SetIsParentAssigned(IsParentAssigned bool) *ListIpParam {
	params.IsParentAssigned = &IsParentAssigned
	return params
}
func (params *ListIpParam) SetPool(Pool string) *ListIpParam {
	params.Pool = &Pool
	return params
}
func (params *ListIpParam) SetStartAddedAt(StartAddedAt int32) *ListIpParam {
	params.StartAddedAt = &StartAddedAt
	return params
}
func (params *ListIpParam) SetEndAddedAt(EndAddedAt int32) *ListIpParam {
	params.EndAddedAt = &EndAddedAt
	return params
}
func (params *ListIpParam) SetRegion(Region Region7) *ListIpParam {
	params.Region = &Region
	return params
}
func (params *ListIpParam) SetIncludeRegion(IncludeRegion bool) *ListIpParam {
	params.IncludeRegion = &IncludeRegion
	return params
}

// This operation returns a list of all IP addresses associated with your account. A sample of IP details is returned with each IP, including which Pools the IP is associated with, whether the IP is set to warm up automatically, and when the IP was last updated.  ### Limitations  The `is_parent_assigned` parameter and `pool` parameter cannot be used at the same time. By definition, an IP cannot be assigned to a Pool if it is not first enabled. You can use either the `before_key` or `after_key` in combination with the `limit` parameter to iterate through paginated results but not both.
func (c *ApiService) ListIp(params *ListIpParam) (interface{}, error) {
	path := "/v3/send_ips/ips"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Ip != nil {
		data.Set("ip", *params.Ip)
	}
	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.AfterKey != nil {
		data.Set("after_key", fmt.Sprint(*params.AfterKey))
	}
	if params != nil && params.BeforeKey != nil {
		data.Set("before_key", *params.BeforeKey)
	}
	if params != nil && params.IsLeased != nil {
		data.Set("is_leased", fmt.Sprint(*params.IsLeased))
	}
	if params != nil && params.IsEnabled != nil {
		data.Set("is_enabled", fmt.Sprint(*params.IsEnabled))
	}
	if params != nil && params.IsParentAssigned != nil {
		data.Set("is_parent_assigned", fmt.Sprint(*params.IsParentAssigned))
	}
	if params != nil && params.Pool != nil {
		data.Set("pool", *params.Pool)
	}
	if params != nil && params.StartAddedAt != nil {
		data.Set("start_added_at", fmt.Sprint(*params.StartAddedAt))
	}
	if params != nil && params.EndAddedAt != nil {
		data.Set("end_added_at", fmt.Sprint(*params.EndAddedAt))
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
		ps := &ListIp200Response{}
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