# ListSsoIntegrationCertificate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSsoIntegrationCertificate**](ListSsoIntegrationCertificate.md#ListSsoIntegrationCertificate) | **Get** /v3/sso/integrations/{IntegrationId}/certificates | Get All SSO Certificates by Integration



## ListSsoIntegrationCertificate

> []SsoCertificateBody ListSsoIntegrationCertificate(ctx, IntegrationId)

Get All SSO Certificates by Integration

**This endpoint allows you to retrieve all your IdP configurations by configuration ID.**  The `integration_id` expected by this endpoint is the `id` returned in the response by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IntegrationId** | **string** | An ID that matches a certificate to a specific IdP integration.

### Other Parameters

Other parameters are passed through a pointer to a ListSsoIntegrationCertificateParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**[]SsoCertificateBody**](SsoCertificateBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

