# UpdateTemplate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateTemplate**](UpdateTemplate.md#UpdateTemplate) | **Patch** /v3/templates/{TemplateId} | Edit a transactional template.



## UpdateTemplate

> TransactionalTemplate UpdateTemplate(ctx, TemplateIdoptional)

Edit a transactional template.

**This endpoint allows you to edit the name of a transactional template.**  To edit the template itself, [create a new transactional template version](https://docs.sendgrid.com/api-reference/transactional-templates-versions/create-a-new-transactional-template-version).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TemplateId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateTemplateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateTemplateRequest** | [**UpdateTemplateRequest**](UpdateTemplateRequest.md) | 

### Return type

[**TransactionalTemplate**](TransactionalTemplate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
