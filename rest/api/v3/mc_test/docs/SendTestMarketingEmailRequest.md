# SendTestMarketingEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The ID of the template that you would like to use. If you use a template that contains a subject and content (either text or HTML), then those values specified at the personalizations or message level will not be used. |
**VersionIdOverride** | **string** |  You can override the active template with an alternative template version by passing the version ID in this field. If this field is blank, the active template version will be used. |[optional] 
**SenderId** | **int32** | This ID must belong to a verified sender. Alternatively, you may supply a `from_address` email. |[optional] 
**CustomUnsubscribeUrl** | **string** | A custom unsubscribe URL. |[optional] 
**SuppressionGroupId** | **int32** |  |[optional] 
**Emails** | **[]string** | An array of email addresses you want to send the test message to. |
**FromAddress** | **string** | You can either specify this address or specify a verified sender ID. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


