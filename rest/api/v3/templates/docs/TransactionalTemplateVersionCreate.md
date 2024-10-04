# TransactionalTemplateVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | [**Active**](Active.md) | Set the version as the active version associated with the template (0 is inactive, 1 is active). Only one version of a template can be active. The first version created for a template will automatically be set to Active. |[optional] 
**Name** | **string** | Name of the transactional template version. |
**HtmlContent** | **string** | The HTML content of the version. Maximum of 1048576 bytes allowed. |[optional] 
**PlainContent** | **string** | Text/plain content of the transactional template version. Maximum of 1048576 bytes allowed. |[optional] [default to "<generated from html_content if left empty>"]
**GeneratePlainContent** | **bool** | If true, plain_content is always generated from html_content. If false, plain_content is not altered. |[optional] [default to true]
**Subject** | **string** | Subject of the new transactional template version. |
**Editor** | [**Editor**](Editor.md) | The editor used in the UI. |[optional] 
**TestData** | **string** | For dynamic templates only, the mock json data that will be used for template preview and test sends. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


