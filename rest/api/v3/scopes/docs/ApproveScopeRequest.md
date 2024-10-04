# ApproveScopeRequest

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveScopeRequest**](ApproveScopeRequest.md#ApproveScopeRequest) | **Patch** /v3/scopes/requests/{RequestId}/approve | Approve access request



## ApproveScopeRequest

> ApproveScopeRequest200Response ApproveScopeRequest(ctx, RequestId)

Approve access request

**This endpoint allows you to approve an access attempt.**  **Note:** Only teammate admins may approve another teammate’s access request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RequestId** | **string** | The ID of the request that you want to approve.

### Other Parameters

Other parameters are passed through a pointer to a ApproveScopeRequestParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ApproveScopeRequest200Response**](ApproveScopeRequest200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

