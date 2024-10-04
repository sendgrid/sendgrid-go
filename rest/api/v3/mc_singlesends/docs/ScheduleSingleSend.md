# ScheduleSingleSend

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduleSingleSend**](ScheduleSingleSend.md#ScheduleSingleSend) | **Put** /v3/marketing/singlesends/{Id}/schedule | Schedule Single Send



## ScheduleSingleSend

> ScheduleSingleSend201Response ScheduleSingleSend(ctx, Idoptional)

Schedule Single Send

**This endpoint allows you to send a Single Send immediately or schedule it to be sent at a later time.**  To send your message immediately, set the `send_at` property value to the string `now`. To schedule the Single Send for future delivery, set the `send_at` value to your desired send time in [ISO 8601 date time format](https://www.iso.org/iso-8601-date-and-time-format.html) (`yyyy-MM-ddTHH:mm:ssZ`).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ScheduleSingleSendParams struct


Name | Type | Description
------------- | ------------- | -------------
**ScheduleSingleSendRequest** | [**ScheduleSingleSendRequest**](ScheduleSingleSendRequest.md) | 

### Return type

[**ScheduleSingleSend201Response**](ScheduleSingleSend201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

