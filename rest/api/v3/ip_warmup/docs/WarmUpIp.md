# WarmUpIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WarmUpIp**](WarmUpIp.md#WarmUpIp) | **Post** /v3/ips/warmup | Start warming up an IP address



## WarmUpIp

> []IpWarmup200Inner WarmUpIp(ctx, optional)

Start warming up an IP address

**This endpoint allows you to put an IP address into warmup mode.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a WarmUpIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**WarmUpIpRequest** | [**WarmUpIpRequest**](WarmUpIpRequest.md) | 

### Return type

[**[]IpWarmup200Inner**](IpWarmup200Inner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

