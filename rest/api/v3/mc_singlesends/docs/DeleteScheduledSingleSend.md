# DeleteScheduledSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScheduledSingleSend**](DeleteScheduledSingleSend.md#DeleteScheduledSingleSend) | **Delete** /v3/marketing/singlesends/{Id}/schedule | Delete Single Send Schedule



## DeleteScheduledSingleSend

> SinglesendSchedule DeleteScheduledSingleSend(ctx, Id)

Delete Single Send Schedule

**This endpoint allows you to cancel a scheduled Single Send using a Single Send ID.**  Making a DELETE request to this endpoint will cancel the scheduled sending of a Single Send. The request will not delete the Single Send itself. Deleting a Single Send can be done by passing a DELETE request to `/marketing/singlesends/{id}`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteScheduledSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SinglesendSchedule**](SinglesendSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

