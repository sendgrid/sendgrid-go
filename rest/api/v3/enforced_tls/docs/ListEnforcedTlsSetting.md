# ListEnforcedTlsSetting

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEnforcedTlsSetting**](ListEnforcedTlsSetting.md#ListEnforcedTlsSetting) | **Get** /v3/user/settings/enforced_tls | Retrieve current Enforced TLS settings.



## ListEnforcedTlsSetting

> EnforcedTlsRequestResponse ListEnforcedTlsSetting(ctx, optional)

Retrieve current Enforced TLS settings.

**This endpoint allows you to retrieve your current Enforced TLS settings.**  The Enforced TLS settings specify whether or not the recipient is required to support TLS or have a valid certificate.  If either `require_tls` or `require_valid_cert` is set to `true`, the recipient must support TLS 1.1 or higher or have a valid certificate. If these conditions are not met, Twilio SendGrid will drop the message and send a block event with “TLS required but not supported” as the description.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEnforcedTlsSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**EnforcedTlsRequestResponse**](EnforcedTlsRequestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

