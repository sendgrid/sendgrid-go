# UpdateSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSingleSend**](UpdateSingleSend.md#UpdateSingleSend) | **Patch** /v3/marketing/singlesends/{Id} | Update Single Send



## UpdateSingleSend

> SinglesendResponse UpdateSingleSend(ctx, Idoptional)

Update Single Send

**This endpoint allows you to update a Single Send using a Single Send ID.**  You only need to pass the properties you want to update. Any blank or missing properties will remain unaltered.  This endpoint will update a draft of the Single Send but will not send it or schedule it to be sent. Any `send_at` property value set with this endpoint will prepopulate the Single Send's send date. However, the Single Send will remain an unscheduled draft until it's updated with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------
**SinglesendRequest** | [**SinglesendRequest**](SinglesendRequest.md) | 

### Return type

[**SinglesendResponse**](SinglesendResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

