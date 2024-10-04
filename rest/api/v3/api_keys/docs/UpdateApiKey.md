# UpdateApiKey

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateApiKey**](UpdateApiKey.md#UpdateApiKey) | **Put** /v3/api_keys/{ApiKeyId} | Update API key name and scopes



## UpdateApiKey

> ApiKeyResponse UpdateApiKey(ctx, ApiKeyIdoptional)

Update API key name and scopes

**This endpoint allows you to update the name and scopes of a given API key.**  You must pass this endpoint a JSON request body with a `name` field and a `scopes` array containing at least one scope. The `name` and `scopes` fields will be used to update the key associated with the `api_key_id` in the request URL.  If you need to update a key's scopes only, pass the `name` field with the key's existing name; the `name` will not be modified. If you need to update a key's name only, use the \"Update API key name\" endpoint.  See the [API Key Permissions List](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/authorization) for all available scopes.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ApiKeyId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateApiKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateApiKeyRequest** | [**UpdateApiKeyRequest**](UpdateApiKeyRequest.md) | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

