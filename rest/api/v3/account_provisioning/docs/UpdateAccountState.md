# UpdateAccountState

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAccountState**](UpdateAccountState.md#UpdateAccountState) | **Put** /v3/partners/accounts/{AccountID}/state | Update an account&#39;s state



## UpdateAccountState

> UpdateAccountState(ctx, AccountIDAccountProvisioningStateWrite)

Update an account's state

Update the state of the specified account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountStateParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

