# ListScope

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListScope**](ListScope.md#ListScope) | **Get** /v3/scopes | Retrieve a list of scopes for which this user has access.



## ListScope

> ListScope200Response ListScope(ctx, optional)

Retrieve a list of scopes for which this user has access.

**This endpoint returns a list of all scopes that this user has access to.**  API Keys are used to authenticate with [SendGrid's v3 API](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/authorization).  API Keys may be assigned certain permissions, or scopes, that limit which API endpoints they are able to access.  This endpoint returns all the scopes assigned to the key you use to authenticate with it. To retrieve the scopes assigned to another key, you can pass an API key ID to the \"Retrieve an existing API key\" endpoint.  For a more detailed explanation of how you can use API Key permissions, please visit our [API Keys documentation](https://sendgrid.com/docs/ui/account-and-settings/api-keys/).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListScopeParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ListScope200Response**](ListScope200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

