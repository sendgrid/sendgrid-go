# TransactionalTemplateVersionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | [**[]TransactionalTemplateWarning**](TransactionalTemplateWarning.md) |  |[optional] 
**Active** | [**Active1**](Active1.md) | Set the version as the active version associated with the template. Only one version of a template can be active. The first version created for a template will automatically be set to Active. |[optional] 
**Name** | **string** | Name of the transactional template version. |
**HtmlContent** | **string** | The HTML content of the Design. |[optional] 
**PlainContent** | **string** | Plain text content of the Design. |[optional] 
**GeneratePlainContent** | **bool** | If true, plain_content is always generated from html_content. If false, plain_content is not altered. |[optional] [default to true]
**Subject** | **string** | Subject of the new transactional template version. |
**Editor** | [**Editor1**](Editor1.md) | The editor used in the UI. |[optional] 
**TestData** | **string** | For dynamic templates only, the mock json data that will be used for template preview and test sends. |[optional] 
**Id** | **string** | ID of the transactional template version. |[optional] 
**TemplateId** | **string** | ID of the transactional template. |[optional] 
**UpdatedAt** | **string** | The date and time that this transactional template version was updated. |[optional] 
**ThumbnailUrl** | **string** | A Thumbnail preview of the template's html content. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

