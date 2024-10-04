# ListApiKey

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiKey**](ListApiKey.md#ListApiKey) | **Get** /v3/api_keys | Retrieve all API Keys belonging to the authenticated user



## ListApiKey

> ListApiKey200Response ListApiKey(ctx, optional)

Retrieve all API Keys belonging to the authenticated user

**This endpoint allows you to retrieve all API Keys that belong to the authenticated user.**  A successful response from this API will include all available API keys' names and IDs.  For security reasons, there is not a way to retrieve the key itself after it's created. If you lose your API key, you must create a new one. Only the \"Create API keys\" endpoint will return a key to you and only at the time of creation.  An `api_key_id` can be used to update or delete the key, as well as retrieve the key's details, such as its scopes.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListApiKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int32** | 
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ListApiKey200Response**](ListApiKey200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

