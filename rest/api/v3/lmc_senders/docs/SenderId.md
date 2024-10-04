# SenderId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | **string** | A nickname for the sender identity. Not used for sending. |
**From** | [**SenderIdRequestFrom**](SenderIdRequestFrom.md) |  |[optional] 
**ReplyTo** | [**SenderIdRequestReplyTo**](SenderIdRequestReplyTo.md) |  |[optional] 
**Address** | **string** | The physical address of the sender identity. |
**Address2** | **string** | Additional sender identity address information. |[optional] 
**City** | **string** | The city of the sender identity. |
**State** | **string** | The state of the sender identity. |[optional] 
**Zip** | **string** | The zipcode of the sender identity. |[optional] 
**Country** | **string** | The country of the sender identity. |
**Id** | **int32** | The unique identifier of the sender identity. |[optional] 
**Verified** | **bool** | If the sender identity is verified or not. Only verified sender identities can be used to send email. |[optional] 
**UpdatedAt** | **int32** | The time the sender identity was last updated. |[optional] 
**CreatedAt** | **int32** | The time the sender identity was created. |[optional] 
**Locked** | **bool** | True when the sender id is associated to a campaign in the Draft, Scheduled, or In Progress status. You cannot update or delete a locked sender identity. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


