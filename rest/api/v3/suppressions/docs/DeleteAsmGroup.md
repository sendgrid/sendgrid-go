# DeleteAsmGroup

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAsmGroup**](DeleteAsmGroup.md#DeleteAsmGroup) | **Delete** /v3/asm/groups/{GroupId} | Delete a Suppression Group



## DeleteAsmGroup

> map[string]interface{} DeleteAsmGroup(ctx, GroupIdoptional)

Delete a Suppression Group

**This endpoint allows you to delete a suppression group.**  If a recipient uses the \"one-click unsubscribe\" option on an email associated with a deleted group, that recipient will be added to the global suppression list.  Deleting a suppression group will remove the suppression, meaning email will once again be sent to the previously suppressed addresses. This should be avoided unless a recipient indicates they wish to receive email from you again. You can use our [bypass filters](https://sendgrid.com/docs/ui/sending-email/index-suppressions/#bypass-suppressions) to deliver messages to otherwise suppressed addresses when exceptions are required.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**GroupId** | **string** | The ID of the suppression group you would like to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAsmGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

