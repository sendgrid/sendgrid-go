# EventWebhookNoDatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if the Event Webhook is enabled. |[optional] 
**Url** | **string** | The URL where SendGrid will send event data. |[optional] 
**AccountStatusChange** | **bool** | Indicates if the webhook is configured to send account status change events related to compliance action taken by SendGrid. |[optional] 
**GroupResubscribe** | **bool** | Indicates if the webhook is configured to send group resubscribe events. Group resubscribes occur when recipients resubscribe to a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**Delivered** | **bool** | Indicates if the webhook is configured to send delivered events. Delivered events occur when a message has been successfully delivered to the receiving server. |[optional] 
**GroupUnsubscribe** | **bool** | Indicates if the webhook is configured to send group unsubscribe events. Group unsubscribes occur when recipients unsubscribe from a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) either by direct link or by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**SpamReport** | **bool** | Indicates if the webhook is configured to send spam report events. Spam reports occur when recipients mark a message as spam. |[optional] 
**Bounce** | **bool** | Indicates if the webhook is configured to send bounce events. A bounce occurs when a receiving server could not or would not accept a message. |[optional] 
**Deferred** | **bool** | Indicates if the webhook is configured to send deferred events. Deferred events occur when a recipient's email server temporarily rejects a message. |[optional] 
**Unsubscribe** | **bool** | Indicates if the webhook is configured to send unsubscribe events. Unsubscribes occur when recipients click on a message's subscription management link. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**Processed** | **bool** | Indicates if the webhook is configured to send processed events. Processed events occur when a message has been received by Twilio SendGrid and is ready to be delivered. |[optional] 
**Open** | **bool** | Indicates if the webhook is configured to send open events. Open events occur when a recipient has opened the HTML message. You must [enable Open Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#open-tracking) to receive this type of event. |[optional] 
**Click** | **bool** | Indicates if the webhook is configured to send click events. Click events occur when a recipient clicks on a link within the message. You must [enable Click Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#click-tracking) to receive this type of event. |[optional] 
**Dropped** | **bool** | Indicates if the webhook is configured to send dropped events. Dropped events occur when your message is not delivered by Twilio SendGrid. Dropped events are accomponied by a `reason` property, which indicates why the message was dropped. Reasons for a dropped message include: Invalid SMTPAPI header, Spam Content (if spam checker app enabled), Unsubscribed Address, Bounced Address, Spam Reporting Address, Invalid, Recipient List over Package Quota. |[optional] 
**FriendlyName** | **string** | An optional friendly name assigned to the Event Webhook to help you differentiate it. The friendly name is for convenience only. You should use the webhook `id` property for any programmatic tasks. |[optional] 
**Id** | **string** | A unique string used to identify the webhook. A webhook's ID is generated programmatically and cannot be changed after creation. You can assign a natural language identifier to your webhook using the `friendly_name` property. |[optional] 
**OauthClientId** | **string** | The OAuth client ID SendGrid sends to your OAuth server or service provider to generate an OAuth access token. |[optional] 
**OauthTokenUrl** | **string** | The URL where SendGrid sends the OAuth client ID and client secret to generate an access token. This should be your OAuth server or service provider. |[optional] 
**PublicKey** | **string** | The public key you can use to verify the SendGrid signature. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


