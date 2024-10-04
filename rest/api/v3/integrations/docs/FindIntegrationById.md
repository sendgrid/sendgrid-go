# FindIntegrationById

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindIntegrationById**](FindIntegrationById.md#FindIntegrationById) | **Get** /v3/marketing/integrations/{Id} | GetIntegration



## FindIntegrationById

> Integration FindIntegrationById(ctx, Id)

GetIntegration

This endpoint returns the data for a specific Integration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Integration you would like to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a FindIntegrationByIdParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Integration**](Integration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

