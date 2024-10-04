# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromEmail** | **string** | The 'From' email address used to deliver the message. This address should be a verified sender in your Twilio SendGrid account. |[optional] [default to "test0@example.com"]
**MsgId** | **string** | A unique ID assigned to the message. This ID can be used to retrieve activity data for the specific message. |[optional] 
**Subject** | **string** | The email's subject line. |[optional] 
**ToEmail** | **string** | The intended recipient's email address. |[optional] 
**Status** | [**Status3**](Status3.md) | The message's status. |[optional] 
**TemplateId** | **string** | The ID associated with a Twilio SendGrid email template used to format the message. |
**AsmGroupId** | **int32** | The unsubscribe group associated with this email. |
**Teammate** | **string** | Teammate's username |
**ApiKeyId** | **string** | The ID of the API Key used to authenticate the sending request for the message. |
**Events** | [**[]Event**](Event.md) | List of events related to email message |
**OriginatingIp** | **string** | This is the IP of the user who sent the message. |
**Categories** | **[]string** | Categories users associated to the message |
**UniqueArgs** | **string** | JSON hash of key-value pairs associated with the message. |[default to "Null"]
**OutboundIp** | **string** | IP used to send to the remote Mail Transfer Agent. |
**OutboundIpType** | [**OutboundIpType**](OutboundIpType.md) | Whether or not the outbound IP is dedicated vs shared |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


