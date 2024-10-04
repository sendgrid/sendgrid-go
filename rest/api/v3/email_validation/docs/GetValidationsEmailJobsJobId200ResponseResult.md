# GetValidationsEmailJobsJobId200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the Bulk Email Address Validation Job. |[optional] 
**Status** | [**Status1**](Status1.md) | The status of the Bulk Email Address Validation Job. |[optional] 
**Segments** | **float32** | The total number of segments in the Bulk Email Address Validation Job. There are 1,500 email addresses per segment. The value is `0` until the Job `status` is `Processing`. |[optional] 
**SegmentsProcessed** | **float32** | The number of segments processed at the time of the request. 100 segments process in parallel at a time. |[optional] 
**IsDownloadAvailable** | **bool** | Boolean indicating whether the results CSV file is available for download. |[optional] 
**StartedAt** | **float32** | The ISO8601 timestamp when the Job was created. This is the time at which the upload request was sent to the `upload_uri`. |[optional] 
**FinishedAt** | **float32** | The ISO8601 timestamp when the Job was finished. |[optional] 
**Errors** | [**[]GetValidationsEmailJobsJobId200ResponseResultErrorsInner**](GetValidationsEmailJobsJobId200ResponseResultErrorsInner.md) | Array containing error messages related to the Bulk Email Address Validation Job. Array is empty if no errors ocurred. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


