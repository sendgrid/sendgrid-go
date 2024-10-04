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

// UpdateTemplateRequest struct for UpdateTemplateRequest
type UpdateTemplateRequest struct {
	// Indicates if you want to enable the legacy email template mail setting.
	Enabled *bool `json:"enabled,omitempty"`
	// The new HTML content for your legacy email template.
	HtmlContent *string `json:"html_content,omitempty"`
}
