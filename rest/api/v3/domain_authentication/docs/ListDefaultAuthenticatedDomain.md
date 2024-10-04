# ListDefaultAuthenticatedDomain

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDefaultAuthenticatedDomain**](ListDefaultAuthenticatedDomain.md#ListDefaultAuthenticatedDomain) | **Get** /v3/whitelabel/domains/default | Get the default authentication



## ListDefaultAuthenticatedDomain

> []interface{} ListDefaultAuthenticatedDomain(ctx, optional)

Get the default authentication

**This endpoint allows you to retrieve the default authentication for a domain.**  When creating or updating a domain authentication, you can set the domain as a default. The default domain will be used to send all mail. If you have multiple authenticated domains, the authenticated domain matching the domain of the From address will be used, and the default will be overridden.  This endpoint will return a default domain and its details only if a default is set. You are not required to set a default. If you do not set a default domain, this endpoint will return general information about your domain authentication status.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDefaultAuthenticatedDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**Domain** | **string** | The domain to find a default authentication.
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

**[]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

