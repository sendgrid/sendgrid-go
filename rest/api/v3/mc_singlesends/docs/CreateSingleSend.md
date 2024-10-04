# CreateSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSingleSend**](CreateSingleSend.md#CreateSingleSend) | **Post** /v3/marketing/singlesends | Create Single Send



## CreateSingleSend

> SinglesendResponse CreateSingleSend(ctx, optional)

Create Single Send

**This endpoint allows you to create a new Single Send.**  Please note that if you are migrating from the previous version of Single Sends, you no longer need to pass a template ID with your request to this endpoint. Instead, you will pass all template data in the `email_config` object.  This endpoint will create a draft of the Single Send but will not send it or schedule it to be sent. Any `send_at` property value set with this endpoint will prepopulate the Single Send's send date. However, the Single Send will remain an unscheduled draft until it's updated with the [**Schedule Single Send**](https://docs.sendgrid.com/api-reference/single-sends/schedule-single-send) endpoint or SendGrid application UI.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSingleSendParams struct


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

