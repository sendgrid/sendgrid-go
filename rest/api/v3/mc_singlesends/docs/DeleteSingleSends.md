# DeleteSingleSends

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSingleSends**](DeleteSingleSends.md#DeleteSingleSends) | **Delete** /v3/marketing/singlesends | Bulk Delete Single Sends



## DeleteSingleSends

> DeleteSingleSends(ctx, optional)

Bulk Delete Single Sends

**This endpoint allows you to delete multiple Single Sends using an array of Single Sends IDs.**  To first retrieve all your Single Sends' IDs, you can make a GET request to the `/marketing/singlensends` endpoint.  Please note that a DELETE request is permanent, and your Single Sends will not be recoverable after deletion.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSingleSendsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ids** | **[]string** | Single Send IDs to delete

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

