# DeleteContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContact**](DeleteContact.md#DeleteContact) | **Delete** /v3/marketing/lists/{Id}/contacts | Remove Contacts from a List



## DeleteContact

> DeleteContact202Response DeleteContact(ctx, ContactIdsId)

Remove Contacts from a List

**This endpoint allows you to remove contacts from a given list.**  The contacts will not be deleted. Only their list membership will be changed.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the list on which you want to perform the operation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteContactParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteContact202Response**](DeleteContact202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

