# GetEmailJobForVerification

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailJobForVerification**](GetEmailJobForVerification.md#GetEmailJobForVerification) | **Get** /v3/validations/email/jobs/{JobId} | This request returns a single Bulk Email Validation Job.



## GetEmailJobForVerification

> GetValidationsEmailJobsJobId200Response GetEmailJobForVerification(ctx, JobId)

This request returns a single Bulk Email Validation Job.

**This endpoint returns a specific Bulk Email Validation Job. You can use this endpoint to check on the progress of a Job.** 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobId** | **string** | The ID of the Bulk Email Address Validation Job you wish to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a GetEmailJobForVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**GetValidationsEmailJobsJobId200Response**](GetValidationsEmailJobsJobId200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

