# ListDefaultBrandedLink

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDefaultBrandedLink**](ListDefaultBrandedLink.md#ListDefaultBrandedLink) | **Get** /v3/whitelabel/links/default | Retrieve the default branded link



## ListDefaultBrandedLink

> LinkBranding200 ListDefaultBrandedLink(ctx, optional)

Retrieve the default branded link

**This endpoint allows you to retrieve the default branded link.**  The default branded link is the actual URL to be used when sending messages. If you have more than one branded link, the default is determined by the following order:  * The validated branded link marked as `default` (set when you call the \"Create a branded link\" endpoint or by calling the \"Update a branded link\" endpoint on an existing link) * Legacy branded links (migrated from the whitelabel wizard) * Default SendGrid-branded links (i.e., `100.ct.sendgrid.net`)  You can submit this request as one of your subusers if you include their ID in the `on-behalf-of` header in the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDefaultBrandedLinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Domain** | **string** | The domain to match against when finding the default branded link.
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**LinkBranding200**](LinkBranding200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

