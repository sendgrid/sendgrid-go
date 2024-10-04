# EventWebhookAllResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAllowed** | **float32** | The maximum number of Event Webhooks you can have enabled under your current Twilio SendGrid plan. See the [Twilio SendGrid pricing page](https://sendgrid.com/pricing) for more information about the features available with each plan. |[optional] 
**Webhooks** | [**[]EventWebhookSignedResponse**](EventWebhookSignedResponse.md) | An array of Event Webhook objects. Each object represents one of your webhooks and contains its configuration settings, including which events it is set to send in the POST request, the URL where it will send those events, and the webhook's ID. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


