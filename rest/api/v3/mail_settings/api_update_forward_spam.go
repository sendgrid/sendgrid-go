/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Mail Settings API
* The Twilio SendGrid Mail Settings API allows you to retrieve and configure the various Mail Settings available. Mail Settings instruct SendGrid to apply specific settings to every email that you send over [SendGrid’s Web API](https://docs.sendgrid.com/api-reference/mail-send/v3-mail-send) or [SMTP relay](https://sendgrid.com/docs/for-developers/sending-email/building-an-x-smtpapi-header/). These settings include how to embed a custom footer, how to manage blocks, spam, and bounces, and more.  For a full list of Twilio SendGrid's Mail Settings, and what each one does, see [**Mail Settings**](https://sendgrid.com/docs/ui/account-and-settings/mail/).  You can also manage your Mail Settings in the Twilio SendGrid application user interface.
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

type UpdateForwardSpamParam struct {
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
	//
	MailSettingsForwardSpam *MailSettingsForwardSpam `json:"MailSettingsForwardSpam,omitempty"`
}

func (params *UpdateForwardSpamParam) SetOnbehalfof(Onbehalfof string) *UpdateForwardSpamParam {
	params.Onbehalfof = &Onbehalfof
	return params
}
func (params *UpdateForwardSpamParam) SetMailSettingsForwardSpam(MailSettingsForwardSpam MailSettingsForwardSpam) *UpdateForwardSpamParam {
	params.MailSettingsForwardSpam = &MailSettingsForwardSpam
	return params
}

// **This endpoint allows you to update your current Forward Spam mail settings.**  Enabling the Forward Spam setting allows you to specify `email` addresses to which spam reports will be forwarded. You can set multiple addresses by passing this endpoint a comma separated list of emails in a single string.  ``` {   \"email\": \"address1@example.com, address2@exapmle.com\",   \"enabled\": true } ```  The Forward Spam setting may also be used to receive emails sent to `abuse@` and `postmaster@` role addresses if you have authenticated your domain.  For example, if you authenticated `example.com` as your root domain and set a custom return path of `sub` for that domain, you could turn on Forward Spam, and any emails sent to `abuse@sub.example.com` or `postmaster@sub.example.com` would be forwarded to the email address you entered in the `email` field.  You can authenticate your domain using the \"Authenticate a domain\" endpoint or in the [Sender Authentication section of the Twilio SendGrid App](https://app.sendgrid.com/settings/sender_auth). You can also configure the Forward Spam mail settings in the [Mail Settings section of the Twilio SendGrid App](https://app.sendgrid.com/settings/mail_settings).
func (c *ApiService) UpdateForwardSpam(params *UpdateForwardSpamParam) (interface{}, error) {
	path := "/v3/mail_settings/forward_spam"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.MailSettingsForwardSpam != nil {
		b, err := json.Marshal(*params.MailSettingsForwardSpam)
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
		ps := &MailSettingsForwardSpam{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}