# SendMailRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The Base64 encoded content of the attachment. |
**Type** | **string** | The MIME type of the content you are attaching (e.g., `image/jpeg` or `application/pdf`). |[optional] 
**Filename** | **string** | The attachment's filename, including the file extension. |
**Disposition** | [**Disposition**](Disposition.md) | The attachment's content-disposition specifies how you would like the attachment to be displayed. For example, `inline` results in the attached file being displayed automatically within the message while `attachment` results in the attached file requiring some action to be taken before it is displayed such as opening or downloading the file. |[optional] 
**ContentId** | **string** | The attachment's content ID. The `content_id` is used when the `disposition` is set to `inline` and the attachment is an image, allowing the file to be displayed within the body of the email. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


