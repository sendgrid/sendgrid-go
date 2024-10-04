# GetDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDesign**](GetDesign.md#GetDesign) | **Get** /v3/designs/{Id} | Get Design



## GetDesign

> DesignOutput GetDesign(ctx, Id)

Get Design

**This endpoint allows you to retrieve a single design**.  A GET request to `/designs/{id}` will retrieve details about a specific design in your Design Library.  This endpoint is valuable when retrieving information stored in a field that you wish to update using a PATCH request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Design you want to duplicate.

### Other Parameters

Other parameters are passed through a pointer to a GetDesignParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DesignOutput**](DesignOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

