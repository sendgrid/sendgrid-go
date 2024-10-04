# ContactdbRecipient200RecipientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this recipient. |[optional] 
**CreatedAt** | **float32** | The time this record was created in your contactdb, in unixtime. |[optional] 
**CustomFields** | [**[]ContactdbCustomFieldIdValue**](ContactdbCustomFieldIdValue.md) | The custom fields assigned to this recipient and their values. |[optional] 
**Email** | **string** | The email address of this recipient. This is a default custom field that SendGrid provides. |
**FirstName** | **string** | The first name of this recipient. This is a default custom field that SendGrid provides. |[optional] 
**LastName** | **string** | The last name of the recipient. |[optional] 
**LastClicked** | **float32** | The last time this recipient clicked a link from one of your campaigns, in unixtime. |[optional] 
**LastEmailed** | **float32** | The last time this user was emailed by one of your campaigns, in unixtime. |[optional] 
**LastOpened** | **float32** | The last time this recipient opened an email from you, in unixtime. |[optional] 
**UpdatedAt** | **float32** | The last update date for this recipient's record. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


