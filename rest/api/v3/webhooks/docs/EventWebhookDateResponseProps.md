# EventWebhookDateResponseProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | [**time.Time**](time.Time.md) | An ISO 8601 timestamp in UTC timezone when the Event Webhook was created. If a Webhook's `created_date` is `null`, it is a [legacy Event Webook](https://www.twilio.com/en-us/changelog/event-webhooks), which means it is your oldest Webhook. |[optional] 
**UpdatedDate** | [**time.Time**](time.Time.md) | An ISO 8601 timestamp in UTC timezone when the Event Webhook was last modified. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


