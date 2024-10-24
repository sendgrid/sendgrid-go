# ActivateTemplateVersion

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTemplateVersion**](ActivateTemplateVersion.md#ActivateTemplateVersion) | **Post** /v3/templates/{TemplateId}/versions/{VersionId}/activate | Activate a transactional template version.



## ActivateTemplateVersion

> TransactionalTemplateVersionOutput ActivateTemplateVersion(ctx, TemplateIdVersionIdoptional)

Activate a transactional template version.

**This endpoint allows you to activate a version of one of your templates.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TemplateId** | **string** | The ID of the original template
**VersionId** | **string** | The ID of the template version

### Other Parameters

Other parameters are passed through a pointer to a ActivateTemplateVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**TransactionalTemplateVersionOutput**](TransactionalTemplateVersionOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
