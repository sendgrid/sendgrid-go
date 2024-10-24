# UpdateFooter

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateFooter**](UpdateFooter.md#UpdateFooter) | **Patch** /v3/mail_settings/footer | Update footer mail settings



## UpdateFooter

> MailSettingsFooter UpdateFooter(ctx, optional)

Update footer mail settings

**This endpoint allows you to update your current Footer mail settings.**  The Footer setting will insert a custom footer at the bottom of your text and HTML email message bodies.  You can insert your HTML or plain text directly using this endpoint, or you can create the footer using the [Mail Settings menu in the Twilio SendGrid App](https://app.sendgrid.com/settings/mail_settings).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFooterParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**MailSettingsFooter** | [**MailSettingsFooter**](MailSettingsFooter.md) | 

### Return type

[**MailSettingsFooter**](MailSettingsFooter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
