# ListVerifiedSenderStepsCompleted

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVerifiedSenderStepsCompleted**](ListVerifiedSenderStepsCompleted.md#ListVerifiedSenderStepsCompleted) | **Get** /v3/verified_senders/steps_completed | Completed Steps



## ListVerifiedSenderStepsCompleted

> ListVerifiedSenderStepsCompleted200Response ListVerifiedSenderStepsCompleted(ctx, )

Completed Steps

**This endpoint allows you to determine which of SendGrid’s verification processes have been completed for an account**.  This endpoint returns boolean values, `true` and `false`, for [Domain Authentication](https://sendgrid.com/docs/for-developers/sending-email/sender-identity/#domain-authentication), `domain_verified`, and [Single Sender Verification](https://sendgrid.com/docs/for-developers/sending-email/sender-identity/#single-sender-verification), `sender_verified`, for the account.  An account may have one, both, or neither verification steps completed. If you need to authenticate a domain rather than a Single Sender, see the \"Authenticate a domain\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerifiedSenderStepsCompletedParams struct


### Return type

[**ListVerifiedSenderStepsCompleted200Response**](ListVerifiedSenderStepsCompleted200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

