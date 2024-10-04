# GetSsoCertificate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSsoCertificate**](GetSsoCertificate.md#GetSsoCertificate) | **Get** /v3/sso/certificates/{CertId} | Get an SSO Certificate



## GetSsoCertificate

> SsoCertificateBody GetSsoCertificate(ctx, CertId)

Get an SSO Certificate

**This endpoint allows you to retrieve an individual SSO certificate.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CertId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetSsoCertificateParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SsoCertificateBody**](SsoCertificateBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

