# DeleteDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDesign**](DeleteDesign.md#DeleteDesign) | **Delete** /v3/designs/{Id} | Delete Design



## DeleteDesign

> map[string]interface{} DeleteDesign(ctx, Id)

Delete Design

**This endpoint allows you to delete a single design**.  Be sure to check the ID of the design you intend to delete before making this request; deleting a design is a permanent action.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Design you want to duplicate.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDesignParams struct


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

