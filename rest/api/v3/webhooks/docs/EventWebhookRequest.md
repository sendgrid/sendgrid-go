# EventWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set this property to `true` to enable the Event Webhook or `false` to disable it. |[optional] 
**Url** | **string** | Set this property to the URL where you want the Event Webhook to send event data. |
**GroupResubscribe** | **bool** | Set this property to `true` to receive group resubscribe events. Group resubscribes occur when recipients resubscribe to a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**Delivered** | **bool** | Set this property to `true` to receive delivered events. Delivered events occur when a message has been successfully delivered to the receiving server. |[optional] 
**GroupUnsubscribe** | **bool** | Set this property to `true` to receive group unsubscribe events. Group unsubscribes occur when recipients unsubscribe from a specific [unsubscribe group](https://docs.sendgrid.com/ui/sending-email/create-and-manage-unsubscribe-groups) either by direct link or by updating their subscription preferences. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**SpamReport** | **bool** | Set this property to `true` to receive spam report events. Spam reports occur when recipients mark a message as spam. |[optional] 
**Bounce** | **bool** | Set this property to `true` to receive bounce events. A bounce occurs when a receiving server could not or would not accept a message. |[optional] 
**Deferred** | **bool** | Set this property to `true` to receive deferred events. Deferred events occur when a recipient's email server temporarily rejects a message. |[optional] 
**Unsubscribe** | **bool** | Set this property to `true` to receive unsubscribe events. Unsubscribes occur when recipients click on a message's subscription management link. You must [enable Subscription Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#subscription-tracking) to receive this type of event. |[optional] 
**Processed** | **bool** | Set this property to `true` to receive processed events. Processed events occur when a message has been received by Twilio SendGrid and the message is ready to be delivered. |[optional] 
**Open** | **bool** | Set this property to `true` to receive open events. Open events occur when a recipient has opened the HTML message. You must [enable Open Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#open-tracking) to receive this type of event. |[optional] 
**Click** | **bool** | Set this property to `true` to receive click events. Click events occur when a recipient clicks on a link within the message. You must [enable Click Tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking#click-tracking) to receive this type of event. |[optional] 
**Dropped** | **bool** | Set this property to `true` to receive dropped events. Dropped events occur when your message is not delivered by Twilio SendGrid. Dropped events are accomponied by a `reason` property, which indicates why the message was dropped. Reasons for a dropped message include: Invalid SMTPAPI header, Spam Content (if spam checker app enabled), Unsubscribed Address, Bounced Address, Spam Reporting Address, Invalid, Recipient List over Package Quota. |[optional] 
**FriendlyName** | **string** | Optionally set this property to a friendly name for the Event Webhook. A friendly name may be assigned to each of your webhooks to help you differentiate them. The friendly name is for convenience only. You should use the webhook `id` property for any programmatic tasks. |[optional] 
**OauthClientId** | **string** | Set this property to the OAuth client ID that SendGrid will pass to your OAuth server or service provider to generate an OAuth access token. When passing data in this property, you must also include the `oauth_token_url` property. |[optional] 
**OauthClientSecret** | **string** | Set this property to the OAuth client secret that SendGrid will pass to your OAuth server or service provider to generate an OAuth access token. This secret is needed only once to create an access token. SendGrid will store the secret, allowing you to update your client ID and Token URL without passing the secret to SendGrid again. When passing data in this field, you must also include the `oauth_client_id` and `oauth_token_url` properties. |[optional] 
**OauthTokenUrl** | **string** | Set this property to the URL where SendGrid will send the OAuth client ID and client secret to generate an OAuth access token. This should be your OAuth server or service provider. When passing data in this field, you must also include the `oauth_client_id` property. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


