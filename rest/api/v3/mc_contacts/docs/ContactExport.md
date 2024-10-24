# ContactExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Status** | [**Status**](Status.md) | The export job's status. Allowed values: `pending`, `ready`, or `failure`. |
**CreatedAt** | **string** | The ISO8601 timestamp when the export was begun. |
**UpdatedAt** | **string** | The ISO8601 timestamp when the export was updated. |
**CompletedAt** | **string** | The ISO8601 timestamp when the export was completed. |[optional] 
**ExpiresAt** | **string** | The ISO8601 timestamp when the exported file on S3 will expire. |
**Urls** | **[]string** | One or more download URLs for the contact file if the status is `ready`. |[optional] 
**Message** | **string** | A human readable message if the status is `failure`. |[optional] 
**Metadata** | [**Metadata**](Metadata.md) |  |[optional] 
**ContactCount** | **int32** | The total number of exported contacts. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


