# GetSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSegment**](GetSegment.md#GetSegment) | **Get** /v3/marketing/segments/2.0/{SegmentId} | Get Segment by ID



## GetSegment

> Segment2xx GetSegment(ctx, SegmentIdoptional)

Get Segment by ID

Get Marketing Campaigns Segment by ID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContactsSample** | **bool** | Defaults to `true`. Set to `false` to exclude the contacts_sample in the response.

### Return type

[**Segment2xx**](Segment2xx.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

