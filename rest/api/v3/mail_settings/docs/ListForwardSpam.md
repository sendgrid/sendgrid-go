# ListForwardSpam

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListForwardSpam**](ListForwardSpam.md#ListForwardSpam) | **Get** /v3/mail_settings/forward_spam | Retrieve forward spam mail settings



## ListForwardSpam

> MailSettingsForwardSpam ListForwardSpam(ctx, optional)

Retrieve forward spam mail settings

**This endpoint allows you to retrieve your current Forward Spam mail settings.**  Enabling the Forward Spam setting allows you to specify `email` addresses to which spam reports will be forwarded. This endpoint returns any email address(es) you have set to receive forwarded spam and an `enabled` status indicating if the setting is active.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListForwardSpamParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**MailSettingsForwardSpam**](MailSettingsForwardSpam.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

