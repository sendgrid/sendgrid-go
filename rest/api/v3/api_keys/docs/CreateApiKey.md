# CreateApiKey

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](CreateApiKey.md#CreateApiKey) | **Post** /v3/api_keys | Create API keys



## CreateApiKey

> CreateApiKey201Response CreateApiKey(ctx, optional)

Create API keys

**This endpoint allows you to create a new API Key for the user.**  To create your initial SendGrid API Key, you should [use the SendGrid App](https://app.sendgrid.com/settings/api_keys). Once you have created a first key with scopes to manage additional API keys, you can use this API for all other key management. A JSON request body containing a `name` property is required when making requests to this endpoint. If the number of maximum keys, 100, is reached, a `403` status will be returned. Though the `name` field is required, it does not need to be unique. A unique API key ID will be generated for each key you create and returned in the response body. It is not necessary to pass a `scopes` field to the API when creating a key, but you should be aware that omitting the `scopes` field from your request will create a key with \"Full Access\" permissions by default. See the [API Key Permissions List](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/authorization) for all available scopes. An API key's scopes can be updated after creation using the \"Update API keys\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateApiKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**CreateApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) | 

### Return type

[**CreateApiKey201Response**](CreateApiKey201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

