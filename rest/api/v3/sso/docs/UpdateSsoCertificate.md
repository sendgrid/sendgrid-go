# UpdateSsoCertificate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSsoCertificate**](UpdateSsoCertificate.md#UpdateSsoCertificate) | **Patch** /v3/sso/certificates/{CertId} | Update SSO Certificate



## UpdateSsoCertificate

> SsoCertificateBody UpdateSsoCertificate(ctx, CertIdoptional)

Update SSO Certificate

**This endpoint allows you to update an existing certificate by ID.**  You can retrieve a certificate's ID from the response provided by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CertId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSsoCertificateParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateSsoCertificateRequest** | [**UpdateSsoCertificateRequest**](UpdateSsoCertificateRequest.md) | 

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

