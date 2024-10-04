# ListTemplate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTemplate**](ListTemplate.md#ListTemplate) | **Get** /v3/templates | Retrieve paged transactional templates.



## ListTemplate

> ListTemplate200Response ListTemplate(ctx, PageSizeoptional)

Retrieve paged transactional templates.

**This endpoint allows you to retrieve all transactional templates.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTemplateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Generations** | [**Generations**](GenerationsGenerations.md) | Comma-delimited list specifying which generations of templates to return. Options are `legacy`, `dynamic` or `legacy,dynamic`.
**PageToken** | **string** | A token corresponding to a specific page of results, as provided by metadata
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ListTemplate200Response**](ListTemplate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

