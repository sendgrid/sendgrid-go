# AddSuppressionToAsmGroup

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSuppressionToAsmGroup**](AddSuppressionToAsmGroup.md#AddSuppressionToAsmGroup) | **Post** /v3/asm/groups/{GroupId}/suppressions | Add suppressions to a suppression group



## AddSuppressionToAsmGroup

> AddSuppressionToAsmGroup201Response AddSuppressionToAsmGroup(ctx, GroupIdoptional)

Add suppressions to a suppression group

**This endpoint allows you to add email addresses to an unsubscribe group.**  If you attempt to add suppressions to a group that has been deleted or does not exist, the suppressions will be added to the global suppressions list.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**GroupId** | **string** | The id of the unsubscribe group that you are adding suppressions to.

### Other Parameters

Other parameters are passed through a pointer to a AddSuppressionToAsmGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**SuppressionsRequest** | [**SuppressionsRequest**](SuppressionsRequest.md) | 

### Return type

[**AddSuppressionToAsmGroup201Response**](AddSuppressionToAsmGroup201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

