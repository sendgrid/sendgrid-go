# UpdateSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSegment**](UpdateSegment.md#UpdateSegment) | **Patch** /v3/marketing/segments/2.0/{SegmentId} | Update Segment



## UpdateSegment

> Segment2xx UpdateSegment(ctx, SegmentIdoptional)

Update Segment

Segment `name` has to be unique. A user can not create a new segment with an existing segment name.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**SegmentUpdate** | [**SegmentUpdate**](SegmentUpdate.md) | 

### Return type

[**Segment2xx**](Segment2xx.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

