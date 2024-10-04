# UpdateMarketingList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMarketingList**](UpdateMarketingList.md#UpdateMarketingList) | **Patch** /v3/marketing/lists/{Id} | Update List



## UpdateMarketingList

> List UpdateMarketingList(ctx, Idoptional)

Update List

**This endpoint updates the name of a list.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the list on which you want to perform the operation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMarketingListParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateMarketingListRequest** | [**UpdateMarketingListRequest**](UpdateMarketingListRequest.md) | 

### Return type

[**List**](List.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

