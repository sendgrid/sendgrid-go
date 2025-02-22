# UpdateSubuser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubuser**](UpdateSubuser.md#UpdateSubuser) | **Patch** /v3/subusers/{SubuserName} | Enable/disable a subuser



## UpdateSubuser

> UpdateSubuser(ctx, SubuserNameoptional)

Enable/disable a subuser

**This endpoint allows you to enable or disable a subuser.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubuserParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateSubuserRequest** | [**UpdateSubuserRequest**](UpdateSubuserRequest.md) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

