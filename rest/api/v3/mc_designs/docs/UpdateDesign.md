# UpdateDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateDesign**](UpdateDesign.md#UpdateDesign) | **Patch** /v3/designs/{Id} | Update Design



## UpdateDesign

> DesignOutput UpdateDesign(ctx, Idoptional)

Update Design

**This endpoint allows you to edit a design**.  The Design API supports PATCH requests, which allow you to make partial updates to a single design. Passing data to a specific field will update only the data stored in that field; all other fields will be unaltered.  For example, updating a design's name requires that you make a PATCH request to this endpoint with data specified for the `name` field only.  ``` {     \"name\": \"<Updated Name>\" } ```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Design you want to duplicate.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDesignParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateDesignRequest** | [**UpdateDesignRequest**](UpdateDesignRequest.md) | 

### Return type

[**DesignOutput**](DesignOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

