# UpdateAlert200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int32** | A Unix timestamp indicating when the alert was created. |
**EmailTo** | **string** | The email address that the alert will be sent to. |
**Frequency** | **string** | If the alert is of type stats_notification, this indicates how frequently the stats notifications will be sent. For example: \"daily\", \"weekly\", or \"monthly\". |[optional] 
**Id** | **int32** | The ID of the alert. |
**Type** | [**Type3**](Type3.md) | The type of alert. |
**UpdatedAt** | **int32** | A Unix timestamp indicating when the alert was last modified. |
**Percentage** | **int32** | If the alert is of type usage_limit, this indicates the percentage of email usage that must be reached before the alert will be sent. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


