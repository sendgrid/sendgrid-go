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

// SinglesendResponseEmailConfig struct for SinglesendResponseEmailConfig
type SinglesendResponseEmailConfig struct {
	// The subject line of the Single Send. This property is not used when a `design_id` value is set.
	Subject *string `json:"subject,omitempty"`
	// The HTML content of the Single Send. This property is not used when a `design_id` value is set.
	HtmlContent *string `json:"html_content,omitempty"`
	// The plain text content of the Single Send. This property is not used when a `design_id` value is set.
	PlainContent *string `json:"plain_content,omitempty"`
	// If this property is set to `true`, `plain_content` is always generated from `html_content`. If it's set to false, `plain_content` is not altered.
	GeneratePlainContent *bool `json:"generate_plain_content,omitempty"`
	// A `design_id` can be used in place of `html_content`, `plain_content`, and/or `subject`. You can retrieve a design's ID from the [**List Designs** endpoint](https://docs.sendgrid.com/api-reference/designs-api/list-designs) or by pulling it from the design's detail page URL in the Marketing Campaigns App.
	DesignId *string `json:"design_id,omitempty"`
	// The editor, `design` or `code`, used to modify the Single Send's design in the Marketing Campaigns application user interface.
	Editor *Editor1 `json:"editor,omitempty"`
	// The ID of the Suppression Group to allow recipients to unsubscribe. You must provide a `suppression_group_id` or the `custom_unsubscribe_url` with your Single Send.
	SuppressionGroupId *int32 `json:"suppression_group_id,omitempty"`
	// The URL allowing recipients to unsubscribe. You must provide a `custom_unsubscribe_url` or a `suppression_group_id` with your Single Send.
	CustomUnsubscribeUrl *string `json:"custom_unsubscribe_url,omitempty"`
	// The ID of the verified sender from whom the Single Send will be sent. You can retrieve a verified sender's ID from the [**Get Verified Senders** endpoint](https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/get-all-verified-senders) or by pulling it from the sender's detail page URL in the SendGrid App.
	SenderId *int32 `json:"sender_id,omitempty"`
	// The name of the IP Pool from which the Single Send emails are sent.
	IpPool *string `json:"ip_pool,omitempty"`
}