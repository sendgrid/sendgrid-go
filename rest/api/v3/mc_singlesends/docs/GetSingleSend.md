# GetSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSingleSend**](GetSingleSend.md#GetSingleSend) | **Get** /v3/marketing/singlesends/{Id} | Get Single Send by ID



## GetSingleSend

> SinglesendResponse GetSingleSend(ctx, Id)

Get Single Send by ID

**This endpoint allows you to retrieve details about one Single Send using a Single Send ID.**  You can retrieve all of your Single Sends by making a GET request to the `/marketing/singlesends` endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SinglesendResponse**](SinglesendResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

