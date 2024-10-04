# CreateAccount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](CreateAccount.md#CreateAccount) | **Post** /v3/partners/accounts | Create an account



## CreateAccount

> AccountProvisioningAccountId CreateAccount(ctx, CreateAccountParamsoptional)

Create an account

Creates a new account, with specified offering, under the organization.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**TTestAccount** | **string** | **OPTIONAL** Custom request header provided ONLY for a test account

### Return type

[**AccountProvisioningAccountId**](AccountProvisioningAccountId.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

