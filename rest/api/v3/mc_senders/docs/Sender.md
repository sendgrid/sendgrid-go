# Sender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the Sender. |
**Nickname** | **string** | A nickname for the Sender. Not used for sending. |
**From** | [**CreateSenderRequestFrom**](CreateSenderRequestFrom.md) |  |
**ReplyTo** | [**CreateSenderRequestReplyTo**](CreateSenderRequestReplyTo.md) |  |
**Address** | **string** | The physical address of the Sender. |
**Address2** | **string** | Additional Sender address information. |[optional] 
**City** | **string** | The city of the Sender. |
**State** | **string** | The state of the Sender. |[optional] 
**Zip** | **string** | The zipcode of the Sender. |[optional] 
**Country** | **string** | The country of the Sender. |
**Verified** | **bool** | A boolean flag indicating whether the Sender is verified or not. Only verified Senders can be used to send email. |
**Locked** | **bool** | A boolean flag that is `true` when the Sender is associated with a campaign in Draft, Scheduled, or In Progress status. You cannot update or delete a locked Sender. |
**UpdatedAt** | **int32** | The time the Sender was last updated. |
**CreatedAt** | **int32** | The time the Sender was created. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


