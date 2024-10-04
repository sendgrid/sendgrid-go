# SegmentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**ContactsCount** | **int32** |  |
**CreatedAt** | [**time.Time**](time.Time.md) | ISO8601 of created timestamp  |
**Name** | **string** |  |[optional] 
**ParentListId** | **string** | The id of the list if this segment is a child of a list.  This implies the query `AND CONTAINS(list_ids, ${parent_list_id})` |[optional] 
**SampleUpdatedAt** | [**time.Time**](time.Time.md) | ISO8601 timestamp the sample was last updated |
**UpdatedAt** | [**time.Time**](time.Time.md) | ISO8601 timestamp the object was last updated |
**NextSampleUpdate** | **string** | ISO8601 string that is equal to `sample_updated_at` plus an internally calculated offset that depends on how often contacts enter or exit segments as the scheduled pipeline updates the samples. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


