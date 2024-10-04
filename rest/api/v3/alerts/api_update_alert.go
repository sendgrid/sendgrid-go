/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Alerts API
* The Twilio SendGrid Alerts API allows you to specify an email address to receive notifications regarding your email usage or statistics. You can set up alerts to be sent to a specific email address on a recurring basis, whether for informational purposes or when specific account actions occur.  For most alerts, you can choose to have the alert sent to you as needed, hourly, daily, weekly, or monthly. The information contained in your alert will be for the last period of the alert. For example, if you choose weekly for the statistics alert, you will receive the statistics for the last week.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

type UpdateAlertParam struct {
	// The ID of the alert you would like to retrieve.
	AlertId *int32 `json:"alert_id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	UpdateAlertRequest *UpdateAlertRequest `json:"UpdateAlertRequest,omitempty"`
}

func (params *UpdateAlertParam) SetAlertId(AlertId int32) *UpdateAlertParam {
	params.AlertId = &AlertId
	return params
}
func (params *UpdateAlertParam) SetOnbehalfof(Onbehalfof string) *UpdateAlertParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateAlertParam) SetUpdateAlertRequest(UpdateAlertRequest UpdateAlertRequest) *UpdateAlertParam {
	params.UpdateAlertRequest = &UpdateAlertRequest
	return params
}

// **This endpoint allows you to update an alert.**
func (c *ApiService) UpdateAlert(params *UpdateAlertParam) (interface{}, error) {
	path := "/v3/alerts/{AlertId}"
	if params != nil && params.AlertId != nil {
		path = strings.Replace(path, "{"+"AlertId"+"}", fmt.Sprint(*params.AlertId), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateAlertRequest != nil {
		b, err := json.Marshal(*params.UpdateAlertRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &UpdateAlert200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
