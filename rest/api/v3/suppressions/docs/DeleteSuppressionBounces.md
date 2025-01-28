# DeleteSuppressionBounces

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSuppressionBounces**](DeleteSuppressionBounces.md#DeleteSuppressionBounces) | **Delete** /v3/suppression/bounces | Delete bounces



## DeleteSuppressionBounces

> interface{} DeleteSuppressionBounces(ctx, optional)

Delete bounces

**This endpoint allows you to delete all emails on your bounces list.**  There are two options for deleting bounced emails:   1. You can delete all bounced emails by setting `delete_all` to `true` in the request body.  2. You can delete a selection of bounced emails by specifying the email addresses in the `emails` array of the request body.   **WARNING:** You can not have both `emails` and `delete_all` set.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSuppressionBouncesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**DeleteSuppressionBouncesRequest** | [**DeleteSuppressionBouncesRequest**](DeleteSuppressionBouncesRequest.md) | 

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

