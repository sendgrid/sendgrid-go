# AddRecipientsToContactDbList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecipientsToContactDbList**](AddRecipientsToContactDbList.md#AddRecipientsToContactDbList) | **Post** /v3/contactdb/lists/{ListId}/recipients | Add Multiple Recipients to a List



## AddRecipientsToContactDbList

> interface{} AddRecipientsToContactDbList(ctx, ListIdoptional)

Add Multiple Recipients to a List

**This endpoint allows you to add multiple recipients to a list.**  Adds existing recipients to a list, passing in the recipient IDs to add. Recipient IDs (base64-encoded email addresses) should be passed exactly as they are returned from recipient endpoints.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ListId** | **int32** | The id of the list of recipients you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a AddRecipientsToContactDbListParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**RequestBody** | **[]string** | 

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

