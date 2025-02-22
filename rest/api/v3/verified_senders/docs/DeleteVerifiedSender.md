# DeleteVerifiedSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVerifiedSender**](DeleteVerifiedSender.md#DeleteVerifiedSender) | **Delete** /v3/verified_senders/{Id} | Delete Verified Sender



## DeleteVerifiedSender

> DeleteVerifiedSender(ctx, Id)

Delete Verified Sender

**This endpoint allows you to delete a Sender Identity**.  Pass the `id` assigned to a Sender Identity to this endpoint to delete the Sender Identity from your account.  You can retrieve the IDs associated with Sender Identities using the \"Get All Verified Senders\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteVerifiedSenderParams struct


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

