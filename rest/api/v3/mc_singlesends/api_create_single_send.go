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
)

type CreateSingleSendParam struct {
	//
	SinglesendRequest *SinglesendRequest `json:"SinglesendRequest,omitempty"`
}

func (params *CreateSingleSendParam) SetSinglesendRequest(SinglesendRequest SinglesendRequest) *CreateSingleSendParam {
	params.SinglesendRequest = &SinglesendRequest
	return params
}

// **This endpoint allows you to create a new Single Send.**  Please note that if you are migrating from the previous version of Single Sends, you no longer need to pass a template ID with your request to this endpoint. Instead, you will pass all template data in the `email_config` object.  This endpoint will create a draft of the Single Send but will not send it or schedule it to be sent. Any `send_at` property value set with this endpoint will prepopulate the Single Send's send date. However, the Single Send will remain an unscheduled draft until it's updated with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI.
func (c *ApiService) CreateSingleSend(params *CreateSingleSendParam) (interface{}, error) {
	path := "/v3/marketing/singlesends"

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

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
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
	if resp.StatusCode == 500 {
		ps := &ListSingleSend500Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
