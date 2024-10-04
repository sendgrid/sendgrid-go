# DeleteIpFromIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpFromIpPool**](DeleteIpFromIpPool.md#DeleteIpFromIpPool) | **Delete** /v3/ips/pools/{PoolName}/ips/{Ip} | Remove an IP address from a pool



## DeleteIpFromIpPool

> map[string]interface{} DeleteIpFromIpPool(ctx, PoolNameIp)

Remove an IP address from a pool

**This endpoint allows you to remove an IP address from an IP pool.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PoolName** | **string** | The name of the IP pool that you are removing the IP address from.
**Ip** | **string** | The IP address that you wish to remove.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpFromIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

