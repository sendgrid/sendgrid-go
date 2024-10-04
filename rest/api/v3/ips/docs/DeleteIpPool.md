# DeleteIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpPool**](DeleteIpPool.md#DeleteIpPool) | **Delete** /v3/ips/pools/{PoolName} | Delete an IP pool



## DeleteIpPool

> map[string]interface{} DeleteIpPool(ctx, PoolName)

Delete an IP pool

**This endpoint allows you to delete an IP pool.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PoolName** | **string** | The name of the IP pool that you want to retrieve the IP addresses for.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpPoolParams struct


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

