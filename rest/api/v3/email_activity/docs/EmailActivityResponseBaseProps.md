# EmailActivityResponseBaseProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromEmail** | **string** | The 'From' email address used to deliver the message. This address should be a verified sender in your Twilio SendGrid account. |[optional] [default to "test0@example.com"]
**MsgId** | **string** | A unique ID assigned to the message. This ID can be used to retrieve activity data for the specific message. |[optional] 
**Subject** | **string** | The email's subject line. |[optional] 
**ToEmail** | **string** | The intended recipient's email address. |[optional] 
**Status** | [**Status3**](Status3.md) | The message's status. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


