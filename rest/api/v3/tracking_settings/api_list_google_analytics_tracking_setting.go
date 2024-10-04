/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Tracking Settings API
* The Twilio SendGrid Tracking Settings API allows you to configure which tracking settings are enabled for your messages. You can track many of your recipients' interactions with your emails, including which emails they open, which links they click, and when they subscribe to or unsubscribe from your emails. See [**Tracking Settings**](https://docs.sendgrid.com/ui/account-and-settings/tracking) for more information.
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

type ListGoogleAnalyticsTrackingSettingParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *ListGoogleAnalyticsTrackingSettingParam) SetOnbehalfof(Onbehalfof string) *ListGoogleAnalyticsTrackingSettingParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to retrieve your current setting for Google Analytics.**   Google Analytics helps you understand how users got to your site and what they're doing there. For more information about using Google Analytics, please refer to [Google’s URL Builder](https://support.google.com/analytics/answer/1033867?hl=en) and their article on [\"Best Practices for Campaign Building\"](https://support.google.com/analytics/answer/1037445).  We default the settings to Google’s recommendations. For more information, see [Google Analytics Demystified](https://sendgrid.com/docs/ui/analytics-and-reporting/google-analytics/).
func (c *ApiService) ListGoogleAnalyticsTrackingSetting(params *ListGoogleAnalyticsTrackingSettingParam) (interface{}, error) {
	path := "/v3/tracking_settings/google_analytics"

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
		ps := &GoogleAnalyticsSettings{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
