# CreateFieldDefinition

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFieldDefinition**](CreateFieldDefinition.md#CreateFieldDefinition) | **Post** /v3/marketing/field_definitions | Create Custom Field Definition



## CreateFieldDefinition

> CreateFieldDefinition200Response CreateFieldDefinition(ctx, optional)

Create Custom Field Definition

**This endpoint creates a new custom field definition.**  Custom field definitions are created with the given `name` and `field_type`. Although field names are stored in a case-sensitive manner, all field names must be case-insensitively unique. This means you may create a field named `CamelCase` or `camelcase`, but not both. Additionally, a Custom Field name cannot collide with any Reserved Field names. You should save the returned `id` value in order to update or delete the field at a later date. You can have up to 500 custom fields.  The custom field name should be created using only alphanumeric characters (A-Z and 0-9) and underscores (\\_). Custom fields can only begin with letters  A-Z or underscores (_). The field type can be date, text, or number fields. The field type is important for creating segments from your contact database.  **Note: Creating a custom field that begins with a number will cause issues with sending in Marketing Campaigns.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldDefinitionParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateFieldDefinitionRequest** | [**CreateFieldDefinitionRequest**](CreateFieldDefinitionRequest.md) | 

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

