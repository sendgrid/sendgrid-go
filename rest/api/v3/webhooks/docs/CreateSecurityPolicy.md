# CreateSecurityPolicy

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityPolicy**](CreateSecurityPolicy.md#CreateSecurityPolicy) | **Post** /v3/user/webhooks/security/policies | Create a new webhook security policy



## CreateSecurityPolicy

> CreateSecurityPolicy201Response CreateSecurityPolicy(ctx, optional)

Create a new webhook security policy

Create a new webhook security policy. Note: One of signature or oauth must be given to have a valid security policy. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSecurityPolicyParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**CreateSecurityPolicyRequest** | [**CreateSecurityPolicyRequest**](CreateSecurityPolicyRequest.md) | 

### Return type

[**CreateSecurityPolicy201Response**](CreateSecurityPolicy201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

