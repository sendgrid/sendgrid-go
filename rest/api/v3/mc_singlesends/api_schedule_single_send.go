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
	"net/url"

	"strings"
)

type ScheduleSingleSendParam struct {
	//
	Id *string `json:"id"`
	//
	ScheduleSingleSendRequest *ScheduleSingleSendRequest `json:"ScheduleSingleSendRequest,omitempty"`
}

func (params *ScheduleSingleSendParam) SetId(Id string) *ScheduleSingleSendParam {
	params.Id = &Id
	return params
}
func (params *ScheduleSingleSendParam) SetScheduleSingleSendRequest(ScheduleSingleSendRequest ScheduleSingleSendRequest) *ScheduleSingleSendParam {
	params.ScheduleSingleSendRequest = &ScheduleSingleSendRequest
	return params
}

// **This endpoint allows you to send a Single Send immediately or schedule it to be sent at a later time.**  To send your message immediately, set the `send_at` property value to the string `now`. To schedule the Single Send for future delivery, set the `send_at` value to your desired send time in [ISO 8601 date time format](https://www.iso.org/iso-8601-date-and-time-format.html) (`yyyy-MM-ddTHH:mm:ssZ`).
func (c *ApiService) ScheduleSingleSend(params *ScheduleSingleSendParam) (interface{}, error) {
	path := "/v3/marketing/singlesends/{Id}/schedule"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.ScheduleSingleSendRequest != nil {
		b, err := json.Marshal(*params.ScheduleSingleSendRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Put(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &ScheduleSingleSend201Response{}
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
	return nil, nil
}
