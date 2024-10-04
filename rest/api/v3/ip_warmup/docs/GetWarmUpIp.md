# GetWarmUpIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWarmUpIp**](GetWarmUpIp.md#GetWarmUpIp) | **Get** /v3/ips/warmup/{IpAddress} | Retrieve the warmup status for a specific IP address



## GetWarmUpIp

> []IpWarmup200Inner GetWarmUpIp(ctx, IpAddress)

Retrieve the warmup status for a specific IP address

**This endpoint allows you to retrieve the warmup status for a specific IP address.**  You can retrieve all of your warming IPs using the \"Retrieve all IPs currently in warmup\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAddress** | **string** | The IP address that you want to retrieve the warmup status for.

### Other Parameters

Other parameters are passed through a pointer to a GetWarmUpIpParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**[]IpWarmup200Inner**](IpWarmup200Inner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

