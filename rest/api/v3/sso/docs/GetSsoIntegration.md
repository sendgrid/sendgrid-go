# GetSsoIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSsoIntegration**](GetSsoIntegration.md#GetSsoIntegration) | **Get** /v3/sso/integrations/{Id} | Get an SSO Integration



## GetSsoIntegration

> SsoIntegration GetSsoIntegration(ctx, Idoptional)

Get an SSO Integration

**This endpoint allows you to retrieve an SSO integration by ID.**  You can retrieve the IDs for your configurations from the response provided by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetSsoIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Si** | **bool** | If this parameter is set to `true`, the response will include the `completed_integration` field.

### Return type

[**SsoIntegration**](SsoIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

