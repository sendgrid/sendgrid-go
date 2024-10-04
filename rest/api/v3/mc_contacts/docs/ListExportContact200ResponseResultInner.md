# ListExportContact200ResponseResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Export jobs ID. |[optional] 
**Status** | **string** | Allowed values: `pending`, `ready`, or `failure`. |[optional] 
**CreatedAt** | **string** | This ISO8601 timestamp when the export was created. |[optional] 
**CompletedAt** | **string** | This ISO8601 timestamp when the export was completed. |[optional] 
**ExpiresAt** | **string** | This ISO8601 timestamp when the export expires. |[optional] 
**Urls** | **[]string** | One or more download URLs for the contact file(s) if the status is `ready`. |[optional] 
**UserId** | **string** | User ID. |[optional] 
**ExportType** | **string** | Allowed types: `contacts_export`, `list_export`, or `segment_export`. |[optional] 
**Segments** | [**[]ListExportContact200ResponseResultInnerSegmentsInner**](ListExportContact200ResponseResultInnerSegmentsInner.md) |  |[optional] 
**Lists** | [**[]ListExportContact200ResponseResultInnerSegmentsInner**](ListExportContact200ResponseResultInnerSegmentsInner.md) |  |[optional] 
**Metadata** | [**ListExportContact200ResponseResultInnerMetadata**](ListExportContact200ResponseResultInnerMetadata.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


