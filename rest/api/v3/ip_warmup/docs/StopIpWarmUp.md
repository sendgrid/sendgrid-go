# StopIpWarmUp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StopIpWarmUp**](StopIpWarmUp.md#StopIpWarmUp) | **Delete** /v3/ips/warmup/{IpAddress} | Stop warming up an IP address



## StopIpWarmUp

> StopIpWarmUp(ctx, IpAddress)

Stop warming up an IP address

**This endpoint allows you to remove an IP address from warmup mode.**  Your request will return a 204 status code if the specified IP was successfully removed from warmup mode. To retrieve details of the IP’s warmup status *before* removing it from warmup mode, call the  \"Retrieve the warmpup status for a specific IP address\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAddress** | **string** | The IP address that you want to retrieve the warmup status for.

### Other Parameters

Other parameters are passed through a pointer to a StopIpWarmUpParams struct


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

