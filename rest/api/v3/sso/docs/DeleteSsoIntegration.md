# DeleteSsoIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSsoIntegration**](DeleteSsoIntegration.md#DeleteSsoIntegration) | **Delete** /v3/sso/integrations/{Id} | Delete an SSO Integration



## DeleteSsoIntegration

> DeleteSsoIntegration(ctx, Id)

Delete an SSO Integration

**This endpoint allows you to delete an IdP configuration by ID.**  You can retrieve the IDs for your configurations from the response provided by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSsoIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

