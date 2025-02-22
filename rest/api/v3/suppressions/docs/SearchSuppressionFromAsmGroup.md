# SearchSuppressionFromAsmGroup

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSuppressionFromAsmGroup**](SearchSuppressionFromAsmGroup.md#SearchSuppressionFromAsmGroup) | **Post** /v3/asm/groups/{GroupId}/suppressions/search | Search for suppressions within a group



## SearchSuppressionFromAsmGroup

> []string SearchSuppressionFromAsmGroup(ctx, GroupIdoptional)

Search for suppressions within a group

**This endpoint allows you to search a suppression group for multiple suppressions.**  When given a list of email addresses and a group ID, this endpoint will only return the email addresses that have been unsubscribed from the given group.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**GroupId** | **string** | The ID of the suppression group that you would like to search.

### Other Parameters

Other parameters are passed through a pointer to a SearchSuppressionFromAsmGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**SuppressionsRequest** | [**SuppressionsRequest**](SuppressionsRequest.md) | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

