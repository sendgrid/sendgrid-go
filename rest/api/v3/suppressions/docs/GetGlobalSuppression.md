# GetGlobalSuppression

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalSuppression**](GetGlobalSuppression.md#GetGlobalSuppression) | **Get** /v3/asm/suppressions/global/{Email} | Retrieve a Global Suppression



## GetGlobalSuppression

> RetrieveAGlobalSuppressionResponse GetGlobalSuppression(ctx, Emailoptional)

Retrieve a Global Suppression

**This endpoint allows you to retrieve a global suppression. You can also use this endpoint to confirm if an email address is already globally suppresed.**  If the email address you include in the URL path parameter `{email}` is already globally suppressed, the response will include that email address. If the address you enter for `{email}` is not globally suppressed, an empty JSON object `{}` will be returned.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Email** | **string** | The email address of the global suppression you want to retrieve. Or, if you want to check if an email address is on the global suppressions list, enter that email address here.

### Other Parameters

Other parameters are passed through a pointer to a GetGlobalSuppressionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**RetrieveAGlobalSuppressionResponse**](RetrieveAGlobalSuppressionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

