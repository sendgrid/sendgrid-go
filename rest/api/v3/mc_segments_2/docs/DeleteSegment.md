# DeleteSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSegment**](DeleteSegment.md#DeleteSegment) | **Delete** /v3/marketing/segments/2.0/{SegmentId} | Delete segment



## DeleteSegment

> DeleteSegment(ctx, SegmentId)

Delete segment

**This endpoint allows you to delete a segment by ID.**

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

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

