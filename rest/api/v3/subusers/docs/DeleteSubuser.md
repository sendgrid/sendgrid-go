# DeleteSubuser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubuser**](DeleteSubuser.md#DeleteSubuser) | **Delete** /v3/subusers/{SubuserName} | Delete a subuser



## DeleteSubuser

> DeleteSubuser(ctx, SubuserName)

Delete a subuser

**This endpoint allows you to delete a subuser.**  This is a permanent action. Once deleted, a subuser cannot be retrieved.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubuserParams struct


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

