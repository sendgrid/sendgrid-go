# GetContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContact**](GetContact.md#GetContact) | **Get** /v3/marketing/contacts/{Id} | Get a Contact by ID



## GetContact

> ContactDetails3 GetContact(ctx, Id)

Get a Contact by ID

**This endpoint returns the full details and all fields for the specified contact**. The \"Get Contacts by Identifier\" endpoint can be used to get the ID of a contact.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetContactParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContactDetails3**](ContactDetails3.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

