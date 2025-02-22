# ResendVerifiedSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendVerifiedSender**](ResendVerifiedSender.md#ResendVerifiedSender) | **Post** /v3/verified_senders/resend/{Id} | Resend Verified Sender Request



## ResendVerifiedSender

> ResendVerifiedSender(ctx, Id)

Resend Verified Sender Request

**This endpoint allows you to resend a verification email to a specified Sender Identity**.  Passing the `id` assigned to a Sender Identity to this endpoint will resend a verification email to the `from_address` associated with the Sender Identity. This can be useful if someone loses their verification email or needs to have it resent for any other reason.  You can retrieve the IDs associated with Sender Identities by passing a \"Get All Verified Senders\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ResendVerifiedSenderParams struct


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

