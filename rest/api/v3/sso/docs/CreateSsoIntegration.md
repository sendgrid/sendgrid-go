# CreateSsoIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSsoIntegration**](CreateSsoIntegration.md#CreateSsoIntegration) | **Post** /v3/sso/integrations | Create an SSO Integration



## CreateSsoIntegration

> SsoIntegration CreateSsoIntegration(ctx, optional)

Create an SSO Integration

**This endpoint allows you to create an SSO integration.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSsoIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PostPatchIntegrationRequest** | [**PostPatchIntegrationRequest**](PostPatchIntegrationRequest.md) | 

### Return type

[**SsoIntegration**](SsoIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

