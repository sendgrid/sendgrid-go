# AllSegments200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID assigned to the segment when created. |
**Name** | **string** | Name of the segment. |
**ContactsCount** | **int32** | Total number of contacts present in the segment |
**CreatedAt** | **string** | ISO8601 timestamp of when the object was created |
**UpdatedAt** | **string** | ISO8601 timestamp of when the object was last updated |
**SampleUpdatedAt** | **string** | ISO8601 timestamp of when the samples were last updated |
**NextSampleUpdate** | **string** | ISO8601 timestamp of when the samples will be next updated |
**ParentListIds** | **[]string** | The array of list ids to filter contacts on when building this segment. It allows only one such list id for now. We will support more in future |
**QueryVersion** | **string** | If not set, segment contains a query for use with Segment v1 APIs. If set to '2', segment contains a SQL query for use in v2. |
**Metadata** | [**Metadata**](Metadata.md) |  |[optional] 
**Status** | [**SegmentStatusResponse**](SegmentStatusResponse.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


