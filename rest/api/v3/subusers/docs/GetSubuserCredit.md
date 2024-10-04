# GetSubuserCredit

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubuserCredit**](GetSubuserCredit.md#GetSubuserCredit) | **Get** /v3/subusers/{SubuserName}/credits | Get the Credits for a Subuser



## GetSubuserCredit

> SubuserCredits GetSubuserCredit(ctx, SubuserName)

Get the Credits for a Subuser

**This endpoint allows you to retrieve a Credits overview for a Subuser.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a GetSubuserCreditParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SubuserCredits**](SubuserCredits.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

