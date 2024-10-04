# UpdateSubuserRemainingCredit

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubuserRemainingCredit**](UpdateSubuserRemainingCredit.md#UpdateSubuserRemainingCredit) | **Patch** /v3/subusers/{SubuserName}/credits/remaining | Update the remaining credits for a Subuser



## UpdateSubuserRemainingCredit

> SubuserCredits UpdateSubuserRemainingCredit(ctx, SubuserNameUpdateSubuserRemainingCreditRequest)

Update the remaining credits for a Subuser

**This endpoint allows you to update the remaining credits for a Subuser.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubuserRemainingCreditParams struct


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

