# CreateSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](CreateSegment.md#CreateSegment) | **Post** /v3/marketing/segments/2.0 | Create Segment



## CreateSegment

> Segment2xx CreateSegment(ctx, optional)

Create Segment

Segment `name` has to be unique. A user can not create a new segment with an existing segment name.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**SegmentWriteV2** | [**SegmentWriteV2**](SegmentWriteV2.md) | 

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

