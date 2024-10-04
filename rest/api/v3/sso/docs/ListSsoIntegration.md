# ListSsoIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSsoIntegration**](ListSsoIntegration.md#ListSsoIntegration) | **Get** /v3/sso/integrations | Get All SSO Integrations



## ListSsoIntegration

> []SsoIntegration ListSsoIntegration(ctx, optional)

Get All SSO Integrations

**This endpoint allows you to retrieve all SSO integrations tied to your Twilio SendGrid account.**  The IDs returned by this endpoint can be used by the APIs additional endpoints to modify your SSO integrations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSsoIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Si** | **bool** | If this parameter is set to `true`, the response will include the `completed_integration` field.

### Return type

[**[]SsoIntegration**](SsoIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

