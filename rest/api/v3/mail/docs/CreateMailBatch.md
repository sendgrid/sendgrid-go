# CreateMailBatch

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailBatch**](CreateMailBatch.md#CreateMailBatch) | **Post** /v3/mail/batch | Create a batch ID.



## CreateMailBatch

> MailBatchResponse CreateMailBatch(ctx, optional)

Create a batch ID.

**This operation allows you to generate a new mail batch ID.**   Once a batch ID is created, you can associate it with a mail send by passing it in the request body of the [Mail Send operation](https://docs.sendgrid.com/api-reference/mail-send/mail-send). This makes it possible to group multiple requests to the Mail Send operation by assigning them the same batch ID.   A batch ID that's associated with a mail send can be used to access and modify the associated send. For example, you can pause or cancel a send using its batch ID. See the [Scheduled Sends API](https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends) for more information about pausing and cancelling a mail send. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateMailBatchParams struct


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

