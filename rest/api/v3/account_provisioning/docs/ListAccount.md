# ListAccount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccount**](ListAccount.md#ListAccount) | **Get** /v3/partners/accounts | Get all accounts



## ListAccount

> AccountList ListAccount(ctx, optional)

Get all accounts

Retrieves all accounts under the organization.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**Offset** | **string** | The last item successfully retrieved
**Limit** | **int32** | The number of items to return

### Return type

[**AccountList**](AccountList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

