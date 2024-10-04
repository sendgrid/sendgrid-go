# DeleteSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSingleSend**](DeleteSingleSend.md#DeleteSingleSend) | **Delete** /v3/marketing/singlesends/{Id} | Delete Single Send by ID



## DeleteSingleSend

> DeleteSingleSend(ctx, Id)

Delete Single Send by ID

**This endpoint allows you to delete one Single Send using a Single Send ID.**  To first retrieve all your Single Sends' IDs, you can make a GET request to the `/marketing/singlensends` endpoint.  Please note that a `DELETE` request is permanent, and your Single Send will not be recoverable after deletion.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSingleSendParams struct


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

