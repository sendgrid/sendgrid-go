# UpdateSubuserWebsiteAccess

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubuserWebsiteAccess**](UpdateSubuserWebsiteAccess.md#UpdateSubuserWebsiteAccess) | **Patch** /v3/subusers/{SubuserName}/website_access | Enable/Disable website access for a Subuser



## UpdateSubuserWebsiteAccess

> map[string]interface{} UpdateSubuserWebsiteAccess(ctx, SubuserNameoptional)

Enable/Disable website access for a Subuser

Enable/Disable website access for a Subuser, while still preserving email send functionality.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubuserWebsiteAccessParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateSubuserWebsiteAccessRequest** | [**UpdateSubuserWebsiteAccessRequest**](UpdateSubuserWebsiteAccessRequest.md) | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

