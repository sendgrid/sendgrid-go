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

// SendMailRequestTrackingSettings Settings to determine how you would like to track the metrics of how your recipients interact with your email. See [**Tracking Settings**](https://docs.sendgrid.com/ui/account-and-settings/tracking) for more information.
type SendMailRequestTrackingSettings struct {
	ClickTracking        *SendMailRequestTrackingSettingsClickTracking        `json:"click_tracking,omitempty"`
	OpenTracking         *SendMailRequestTrackingSettingsOpenTracking         `json:"open_tracking,omitempty"`
	SubscriptionTracking *SendMailRequestTrackingSettingsSubscriptionTracking `json:"subscription_tracking,omitempty"`
	Ganalytics           *SendMailRequestTrackingSettingsGanalytics           `json:"ganalytics,omitempty"`
}