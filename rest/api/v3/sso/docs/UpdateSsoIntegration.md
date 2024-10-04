# UpdateSsoIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSsoIntegration**](UpdateSsoIntegration.md#UpdateSsoIntegration) | **Patch** /v3/sso/integrations/{Id} | Update an SSO Integration



## UpdateSsoIntegration

> SsoIntegration UpdateSsoIntegration(ctx, Idoptional)

Update an SSO Integration

**This endpoint allows you to modify an exisiting SSO integration.**  You can retrieve the IDs for your configurations from the response provided by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSsoIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Si** | **bool** | If this parameter is set to `true`, the response will include the `completed_integration` field.
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

