# DeleteMarketingList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingList**](DeleteMarketingList.md#DeleteMarketingList) | **Delete** /v3/marketing/lists/{Id} | Delete a list



## DeleteMarketingList

> DeleteMarketingList200Response DeleteMarketingList(ctx, Idoptional)

Delete a list

**This endpoint allows you to deletes a specific list.**  Optionally, you can also delete contacts associated to the list. The query parameter, `delete_contacts=true`, will delete the list and start an asynchronous job to delete associated contacts.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the list on which you want to perform the operation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMarketingListParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteContacts** | **bool** | Flag indicates that all contacts on the list are also to be deleted.

### Return type

[**DeleteMarketingList200Response**](DeleteMarketingList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

