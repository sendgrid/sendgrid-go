# UpdateDesignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Design. |[optional] [default to "My Design"]
**HtmlContent** | **string** | The HTML content of the Design. |[optional] 
**PlainContent** | **string** | Plain text content of the Design. |[optional] [default to "<generated from html_content if left empty>"]
**GeneratePlainContent** | **bool** | If true, plain_content is always generated from html_content. If false, plain_content is not altered. |[optional] [default to true]
**Subject** | **string** | Subject of the Design. |[optional] 
**Categories** | **[]string** | The list of categories applied to the design |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


