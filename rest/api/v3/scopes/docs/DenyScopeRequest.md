# DenyScopeRequest

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DenyScopeRequest**](DenyScopeRequest.md#DenyScopeRequest) | **Delete** /v3/scopes/requests/{RequestId} | Deny access request



## DenyScopeRequest

> DenyScopeRequest(ctx, RequestId)

Deny access request

**This endpoint allows you to deny an attempt to access your account.**  **Note:** Only teammate admins may delete a teammate's access request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RequestId** | **string** | The ID of the request that you want to deny.

### Other Parameters

Other parameters are passed through a pointer to a DenyScopeRequestParams struct


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

