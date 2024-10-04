# RefreshSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshSegment**](RefreshSegment.md#RefreshSegment) | **Post** /v3/marketing/segments/2.0/refresh/{SegmentId} | Manually refresh a segment



## RefreshSegment

> SegmentRefresh202 RefreshSegment(ctx, SegmentIdSegmentRefreshRequest)

Manually refresh a segment

Manually refresh a segment by segment ID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a RefreshSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SegmentRefresh202**](SegmentRefresh202.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

