# ListContactCount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContactCount**](ListContactCount.md#ListContactCount) | **Get** /v3/marketing/lists/{Id}/contacts/count | Get List Contact Count



## ListContactCount

> ListContactCount200Response ListContactCount(ctx, Id)

Get List Contact Count

**This endpoint returns the number of contacts on a specific list.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListContactCountParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ListContactCount200Response**](ListContactCount200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

