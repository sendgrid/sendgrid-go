# DeleteSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSegment**](DeleteSegment.md#DeleteSegment) | **Delete** /v3/marketing/segments/{SegmentId} | Delete Segment



## DeleteSegment

> map[string]interface{} DeleteSegment(ctx, SegmentId)

Delete Segment

**This endpoint allows you to delete a segment by `segment_id`.**  Note that deleting a segment does not delete the contacts associated with the segment by default. Contacts associated with a deleted segment will remain in your list of all contacts and any other segments they belong to.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

