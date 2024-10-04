# GetMailBatch

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMailBatch**](GetMailBatch.md#GetMailBatch) | **Get** /v3/mail/batch/{BatchId} | Validate a batch ID.



## GetMailBatch

> MailBatchResponse GetMailBatch(ctx, optional)

Validate a batch ID.

**This operation allows you to validate a mail batch ID.**  If you provide a valid batch ID, this operation will return a `200` status code and the batch ID itself. If you provide an invalid batch ID, you will receive a `400` level status code and an error message. A batch ID does not need to be assigned to a send to be considered valid. A successful response means only that the batch ID has been created, but it does not indicate that the ID has been assigned to a send.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a GetMailBatchParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | Use the `on-behalf-of` header to make API calls for a particular Subuser through the parent account. You can use this header to automate bulk updates or to administer a Subuser without changing the authentication in your code. You will use the parent account's API key when using this header.

### Return type

[**MailBatchResponse**](MailBatchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

