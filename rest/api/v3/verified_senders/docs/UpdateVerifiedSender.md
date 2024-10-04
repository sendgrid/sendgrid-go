# UpdateVerifiedSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateVerifiedSender**](UpdateVerifiedSender.md#UpdateVerifiedSender) | **Patch** /v3/verified_senders/{Id} | Edit Verified Sender



## UpdateVerifiedSender

> VerifiedSenderResponse UpdateVerifiedSender(ctx, Idoptional)

Edit Verified Sender

**This endpoint allows you to update an existing Sender Identity**.  Pass the `id` assigned to a Sender Identity to this endpoint as a path parameter. Include any fields you wish to update in the request body in JSON format.  You can retrieve the IDs associated with Sender Identities by passing a `GET` request to the Get All Verified Senders endpoint, `/verified_senders`.  **Note:** Unlike a `PUT` request, `PATCH` allows you to update only the fields you wish to edit. Fields that are not passed as part of a request will remain unaltered.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateVerifiedSenderParams struct


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

