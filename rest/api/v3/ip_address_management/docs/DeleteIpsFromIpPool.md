# DeleteIpsFromIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpsFromIpPool**](DeleteIpsFromIpPool.md#DeleteIpsFromIpPool) | **Post** /v3/send_ips/pools/{Poolid}/ips:batchDelete | Delete a Batch of IPs from an IP Pool



## DeleteIpsFromIpPool

> DeleteIpsFromIpPool(ctx, Poolidoptional)

Delete a Batch of IPs from an IP Pool

This operation removes a batch of IPs from an IP Pool. All IPs associated with the Pool will be unassigned from the deleted Pool. However, this operation does not remove the IPs from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Poolid** | **string** | Specifies the unique ID for an IP Pool.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpsFromIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteIpsFromIpPoolRequest** | [**DeleteIpsFromIpPoolRequest**](DeleteIpsFromIpPoolRequest.md) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

