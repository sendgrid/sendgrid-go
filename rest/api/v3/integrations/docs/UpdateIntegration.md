# UpdateIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIntegration**](UpdateIntegration.md#UpdateIntegration) | **Patch** /v3/marketing/integrations/{Id} | UpdateIntegration



## UpdateIntegration

> Integration UpdateIntegration(ctx, IdIntegrationPatch)

UpdateIntegration

This endpoint updates an existing Integration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Integration you would like to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIntegrationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Integration**](Integration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

