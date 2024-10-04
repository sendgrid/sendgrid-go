# UpdateSubuserCredit

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubuserCredit**](UpdateSubuserCredit.md#UpdateSubuserCredit) | **Put** /v3/subusers/{SubuserName}/credits | Update the Credits for a Subuser



## UpdateSubuserCredit

> SubuserCredits UpdateSubuserCredit(ctx, SubuserNameSubuserCreditsRequest)

Update the Credits for a Subuser

**This endpoint allows you to update the Credits for a Subuser.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubuserCreditParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SubuserCredits**](SubuserCredits.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

