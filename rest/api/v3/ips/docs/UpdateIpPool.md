# UpdateIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIpPool**](UpdateIpPool.md#UpdateIpPool) | **Put** /v3/ips/pools/{PoolName} | Rename an IP pool



## UpdateIpPool

> IpPools200 UpdateIpPool(ctx, PoolNameoptional)

Rename an IP pool

**This endpoint allows you to update the name of an IP pool.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PoolName** | **string** | The name of the IP pool that you want to retrieve the IP addresses for.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateIpPoolRequest** | [**UpdateIpPoolRequest**](UpdateIpPoolRequest.md) | 

### Return type

[**IpPools200**](IpPools200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

