# ListAsmGroup

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAsmGroup**](ListAsmGroup.md#ListAsmGroup) | **Get** /v3/asm/groups | Retrieve all suppression groups associated with the user.



## ListAsmGroup

> []SuppressionGroup ListAsmGroup(ctx, optional)

Retrieve all suppression groups associated with the user.

**This endpoint allows you to retrieve a list of all suppression groups created by this user.**  This endpoint can also return information for multiple group IDs that you include in your request. To add a group ID to your request, simply append `?id=123456&id=123456`, with the appropriate group IDs.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAsmGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**Id** | **int32** | The ID of the suppression group(s) you want to retrieve.
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**[]SuppressionGroup**](SuppressionGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

