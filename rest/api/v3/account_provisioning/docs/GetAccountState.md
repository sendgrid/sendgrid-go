# GetAccountState

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountState**](GetAccountState.md#GetAccountState) | **Get** /v3/partners/accounts/{AccountID}/state | Get an account&#39;s state



## GetAccountState

> AccountProvisioningStateRead GetAccountState(ctx, AccountID)

Get an account's state

Retrieve the state of the specified account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a GetAccountStateParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AccountProvisioningStateRead**](AccountProvisioningStateRead.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

