# DeleteFieldDefinition

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFieldDefinition**](DeleteFieldDefinition.md#DeleteFieldDefinition) | **Delete** /v3/marketing/field_definitions/{CustomFieldId} | Delete Custom Field Definition



## DeleteFieldDefinition

> DeleteFieldDefinition(ctx, CustomFieldId)

Delete Custom Field Definition

**This endpoint deletes a defined Custom Field.**  You can delete only Custom Fields; Reserved Fields cannot be deleted.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomFieldId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldDefinitionParams struct


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

