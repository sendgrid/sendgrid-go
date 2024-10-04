# DeleteContactDbList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContactDbList**](DeleteContactDbList.md#DeleteContactDbList) | **Delete** /v3/contactdb/lists/{ListId} | Delete a List



## DeleteContactDbList

> interface{} DeleteContactDbList(ctx, ListIdoptional)

Delete a List

**This endpoint allows you to delete a specific recipient list with the given ID.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ListId** | **int32** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteContactDbListParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteContacts** | [**DeleteContacts**](DeleteContactsDeleteContacts.md) | Adds the ability to delete all contacts on the list in addition to deleting the list.
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Body** | **interface{}** | 

### Return type

**interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

