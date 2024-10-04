/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Subusers
* The Twilio SendGrid Subusers API allows you to create and manage your Subuser accounts. Subusers are available on [Pro and Premier plans](https://sendgrid.com/pricing), and you can think of them as sub-accounts. Each Subuser can have its own sending domains, IP addresses, and reporting. SendGrid recommends creating Subusers for each of the different types of emails you send—one Subuser for transactional emails and another for marketing emails. Independent Software Vendor (ISV) customers may also create Subusers for each of their customers.  You can also manage Subusers in the [Twilio SendGrid application user interface](https://app.sendgrid.com/settings/subusers). See [**Subusers**](https://docs.sendgrid.com/ui/account-and-settings/subusers) for more information.
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

type ListSubuserMonthlyStatParam struct {
	// The date of the month to retrieve statistics for. Must be formatted YYYY-MM-DD
	Date *string `json:"date"`
	// The username of the Subuser.
	SubuserName *string `json:"subuser_name"`
	// The metric that you want to sort by. Metrics that you can sort by are: `blocks`, `bounces`, `clicks`, `delivered`, `opens`, `requests`, `unique_clicks`, `unique_opens`, and `unsubscribes`.'
	SortByMetric *string `json:"sort_by_metric,omitempty"`
	// The direction you want to sort.
	SortByDirection *SortByDirection `json:"sort_by_direction,omitempty"`
	// Optional field to limit the number of results returned.
	Limit *int32 `json:"limit,omitempty"`
	// Optional beginning point in the list to retrieve from.
	Offset *int32 `json:"offset,omitempty"`
}

func (params *ListSubuserMonthlyStatParam) SetDate(Date string) *ListSubuserMonthlyStatParam {
	params.Date = &Date
	return params
}
func (params *ListSubuserMonthlyStatParam) SetSubuserName(SubuserName string) *ListSubuserMonthlyStatParam {
	params.SubuserName = &SubuserName
	return params
}
func (params *ListSubuserMonthlyStatParam) SetSortByMetric(SortByMetric string) *ListSubuserMonthlyStatParam {
	params.SortByMetric = &SortByMetric
	return params
}
func (params *ListSubuserMonthlyStatParam) SetSortByDirection(SortByDirection SortByDirection) *ListSubuserMonthlyStatParam {
	params.SortByDirection = &SortByDirection
	return params
}
func (params *ListSubuserMonthlyStatParam) SetLimit(Limit int32) *ListSubuserMonthlyStatParam {
	params.Limit = &Limit
	return params
}
func (params *ListSubuserMonthlyStatParam) SetOffset(Offset int32) *ListSubuserMonthlyStatParam {
	params.Offset = &Offset
	return params
}

// **This endpoint allows you to retrive the monthly email statistics for a specific subuser.**  When using the `sort_by_metric` to sort your stats by a specific metric, you can not sort by the following metrics: `bounce_drops`, `deferred`, `invalid_emails`, `processed`, `spam_report_drops`, `spam_reports`, or `unsubscribe_drops`.
func (c *ApiService) ListSubuserMonthlyStat(params *ListSubuserMonthlyStatParam) (interface{}, error) {
	path := "/v3/subusers/{SubuserName}/stats/monthly"
	if params != nil && params.SubuserName != nil {
		path = strings.Replace(path, "{"+"SubuserName"+"}", *params.SubuserName, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Date != nil {
		data.Set("date", *params.Date)
	}
	if params != nil && params.SortByMetric != nil {
		data.Set("sort_by_metric", *params.SortByMetric)
	}
	if params != nil && params.SortByDirection != nil {
		data.Set("sort_by_direction", fmt.Sprint(*params.SortByDirection))
	}
	if params != nil && params.Limit != nil {
		data.Set("limit", fmt.Sprint(*params.Limit))
	}
	if params != nil && params.Offset != nil {
		data.Set("offset", fmt.Sprint(*params.Offset))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &SubuserStats{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
