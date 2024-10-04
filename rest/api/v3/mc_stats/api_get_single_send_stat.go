/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Statistics API
* The Marketing Campaigns Stats API allows you to retrieve statistics for both Automations and Single Sends. The statistics provided include bounces, clicks, opens, and more. You can export stats in CSV format for use in other applications. You can also retrieve Marketing Campaigns stats in the [Marketing Campaigns application user interface](https://mc.sendgrid.com/).  This API provides statistics for Marketing Campaigns only. For stats related to event tracking, please see the [Stats API](https://docs.sendgrid.com/api-reference/stats).
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

type GetSingleSendStatParam struct {
	// The ID of Single Send for which you want to retrieve stats.
	Id *string `json:"id"`
	// Dictates how the stats are time-sliced. Currently, `\"total\"` and `\"day\"` are supported.
	AggregatedBy *AggregatedBy `json:"aggregated_by,omitempty"`
	// Format: `YYYY-MM-DD`. If this parameter is included, the stats' start date is included in the search.
	StartDate *string `json:"start_date,omitempty"`
	// Format: `YYYY-MM-DD`.If this parameter is included, the stats' end date is included in the search.
	EndDate *string `json:"end_date,omitempty"`
	// [IANA Area/Region](https://en.wikipedia.org/wiki/Tz_database#Names_of_timezones) string representing the timezone in which the stats are to be presented, e.g., \"America/Chicago\".
	Timezone *string `json:"timezone,omitempty"`
	// The number of elements you want returned on each page.
	PageSize *int32 `json:"page_size,omitempty"`
	// The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.
	PageToken *string `json:"page_token,omitempty"`
	// A/B Single Sends have multiple variation IDs and phase IDs. Including these additional fields allows further granularity of stats by these fields.
	GroupBy *[]Items1 `json:"group_by,omitempty"`
}

func (params *GetSingleSendStatParam) SetId(Id string) *GetSingleSendStatParam {
	params.Id = &Id
	return params
}
func (params *GetSingleSendStatParam) SetAggregatedBy(AggregatedBy AggregatedBy) *GetSingleSendStatParam {
	params.AggregatedBy = &AggregatedBy
	return params
}
func (params *GetSingleSendStatParam) SetStartDate(StartDate string) *GetSingleSendStatParam {
	params.StartDate = &StartDate
	return params
}
func (params *GetSingleSendStatParam) SetEndDate(EndDate string) *GetSingleSendStatParam {
	params.EndDate = &EndDate
	return params
}
func (params *GetSingleSendStatParam) SetTimezone(Timezone string) *GetSingleSendStatParam {
	params.Timezone = &Timezone
	return params
}
func (params *GetSingleSendStatParam) SetPageSize(PageSize int32) *GetSingleSendStatParam {
	params.PageSize = &PageSize
	return params
}
func (params *GetSingleSendStatParam) SetPageToken(PageToken string) *GetSingleSendStatParam {
	params.PageToken = &PageToken
	return params
}
func (params *GetSingleSendStatParam) SetGroupBy(GroupBy []Items1) *GetSingleSendStatParam {
	params.GroupBy = &GroupBy
	return params
}

// **This endpoint allows you to retrieve stats for an individual Single Send using a Single Send ID.**  Multiple Single Send IDs can be retrieved using the \"Get All Single Sends Stats\" endpoint. Once you have an ID, this endpoint will return detailed stats for the Single Send specified.  You may constrain the stats returned using the `start_date` and `end_date` query string parameters. You can also use the `group_by` and `aggregated_by` query string parameters to further refine the stats returned.
func (c *ApiService) GetSingleSendStat(params *GetSingleSendStatParam) (interface{}, error) {
	path := "/v3/marketing/stats/singlesends/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AggregatedBy != nil {
		data.Set("aggregated_by", fmt.Sprint(*params.AggregatedBy))
	}
	if params != nil && params.StartDate != nil {
		data.Set("start_date", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("end_date", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.Timezone != nil {
		data.Set("timezone", *params.Timezone)
	}
	if params != nil && params.PageSize != nil {
		data.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if params != nil && params.PageToken != nil {
		data.Set("page_token", *params.PageToken)
	}
	if params != nil && params.GroupBy != nil {
		for _, item := range *params.GroupBy {
			v, err := json.Marshal(item)

			if err != nil {
				return nil, err
			}

			data.Add("group_by", string(v))

		}
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &SinglesendsResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
