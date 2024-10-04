# GetMessage

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessage**](GetMessage.md#GetMessage) | **Get** /v3/messages/{MsgId} | Filter messages by message ID



## GetMessage

> Message GetMessage(ctx, MsgId)

Filter messages by message ID

Get all of the details about the specified message.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MsgId** | **string** | The ID of the message you are requesting details for.

### Other Parameters

Other parameters are passed through a pointer to a GetMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Message**](Message.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

