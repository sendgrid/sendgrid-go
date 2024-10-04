# DeleteSsoCertificate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSsoCertificate**](DeleteSsoCertificate.md#DeleteSsoCertificate) | **Delete** /v3/sso/certificates/{CertId} | Delete an SSO Certificate



## DeleteSsoCertificate

> SsoCertificateBody DeleteSsoCertificate(ctx, CertId)

Delete an SSO Certificate

**This endpoint allows you to delete an SSO certificate.**  You can retrieve a certificate's ID from the response provided by the \"Get All SSO Integrations\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CertId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSsoCertificateParams struct


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

