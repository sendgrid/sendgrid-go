# DisassociateBrandedLinkFromSubuser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisassociateBrandedLinkFromSubuser**](DisassociateBrandedLinkFromSubuser.md#DisassociateBrandedLinkFromSubuser) | **Delete** /v3/whitelabel/links/subuser | Disassociate a branded link from a subuser



## DisassociateBrandedLinkFromSubuser

> map[string]interface{} DisassociateBrandedLinkFromSubuser(ctx, Username)

Disassociate a branded link from a subuser

**This endpoint allows you to take a branded link away from a subuser.**  Link branding can be associated with subusers from the parent account. This functionality allows subusers to send mail using their parent's link branding. To associate link branding, the parent account must first create a branded link and validate it. The parent may then associate that branded link with a subuser via the API or the [Subuser Management page of the Twilio SendGrid App](https://app.sendgrid.com/settings/subusers).  Your request will receive a response with a 204 status code if the disassociation was successful.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DisassociateBrandedLinkFromSubuserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

