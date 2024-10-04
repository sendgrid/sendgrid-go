# ListSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSegment**](ListSegment.md#ListSegment) | **Get** /v3/marketing/segments/2.0 | Get List of Segments



## ListSegment

> AllSegments200 ListSegment(ctx, optional)

Get List of Segments

**This endpoint allows you to retrieve a list of segments.**  The query param `parent_list_ids` is treated as a filter.  Any match will be returned.  Zero matches will return a response code of 200 with an empty `results` array.  `parent_list_ids` | `no_parent_list_id` | `ids` | `result` -----------------:|:--------------------:|:-------------:|:-------------: empty | false | empty | all segments values list_ids | false | empty | segments filtered by list_ids values list_ids |true | empty | segments filtered by list_ids and segments with no parent list_ids empty empty | true | empty | segments with no parent list_ids anything | anything | ids | segments with matching segment ids |

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ids** | **[]string** | A list of segment IDs to retrieve. When this parameter is included, the `no_parent_list_ids` and `parent_list_ids` parameters are ignored and only segments with given IDs are returned.
**ParentListIds** | **string** | A comma separated list up to 50 in size, to filter segments on.  Only segments that have any of these list ids as the parent list will be retrieved. This is different from the parameter of the same name used when creating a segment.
**NoParentListId** | **bool** | If set to `true`, segments with an empty value of `parent_list_id` will be returned in the filter.  If the value is not present, it defaults to 'false'.

### Return type

[**AllSegments200**](AllSegments200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

