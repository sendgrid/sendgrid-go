# GetMarketingList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingList**](GetMarketingList.md#GetMarketingList) | **Get** /v3/marketing/lists/{Id} | Get a List by ID



## GetMarketingList

> GetMarketingList200Response GetMarketingList(ctx, Idoptional)

Get a List by ID

**This endpoint returns data about a specific list.** Setting the optional parameter `contact_sample=true` returns the `contact_sample` in the response body. Up to 50 of the most recent contacts uploaded or attached to a list will be returned. The full contact count is also returned.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the list on which you want to perform the operation.

### Other Parameters

Other parameters are passed through a pointer to a GetMarketingListParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContactSample** | **bool** | Setting this parameter to the true  will cause the contact_sample to be returned

### Return type

[**GetMarketingList200Response**](GetMarketingList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

