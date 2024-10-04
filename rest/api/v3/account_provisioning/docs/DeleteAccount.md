# DeleteAccount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](DeleteAccount.md#DeleteAccount) | **Delete** /v3/partners/accounts/{AccountID} | Delete an account



## DeleteAccount

> DeleteAccount(ctx, AccountID)

Delete an account

Delete a specific account under your organization by account ID. Note that this is an **irreversible** action that does the following:   - Revokes API Keys and SSO so that the account user cannot log in or access SendGrid data.  - Removes all offerings and configured SendGrid resources such as dedicated IPs.  - Cancels billing effective immediately.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a DeleteAccountParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

