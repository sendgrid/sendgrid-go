# ListRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRecipient**](ListRecipient.md#ListRecipient) | **Get** /v3/contactdb/recipients | Retrieve recipients



## ListRecipient

> ListRecipientsResponse ListRecipient(ctx, optional)

Retrieve recipients

**This endpoint allows you to retrieve all of your Marketing Campaigns recipients.**  Batch deletion of a page makes it possible to receive an empty page of recipients before reaching the end of the list of recipients. To avoid this issue; iterate over pages until a 404 is retrieved.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRecipientParams struct


Name | Type | Description
------------- | ------------- | -------------
**Page** | **int32** | Page index of first recipients to return (must be a positive integer)
**PageSize** | **int32** | Number of recipients to return at a time (must be a positive integer between 1 and 1000)
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ListRecipientsResponse**](ListRecipientsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

