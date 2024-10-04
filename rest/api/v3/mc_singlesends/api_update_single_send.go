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

type UpdateSingleSendParam struct {
	//
	Id *string `json:"id"`
	//
	SinglesendRequest *SinglesendRequest `json:"SinglesendRequest,omitempty"`
}

func (params *UpdateSingleSendParam) SetId(Id string) *UpdateSingleSendParam {
	params.Id = &Id
	return params
}
func (params *UpdateSingleSendParam) SetSinglesendRequest(SinglesendRequest SinglesendRequest) *UpdateSingleSendParam {
	params.SinglesendRequest = &SinglesendRequest
	return params
}

// **This endpoint allows you to update a Single Send using a Single Send ID.**  You only need to pass the properties you want to update. Any blank or missing properties will remain unaltered.  This endpoint will update a draft of the Single Send but will not send it or schedule it to be sent. Any `send_at` property value set with this endpoint will prepopulate the Single Send's send date. However, the Single Send will remain an unscheduled draft until it's updated with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI.
func (c *ApiService) UpdateSingleSend(params *UpdateSingleSendParam) (interface{}, error) {
	path := "/v3/marketing/singlesends/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.SinglesendRequest != nil {
		b, err := json.Marshal(*params.SinglesendRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 202 {
		ps := &SinglesendResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ListSingleSend500Response{}
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
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
