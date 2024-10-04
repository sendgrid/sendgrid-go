# ExportContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListIds** | **[]string** | IDs of the contact lists you want to export. |[optional] 
**SegmentIds** | **[]string** | IDs of the contact segments you want to export. |[optional] 
**Notifications** | [**ExportContactRequestNotifications**](ExportContactRequestNotifications.md) |  |[optional] 
**FileType** | [**FileType**](FileType.md) | File type for export file. Choose from `json` or `csv`. |[optional] 
**MaxFileSize** | **int32** | The maximum size of an export file in MB. Note that when this option is specified, multiple output files may be returned from the export. |[optional] [default to 5000]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


