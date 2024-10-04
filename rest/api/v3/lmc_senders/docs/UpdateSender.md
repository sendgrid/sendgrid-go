# UpdateSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSender**](UpdateSender.md#UpdateSender) | **Patch** /v3/senders/{SenderId} | Update a Sender Identity



## UpdateSender

> SenderId UpdateSender(ctx, SenderIdoptional)

Update a Sender Identity

**This endpoint allows you to update a sender identity.**  Updates to `from.email` require re-verification.  Partial updates are allowed, but fields that are marked as \"required\" in the POST (create) endpoint must not be nil if that field is included in the PATCH request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SenderId** | **int32** | The ID of the sender identity that you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**SenderIdRequest** | [**SenderIdRequest**](SenderIdRequest.md) | 

### Return type

[**SenderId**](SenderId.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

