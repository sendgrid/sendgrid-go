# UpdateForwardBounce

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateForwardBounce**](UpdateForwardBounce.md#UpdateForwardBounce) | **Patch** /v3/mail_settings/forward_bounce | Update forward bounce mail settings



## UpdateForwardBounce

> MailSettingsForwardBounce UpdateForwardBounce(ctx, optional)

Update forward bounce mail settings

**This endpoint allows you to update your current bounce forwarding mail settings.**  Enabling the Forward Bounce setting allows you to specify an `email` address to which bounce reports will be forwarded.  You can also configure the Forward Spam mail settings in the [Mail Settings section of the Twilio SendGrid App](https://app.sendgrid.com/settings/mail_settings).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateForwardBounceParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**MailSettingsForwardBounce** | [**MailSettingsForwardBounce**](MailSettingsForwardBounce.md) | 

### Return type

[**MailSettingsForwardBounce**](MailSettingsForwardBounce.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

