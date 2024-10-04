# GetSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSegment**](GetSegment.md#GetSegment) | **Get** /v3/marketing/segments/{SegmentId} | Get Segment by ID



## GetSegment

> FullSegment GetSegment(ctx, SegmentIdoptional)

Get Segment by ID

**This endpoint allows you to retrieve a single segment by ID.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**QueryJson** | **bool** | Defaults to `false`.  Set to `true` to return the parsed SQL AST as a JSON object in the field `query_json`

### Return type

[**FullSegment**](FullSegment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

