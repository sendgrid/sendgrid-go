/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Single Sends API
* The Twilio SendGrid Single Sends API allows you to create, manage, and send Single Sends. You can also search Single Sends and retrieve statistics about them to help you maintain, alter, and further develop your campaigns.  A Single Send is a one-time non-automated email message delivered to a list or segment of your audience. A Single Send may be sent immediately or scheduled for future delivery.  Single Sends can serve many use cases, including promotional offers, engagement campaigns, newsletters, announcements, legal notices, or policy updates. You can also create and manage Single Sends in the [Marketing Campaigns application user interface](https://mc.sendgrid.com/single-sends).  The Single Sends API changed on May 6, 2020. See [**Single Sends 2020 Update**](https://docs.sendgrid.com/for-developers/sending-email/single-sends-2020-update) for more information.
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

type DeleteScheduledSingleSendParam struct {
	//
	Id *string `json:"id"`
}

func (params *DeleteScheduledSingleSendParam) SetId(Id string) *DeleteScheduledSingleSendParam {
	params.Id = &Id
	return params
}

// **This endpoint allows you to cancel a scheduled Single Send using a Single Send ID.**  Making a DELETE request to this endpoint will cancel the scheduled sending of a Single Send. The request will not delete the Single Send itself. Deleting a Single Send can be done by passing a DELETE request to `/marketing/singlesends/{id}`.
func (c *ApiService) DeleteScheduledSingleSend(params *DeleteScheduledSingleSendParam) (interface{}, error) {
	path := "/v3/marketing/singlesends/{Id}/schedule"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &SinglesendSchedule{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ListSingleSend500Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 500 {
		ps := &ListSingleSend500Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
