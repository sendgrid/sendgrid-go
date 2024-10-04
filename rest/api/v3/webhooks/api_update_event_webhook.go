/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Webhook Configuration API
* The Twilio SendGrid Webhooks API allows you to configure the Event and Parse Webhooks.  Email is a data-rich channel, and implementing the Event Webhook will allow you to consume those data in nearly real time. This means you can actively monitor and participate in the health of your email program throughout the send cycle.  The Inbound Parse Webhook processes all incoming email for a domain or subdomain, parses the contents and attachments and then POSTs `multipart/form-data` to a URL that you choose.
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

type UpdateEventWebhookParam struct {
	// The ID of the Event Webhook you want to retrieve.
	Id *string `json:"id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	// Use this to include optional fields in the response payload. When this is set to `include=account_status_change`, the `account_status_change` field will be part of the response payload and set to `false` by default. See [Update an event webhook](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) for enabling this webhook notification which lets you subscribe to account status change events related to compliance action taken by SendGrid.
	Include *string `json:"include,omitempty"`
	//
	EventWebhookRequest *EventWebhookRequest `json:"EventWebhookRequest,omitempty"`
}

func (params *UpdateEventWebhookParam) SetId(Id string) *UpdateEventWebhookParam {
	params.Id = &Id
	return params
}
func (params *UpdateEventWebhookParam) SetOnbehalfof(Onbehalfof string) *UpdateEventWebhookParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateEventWebhookParam) SetInclude(Include string) *UpdateEventWebhookParam {
	params.Include = &Include
	return params
}
func (params *UpdateEventWebhookParam) SetEventWebhookRequest(EventWebhookRequest EventWebhookRequest) *UpdateEventWebhookParam {
	params.EventWebhookRequest = &EventWebhookRequest
	return params
}

// **This endpoint allows you to update a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will update and return your oldest webhook by `created_date`. This means the default webhook updated by this endpoint when no ID is provided will be the first one you created. This functionality allows customers who do not have multiple webhooks to use this endpoint to update their only webhook, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  ### Enable or disable the webhook  You can set the `enabled` property to `true` to enable the webhook or `false` to disable it. Disabling a webhook will not delete it from your account, but it will prevent the webhook from sending events to your designated URL.  ### URL  A webhook's URL is the endpoint where you want the webhook to send POST requests containing event data. No more than one webhook may be configured to send to the same URL. SendGrid will return an error if you attempt to set a URL for a webhook that is already in use by the user on another webhook.  ### Event settings  If an event type is marked as `true`, the event webhook will send information about that event type. See the [**Event Webhook Reference**](https://docs.sendgrid.com/for-developers/tracking-events/event#delivery-events) for details about each event type.  ### Webhook identifiers  You may assign an optional friendly name to each of your webhooks. The friendly name is for convenience only and should not be used to programmatically differentiate your webhooks because it does not need to be unique.  ### OAuth  You can configure OAuth for your webhook by passing the required values to this endpoint in the `oauth_client_id`, `oauth_client_secret`, and `oauth_token_url` properties. To disable OAuth, pass an empty string to this endpoint for each of the OAuth properties. You may share one OAuth configuration across all your webhooks or create unique credentials for each. See our [webhook security documentation](https://docs.sendgrid.com/for-developers/tracking-events/getting-started-event-webhook-security-features#oauth-20) for more detailed information about OAuth and the Event Webhook.  ### Signature verification  Enabling signature verification for your webhook is a separate process and cannot be done with this endpoint. You can use the webhook ID to [enable or disable signature verification with the endpoint dedicated for that operation](https://docs.sendgrid.com/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook).
func (c *ApiService) UpdateEventWebhook(params *UpdateEventWebhookParam) (interface{}, error) {
	path := "/v3/user/webhooks/event/settings/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	if params != nil && params.Include != nil {
		data.Set("include", *params.Include)
	}
	body := []byte{}
	if params != nil && params.EventWebhookRequest != nil {
		b, err := json.Marshal(*params.EventWebhookRequest)
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
		ps := &EventWebhookUnsignedResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &CreateEventWebhook400Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &GetSignedEventWebhook404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
