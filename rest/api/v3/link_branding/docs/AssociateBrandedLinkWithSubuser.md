# AssociateBrandedLinkWithSubuser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateBrandedLinkWithSubuser**](AssociateBrandedLinkWithSubuser.md#AssociateBrandedLinkWithSubuser) | **Post** /v3/whitelabel/links/{LinkId}/subuser | Associate a branded link with a subuser



## AssociateBrandedLinkWithSubuser

> LinkBranding200 AssociateBrandedLinkWithSubuser(ctx, LinkIdoptional)

Associate a branded link with a subuser

**This endpoint allows you to associate a branded link with a subuser account.**  Link branding can be associated with subusers from the parent account. This functionality allows subusers to send mail using their parent's link branding. To associate link branding, the parent account must first create a branded link and validate it. The parent may then associate that branded link with a subuser via the API or the [Subuser Management page of the Twilio SendGrid App](https://app.sendgrid.com/settings/subusers).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**LinkId** | **int32** | The ID of the branded link you want to associate.

### Other Parameters

Other parameters are passed through a pointer to a AssociateBrandedLinkWithSubuserParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssociateBrandedLinkWithSubuserRequest** | [**AssociateBrandedLinkWithSubuserRequest**](AssociateBrandedLinkWithSubuserRequest.md) | 

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

