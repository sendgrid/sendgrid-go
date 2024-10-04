# UpdateAccountOffering

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAccountOffering**](UpdateAccountOffering.md#UpdateAccountOffering) | **Put** /v3/partners/accounts/{AccountID}/offerings | Update account offerings



## UpdateAccountOffering

> AccountProvisioningOfferingList UpdateAccountOffering(ctx, AccountIDOfferingsToAdd)

Update account offerings

Changes a package offering for the specified account. Please note that an account can have only one package offering. Also associates one or more add-on offerings such as Marketing Campaigns, Dedicated IP Addresses, and Expert Services to the specified account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountOfferingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AccountProvisioningOfferingList**](AccountProvisioningOfferingList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

