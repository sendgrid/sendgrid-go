# ListSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSingleSend**](ListSingleSend.md#ListSingleSend) | **Get** /v3/marketing/singlesends | Get All Single Sends



## ListSingleSend

> ListSingleSend200Response ListSingleSend(ctx, optional)

Get All Single Sends

**This endpoint allows you to retrieve all your Single Sends.**  Returns all of your Single Sends with condensed details about each, including the Single Sends' IDs. For more details about an individual Single Send, pass the Single Send's ID to the `/marketing/singlesends/{id}` endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | 
**PageToken** | **string** | 

### Return type

[**ListSingleSend200Response**](ListSingleSend200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

