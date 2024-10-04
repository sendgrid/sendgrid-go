# SearchSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSingleSend**](SearchSingleSend.md#SearchSingleSend) | **Post** /v3/marketing/singlesends/search | Get Single Sends Search



## SearchSingleSend

> ListSingleSend200Response SearchSingleSend(ctx, optional)

Get Single Sends Search

**This endpoint allows you to search for Single Sends based on specified criteria.**  You can search for Single Sends by passing a combination of values using the `name`, `status`, and `categories` request body fields.  For example, if you want to search for all Single Sends that are \"drafts\" or \"scheduled\" and also associated with the category \"shoes,\" your request body may look like the example below.  ```javascript {   \"status\": [     \"draft\",     \"scheduled\"   ],   \"categories\": [     \"shoes\"   ], } ```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a SearchSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | 
**PageToken** | **string** | 
**SinglesendSearch** | [**SinglesendSearch**](SinglesendSearch.md) | 

### Return type

[**ListSingleSend200Response**](ListSingleSend200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

