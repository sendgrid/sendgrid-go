# ResendTeammateInvite

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendTeammateInvite**](ResendTeammateInvite.md#ResendTeammateInvite) | **Post** /v3/teammates/pending/{Token}/resend | Resend teammate invite



## ResendTeammateInvite

> ResendTeammateInvite200Response ResendTeammateInvite(ctx, Tokenoptional)

Resend teammate invite

**This endpoint allows you to resend a Teammate invitation.**  Teammate invitations will expire after 7 days. Resending an invitation will reset the expiration date.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Token** | **string** | The token for the invite that you want to resend.

### Other Parameters

Other parameters are passed through a pointer to a ResendTeammateInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ResendTeammateInvite200Response**](ResendTeammateInvite200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

