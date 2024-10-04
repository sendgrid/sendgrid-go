# AddIntegration

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIntegration**](AddIntegration.md#AddIntegration) | **Post** /v3/marketing/integrations | CreateIntegration



## AddIntegration

> Integration AddIntegration(ctx, IntegrationInput)

CreateIntegration

This endpoint creates an Integration for email event forwarding. Each Integration has a maximum number of allowed Integration instances per user. For example, users can create up to 10 Segment Integrations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a AddIntegrationParams struct


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

