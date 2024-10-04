# UpdateFieldDefinition

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateFieldDefinition**](UpdateFieldDefinition.md#UpdateFieldDefinition) | **Patch** /v3/marketing/field_definitions/{CustomFieldId} | Update Custom Field Definition



## UpdateFieldDefinition

> CreateFieldDefinition200Response UpdateFieldDefinition(ctx, CustomFieldIdoptional)

Update Custom Field Definition

**This endpoint allows you to update a defined Custom Field.**  Only your Custom fields can be modified; Reserved Fields cannot be updated.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomFieldId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateFieldDefinitionParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateFieldDefinitionRequest** | [**UpdateFieldDefinitionRequest**](UpdateFieldDefinitionRequest.md) | 

### Return type

[**CreateFieldDefinition200Response**](CreateFieldDefinition200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

