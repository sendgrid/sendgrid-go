# AuthenticateAccount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateAccount**](AuthenticateAccount.md#AuthenticateAccount) | **Post** /v3/partners/accounts/{AccountID}/sso | Authenticate an account with single sign on



## AuthenticateAccount

> AuthenticateAccount(ctx, AccountID)

Authenticate an account with single sign on

Authenticates and logs in a user to Twilio Sendgrid as a specific admin identity configured for SSO by partner. Any additional teammates or subusers will need to log in directly via app.sendgrid.com

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountID** | **string** | Twilio SendGrid account ID

### Other Parameters

Other parameters are passed through a pointer to a AuthenticateAccountParams struct


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

