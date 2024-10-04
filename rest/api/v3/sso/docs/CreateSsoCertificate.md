# CreateSsoCertificate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSsoCertificate**](CreateSsoCertificate.md#CreateSsoCertificate) | **Post** /v3/sso/certificates | Create an SSO Certificate



## CreateSsoCertificate

> SsoCertificateBody CreateSsoCertificate(ctx, optional)

Create an SSO Certificate

**This endpoint allows you to create an SSO certificate.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSsoCertificateParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateSsoCertificateRequest** | [**CreateSsoCertificateRequest**](CreateSsoCertificateRequest.md) | 

### Return type

[**SsoCertificateBody**](SsoCertificateBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

