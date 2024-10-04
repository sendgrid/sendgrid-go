# GetIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpPool**](GetIpPool.md#GetIpPool) | **Get** /v3/ips/pools/{PoolName} | Retrieve all the IPs in a specified pool



## GetIpPool

> GetIpPool200Response GetIpPool(ctx, PoolName)

Retrieve all the IPs in a specified pool

**This endpoint allows you to get all of the IP addresses that are in a specific IP pool.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PoolName** | **string** | The name of the IP pool that you want to retrieve the IP addresses for.

### Other Parameters

Other parameters are passed through a pointer to a GetIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**GetIpPool200Response**](GetIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

