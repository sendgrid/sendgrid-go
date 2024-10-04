# GetPreBuiltDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreBuiltDesign**](GetPreBuiltDesign.md#GetPreBuiltDesign) | **Get** /v3/designs/pre-builts/{Id} | Get SendGrid Pre-built Design



## GetPreBuiltDesign

> DesignOutput GetPreBuiltDesign(ctx, Id)

Get SendGrid Pre-built Design

**This endpoint allows you to retrieve a single pre-built design**.  A GET request to `/designs/pre-builts/{id}` will retrieve details about a specific pre-built design.  This endpoint is valuable when retrieving details about a pre-built design that you wish to duplicate and modify.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the pre-built Design you want to duplicate.

### Other Parameters

Other parameters are passed through a pointer to a GetPreBuiltDesignParams struct


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

