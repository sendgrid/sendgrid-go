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

// EventWebhookRequest struct for EventWebhookRequest
type EventWebhookRequest struct {
	// Set this property to `true` to enable the Event Webhook or `false` to disable it.
	Enabled *bool `json:"enabled,omitempty"`
	// Set this property to the URL where you want the Event Webhook to send event data.
	Url string `json:"url"`
	// Set this property to `true` to receive group resubscribe events. Group resubscribes occur when recipients resubscribe to a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	GroupResubscribe *bool `json:"group_resubscribe,omitempty"`
	// Set this property to `true` to receive delivered events. Delivered events occur when a message has been successfully delivered to the receiving server.
	Delivered *bool `json:"delivered,omitempty"`
	// Set this property to `true` to receive group unsubscribe events. Group unsubscribes occur when recipients unsubscribe from a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) either by direct link or by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	GroupUnsubscribe *bool `json:"group_unsubscribe,omitempty"`
	// Set this property to `true` to receive spam report events. Spam reports occur when recipients mark a message as spam.
	SpamReport *bool `json:"spam_report,omitempty"`
	// Set this property to `true` to receive bounce events. A bounce occurs when a receiving server could not or would not accept a message.
	Bounce *bool `json:"bounce,omitempty"`
	// Set this property to `true` to receive deferred events. Deferred events occur when a recipient's email server temporarily rejects a message.
	Deferred *bool `json:"deferred,omitempty"`
	// Set this property to `true` to receive unsubscribe events. Unsubscribes occur when recipients click on a message's subscription management link. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event.
	Unsubscribe *bool `json:"unsubscribe,omitempty"`
	// Set this property to `true` to receive processed events. Processed events occur when a message has been received by Twilio SendGrid and the message is ready to be delivered.
	Processed *bool `json:"processed,omitempty"`
	// Set this property to `true` to receive open events. Open events occur when a recipient has opened the HTML message. You must [enable Open Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#open-tracking) to receive this type of event.
	Open *bool `json:"open,omitempty"`
	// Set this property to `true` to receive click events. Click events occur when a recipient clicks on a link within the message. You must [enable Click Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#click-tracking) to receive this type of event.
	Click *bool `json:"click,omitempty"`
	// Set this property to `true` to receive dropped events. Dropped events occur when your message is not delivered by Twilio SendGrid. Dropped events are accomponied by a `reason` property, which indicates why the message was dropped. Reasons for a dropped message include: Invalid SMTPAPI header, Spam Content (if spam checker app enabled), Unsubscribed Address, Bounced Address, Spam Reporting Address, Invalid, Recipient List over Package Quota.
	Dropped *bool `json:"dropped,omitempty"`
	// Optionally set this property to a friendly name for the Event Webhook. A friendly name may be assigned to each of your webhooks to help you differentiate them. The friendly name is for convenience only. You should use the webhook `id` property for any programmatic tasks.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Set this property to the OAuth client ID that SendGrid will pass to your OAuth server or service provider to generate an OAuth access token. When passing data in this property, you must also include the `oauth_token_url` property.
	OauthClientId *string `json:"oauth_client_id,omitempty"`
	// Set this property to the OAuth client secret that SendGrid will pass to your OAuth server or service provider to generate an OAuth access token. This secret is needed only once to create an access token. SendGrid will store the secret, allowing you to update your client ID and Token URL without passing the secret to SendGrid again. When passing data in this field, you must also include the `oauth_client_id` and `oauth_token_url` properties.
	OauthClientSecret *string `json:"oauth_client_secret,omitempty"`
	// Set this property to the URL where SendGrid will send the OAuth client ID and client secret to generate an OAuth access token. This should be your OAuth server or service provider. When passing data in this field, you must also include the `oauth_client_id` property.
	OauthTokenUrl *string `json:"oauth_token_url,omitempty"`
}
