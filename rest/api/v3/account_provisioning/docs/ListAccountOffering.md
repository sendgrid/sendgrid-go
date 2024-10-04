# ListAccountOffering

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountOffering**](ListAccountOffering.md#ListAccountOffering) | **Get** /v3/partners/accounts/{AccountID}/offerings | Get account offerings



## ListAccountOffering

> AccountProvisioningOfferingList ListAccountOffering(ctx, AccountID)

Get account offerings

Retrieves offering information about the specified account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a ListAccountOfferingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AccountProvisioningOfferingList**](AccountProvisioningOfferingList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

