# CreateAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Type**](Type.md) | The type of alert you want to create. Can be either usage_limit or stats_notification. Example: usage_limit |
**EmailTo** | **string** | The email address the alert will be sent to. Example: test@example.com |
**Frequency** | **string** | Required for stats_notification. How frequently the alert will be sent. Example: daily |[optional] 
**Percentage** | **int32** | Required for usage_limit. When this usage threshold is reached, the alert will be sent. Example: 90 |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


