# SinglesendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the Single Send. |[optional] 
**Name** | **string** | The name of the Single Send. |[optional] 
**Status** | [**Status2**](Status2.md) | The current status of the Single Send. The status may be `draft`, `scheduled`, or `triggered`. |[optional] 
**Categories** | **[]string** | The categories associated with this Single Send. |[optional] 
**SendAt** | [**time.Time**](time.Time.md) | An ISO 8601 formatted date-time when the Single Send is set to be sent. Please note that any `send_at` property value will have no effect while the Single Send `status` is `draft`. You must update the Single Send with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI to schedule it. |[optional] 
**SendTo** | [**SinglesendResponseSendTo**](SinglesendResponseSendTo.md) |  |[optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | the ISO 8601 time at which the Single Send was last updated. |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | the ISO 8601 time at which the Single Send was created. |[optional] 
**EmailConfig** | [**SinglesendResponseEmailConfig**](SinglesendResponseEmailConfig.md) |  |[optional] 
**Warnings** | [**[]SinglesendResponseWarningsInner**](SinglesendResponseWarningsInner.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


