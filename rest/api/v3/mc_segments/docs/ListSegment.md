# ListSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSegment**](ListSegment.md#ListSegment) | **Get** /v3/marketing/segments | Get List of Segments



## ListSegment

> ListSegment200Response ListSegment(ctx, optional)

Get List of Segments

**This endpoint allows you to retrieve a list of segments.**  The query param `parent_list_ids` is treated as a filter.  Any match will be returned.  Zero matches will return a response code of 200 with an empty `results` array.  `parent_list_ids` | `no_parent_list_id` | `ids` | `result` -----------------:|:--------------------:|:-------------:|:-------------: empty | false | empty | all segments values list_ids | false | empty | segments filtered by list_ids values list_ids |true | empty | segments filtered by list_ids and segments with no parent list_ids empty empty | true | empty | segments with no parent list_ids anything | anything | ids | segments with matching segment ids |

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ids** | **[]string** | A list of segment IDs to retrieve. When this parameter is included, the `no_parent_list_ids` and `parent_list_ids` parameters are ignored and only segments with given IDs are returned.
**ParentListIds** | **string** | A comma separated list of list ids to be used when searching for segments with the specified parent_list_id, no more than 50 is allowed
**NoParentListId** | **bool** | If set to `true` segments with an empty value of `parent_list_id` will be returned in the filter.  If the value is not present it defaults to 'false'.

### Return type

[**ListSegment200Response**](ListSegment200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

