# ListMarketingList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMarketingList**](ListMarketingList.md#ListMarketingList) | **Get** /v3/marketing/lists | Get All Lists



## ListMarketingList

> ListMarketingList200Response ListMarketingList(ctx, optional)

Get All Lists

**This endpoint returns an array of all of your contact lists.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMarketingListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **float32** | Maximum number of elements to return. Defaults to 100, returns 1000 max
**PageToken** | **string** | 

### Return type

[**ListMarketingList200Response**](ListMarketingList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

