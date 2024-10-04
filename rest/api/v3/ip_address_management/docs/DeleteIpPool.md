# DeleteIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpPool**](DeleteIpPool.md#DeleteIpPool) | **Delete** /v3/send_ips/pools/{Poolid} | Delete IP Pool



## DeleteIpPool

> DeleteIpPool(ctx, Poolid)

Delete IP Pool

This operation deletes an IP Pool and unassigns all IP addresses associated with the Pool. IP addresses associated with the deleted Pool will remain in your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Poolid** | **string** | Specifies the unique ID for an IP Pool.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

