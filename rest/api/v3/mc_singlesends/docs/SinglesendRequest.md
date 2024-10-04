# SinglesendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Single Send. |
**Categories** | **[]string** | The categories to associate with this Single Send. |[optional] 
**SendAt** | [**time.Time**](time.Time.md) | Set this property to an ISO 8601 formatted date-time when you would like to send the Single Send. Please note that any `send_at` property value set with this endpoint will prepopulate the send date in the SendGrid user interface (UI). However, the Single Send will remain an unscheduled draft until it's updated with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI. Additionally, the `now` keyword is a valid `send_at` value only when using the Schedule Single Send endpoint. Setting this property to `now` with this endpoint will cause an error. |[optional] 
**SendTo** | [**SinglesendRequestSendTo**](SinglesendRequestSendTo.md) |  |[optional] 
**EmailConfig** | [**SinglesendRequestEmailConfig**](SinglesendRequestEmailConfig.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


