/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Mail API
* The Twilio SendGrid v3 Mail API allows you to send email at scale over HTTP. The Mail Send endpoint supports many levels of functionality, allowing you to send templates, set categories and custom arguments that can be used to analyze your send, and configure which tracking settings to include such as opens and clicks. You can also group mail sends into batches, allowing you to schedule and cancel sends by their batch IDs.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// SendMailRequestMailSettingsBypassSpamManagement Allows you to bypass the spam report list to ensure that the email is delivered to recipients. Bounce and unsubscribe lists will still be checked; addresses on these other lists will not receive the message. This filter cannot be combined with the `bypass_list_management` filter. See the [the \"Bypass suppressions\" section of our **Suppressions Overview**](https://sendgrid.com/docs/ui/sending-email/index-suppressions/#bypass-suppressions) for more about bypass filters.
type SendMailRequestMailSettingsBypassSpamManagement struct {
	// Indicates if this setting is enabled.
	Enable *bool `json:"enable,omitempty"`
}
