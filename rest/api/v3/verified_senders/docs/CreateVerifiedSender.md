# CreateVerifiedSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerifiedSender**](CreateVerifiedSender.md#CreateVerifiedSender) | **Post** /v3/verified_senders | Create Verified Sender Request



## CreateVerifiedSender

> VerifiedSenderResponse CreateVerifiedSender(ctx, optional)

Create Verified Sender Request

**This endpoint allows you to create a new Sender Identify**.  Upon successful submission of a `POST` request to this endpoint, an identity will be created, and a verification email will be sent to the address assigned to the `from_email` field. You must complete the verification process using the sent email to fully verify the sender.  If you need to resend the verification email, you can do so with the Resend Verified Sender Request, `/resend/{id}`, endpoint.  If you need to authenticate a domain rather than a Single Sender, see the [Domain Authentication API](https://docs.sendgrid.com/api-reference/domain-authentication/authenticate-a-domain).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateVerifiedSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**VerifiedSenderRequest** | [**VerifiedSenderRequest**](VerifiedSenderRequest.md) | 

### Return type

[**VerifiedSenderResponse**](VerifiedSenderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

