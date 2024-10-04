# ValidateEmail

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateEmail**](ValidateEmail.md#ValidateEmail) | **Post** /v3/validations/email | Validate an email



## ValidateEmail

> ValidateEmail200Response ValidateEmail(ctx, optional)

Validate an email

**This endpoint allows you to validate an email address.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ValidateEmailParams struct


Name | Type | Description
------------- | ------------- | -------------
**ValidateEmailRequest** | [**ValidateEmailRequest**](ValidateEmailRequest.md) | 

### Return type

[**ValidateEmail200Response**](ValidateEmail200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

