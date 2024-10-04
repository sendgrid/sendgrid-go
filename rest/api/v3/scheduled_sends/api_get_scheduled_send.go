/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Scheduled Sends API
* The Twilio SendGrid Scheduled Sends API allows you to cancel or pause the send of one or more emails using a batch ID.  A `batch_id` groups multiple scheduled mail send requests together with the same ID. You can cancel or pause all of the requests associated with a batch ID up to 10 minutes before the scheduled send time.  See [**Canceling a Scheduled Send**](https://docs.sendgrid.com/for-developers/sending-email/stopping-a-scheduled-send) for a guide on creating a `batch_id`, assigning it to a scheduled send, and modifying the send.  See the Mail API's batching operations to validate a `batch_id` and retrieve all scheduled sends as an array.  When a batch is canceled, all messages associated with that batch will stay in your sending queue. When their `send_at` value is reached, they will be discarded.  When a batch is paused, all messages associated with that batch will stay in your sending queue, even after their `send_at` value has passed. This means you can remove a pause status, and your scheduled send will be delivered once the pause is removed. Any messages left with a pause status that are more than 72 hours old will be discarded as Expired.
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

type GetScheduledSendParam struct {
	//
	BatchId *string `json:"batch_id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *GetScheduledSendParam) SetBatchId(BatchId string) *GetScheduledSendParam {
	params.BatchId = &BatchId
	return params
}
func (params *GetScheduledSendParam) SetOnbehalfof(Onbehalfof string) *GetScheduledSendParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve the cancel/paused scheduled send information for a specific `batch_id`.**
func (c *ApiService) GetScheduledSend(params *GetScheduledSendParam) (interface{}, error) {
	path := "/v3/user/scheduled_sends/{BatchId}"
	if params != nil && params.BatchId != nil {
		path = strings.Replace(path, "{"+"BatchId"+"}", *params.BatchId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &[]ScheduledSendStatus{}
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
	if resp.StatusCode == 401 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 403 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
