# DeleteInvalidEmails

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInvalidEmails**](DeleteInvalidEmails.md#DeleteInvalidEmails) | **Delete** /v3/suppression/invalid_emails | Delete invalid emails



## DeleteInvalidEmails

> DeleteInvalidEmails(ctx, optional)

Delete invalid emails

**This endpoint allows you to remove email addresses from your invalid email address list.**  There are two options for deleting invalid email addresses:   1) You can delete all invalid email addresses by setting `delete_all` to true in the request body. 2) You can delete some invalid email addresses by specifying certain addresses in an array in the request body.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInvalidEmailsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**DeleteInvalidEmailsRequest** | [**DeleteInvalidEmailsRequest**](DeleteInvalidEmailsRequest.md) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

