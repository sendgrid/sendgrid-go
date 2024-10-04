# DuplicateSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DuplicateSingleSend**](DuplicateSingleSend.md#DuplicateSingleSend) | **Post** /v3/marketing/singlesends/{Id} | Duplicate Single Send



## DuplicateSingleSend

> SinglesendResponse DuplicateSingleSend(ctx, Idoptional)

Duplicate Single Send

**This endpoint allows you to duplicate an existing Single Send using its Single Send ID.**  Duplicating a Single Send is useful when you want to create a Single Send but don't want to start from scratch. Once duplicated, you can update or edit the Single Send by making a PATCH request to the `/marketing/singlesends/{id}` endpoint.   If you leave the `name` field blank, your duplicate will be assigned the name of the Single Send it was copied from with the text “Copy of ” prepended to it. The `name` field length is limited to 100 characters, so the end of the new Single Send name, including “Copy of ”, will be trimmed if the name exceeds this limit.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DuplicateSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------
**DuplicateSingleSendRequest** | [**DuplicateSingleSendRequest**](DuplicateSingleSendRequest.md) | 

### Return type

[**SinglesendResponse**](SinglesendResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

