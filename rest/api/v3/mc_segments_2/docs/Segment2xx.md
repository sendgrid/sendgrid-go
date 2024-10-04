# Segment2xx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID assigned to the segment when created. |
**Name** | **string** | Name of the segment. |
**QueryDsl** | **string** | SQL query which will filter contacts based on the conditions provided |
**ContactsCount** | **int32** | Total number of contacts present in the segment |
**ContactsSample** | [**[]ContactResponse**](ContactResponse.md) | A subset of all contacts that are in this segment |
**CreatedAt** | **string** | ISO8601 timestamp of when the object was created |
**UpdatedAt** | **string** | ISO8601 timestamp of when the object was last updated |
**SampleUpdatedAt** | **string** | ISO8601 timestamp of when the samples were last updated |
**NextSampleUpdate** | **string** | ISO8601 timestamp of when the samples will be next updated |
**ParentListIds** | **[]string** | The array of list ids to filter contacts on when building this segment. It allows only one such list id for now. We will support more in future |
**QueryVersion** | **string** | If not set, segment contains a Query for use with Segment v1 APIs. If set to '2', segment contains a SQL query for use in v2. |
**Status** | [**SegmentStatusResponse**](SegmentStatusResponse.md) |  |
**RefreshesUsed** | **int32** | The number of times a segment has been manually refreshed since start of today in the user's timezone. |[optional] 
**MaxRefreshes** | **int32** | The maximum number of manual refreshes allowed per day for this segment. Currently, only 2 are allowed. |[optional] 
**LastRefreshedAt** | **string** | The ISO8601 timestamp when the segment was last refreshed in UTC time. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


