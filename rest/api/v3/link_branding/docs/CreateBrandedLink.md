# CreateBrandedLink

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandedLink**](CreateBrandedLink.md#CreateBrandedLink) | **Post** /v3/whitelabel/links | Create a branded link



## CreateBrandedLink

> LinkBranding200 CreateBrandedLink(ctx, optional)

Create a branded link

**This endpoint allows you to create a new branded link.**  To create the link branding, supply the root domain and, optionally, the subdomain — these go into separate fields in your request body. The root domain should match your FROM email address. If you provide a  subdomain, it must be different from the subdomain you used for authenticating your domain.  You can submit this request as one of your subusers if you include their ID in the `on-behalf-of` header in the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandedLinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**CreateBrandedLinkRequest** | [**CreateBrandedLinkRequest**](CreateBrandedLinkRequest.md) | 

### Return type

[**LinkBranding200**](LinkBranding200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

