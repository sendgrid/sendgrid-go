# UpdateForwardSpam

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateForwardSpam**](UpdateForwardSpam.md#UpdateForwardSpam) | **Patch** /v3/mail_settings/forward_spam | Update forward spam mail settings



## UpdateForwardSpam

> MailSettingsForwardSpam UpdateForwardSpam(ctx, optional)

Update forward spam mail settings

**This endpoint allows you to update your current Forward Spam mail settings.**  Enabling the Forward Spam setting allows you to specify `email` addresses to which spam reports will be forwarded. You can set multiple addresses by passing this endpoint a comma separated list of emails in a single string.  ``` {   \"email\": \"address1@example.com, address2@exapmle.com\",   \"enabled\": true } ```  The Forward Spam setting may also be used to receive emails sent to `abuse@` and `postmaster@` role addresses if you have authenticated your domain.  For example, if you authenticated `example.com` as your root domain and set a custom return path of `sub` for that domain, you could turn on Forward Spam, and any emails sent to `abuse@sub.example.com` or `postmaster@sub.example.com` would be forwarded to the email address you entered in the `email` field.  You can authenticate your domain using the \"Authenticate a domain\" endpoint or in the [Sender Authentication section of the Twilio SendGrid App](https://app.sendgrid.com/settings/sender_auth). You can also configure the Forward Spam mail settings in the [Mail Settings section of the Twilio SendGrid App](https://app.sendgrid.com/settings/mail_settings).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateForwardSpamParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**MailSettingsForwardSpam** | [**MailSettingsForwardSpam**](MailSettingsForwardSpam.md) | 

### Return type

[**MailSettingsForwardSpam**](MailSettingsForwardSpam.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

