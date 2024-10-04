# ListSubuserBrandedLink

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubuserBrandedLink**](ListSubuserBrandedLink.md#ListSubuserBrandedLink) | **Get** /v3/whitelabel/links/subuser | Retrieve a subuser&#39;s branded link



## ListSubuserBrandedLink

> LinkBranding200 ListSubuserBrandedLink(ctx, Username)

Retrieve a subuser's branded link

**This endpoint allows you to retrieve the branded link associated with a subuser.**  Link branding can be associated with subusers from the parent account. This functionality allows subusers to send mail using their parent's link branding. To associate link branding, the parent account must first create a branded link and then validate it. The parent may then associate that branded link with a subuser via the API or the [Subuser Management page of the Twilio SendGrid App](https://app.sendgrid.com/settings/subusers).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubuserBrandedLinkParams struct


Name | Type | Description
------------- | ------------- | -------------

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

