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

// SendMailRequestMailSettings A collection of different mail settings that you can use to specify how you would like this email to be handled. Mail settings provide extra functionality to your mail send. See [**Mail Settings**](https://docs.sendgrid.com/ui/account-and-settings/mail) for more information.
type SendMailRequestMailSettings struct {
	BypassListManagement        *SendMailRequestMailSettingsBypassListManagement        `json:"bypass_list_management,omitempty"`
	BypassSpamManagement        *SendMailRequestMailSettingsBypassSpamManagement        `json:"bypass_spam_management,omitempty"`
	BypassBounceManagement      *SendMailRequestMailSettingsBypassBounceManagement      `json:"bypass_bounce_management,omitempty"`
	BypassUnsubscribeManagement *SendMailRequestMailSettingsBypassUnsubscribeManagement `json:"bypass_unsubscribe_management,omitempty"`
	Footer                      *SendMailRequestMailSettingsFooter                      `json:"footer,omitempty"`
	SandboxMode                 *SendMailRequestMailSettingsSandboxMode                 `json:"sandbox_mode,omitempty"`
}
