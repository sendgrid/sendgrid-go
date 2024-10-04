# ListEmailJobForVerification

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmailJobForVerification**](ListEmailJobForVerification.md#ListEmailJobForVerification) | **Put** /v3/validations/email/jobs | Request a presigned URL and headers for Bulk Email Address Validation list upload.



## ListEmailJobForVerification

> PutValidationsEmailJobs200Response ListEmailJobForVerification(ctx, optional)

Request a presigned URL and headers for Bulk Email Address Validation list upload.

**This endpoint returns a presigned URL and request headers. Use this information to upload a list of email addresses for verification.**  Note that in a successful response the `content-type` header value matches the provided `file_type` parameter in the `PUT` request.  Once you have an `upload_uri` and the `upload_headers`, you're ready to upload your email address list for verification. For the expected format of the email address list and a sample upload request, see the [Bulk Email Address Validation Overview page](https://www.twilio.com/docs/sendgrid/ui/managing-contacts/email-address-validation/bulk-email-address-validation-overview). 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEmailJobForVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ListEmailJobForVerificationRequest** | [**ListEmailJobForVerificationRequest**](ListEmailJobForVerificationRequest.md) | 

### Return type

[**PutValidationsEmailJobs200Response**](PutValidationsEmailJobs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

