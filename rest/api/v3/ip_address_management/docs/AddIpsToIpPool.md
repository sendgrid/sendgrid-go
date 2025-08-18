# AddIpsToIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIpsToIpPool**](AddIpsToIpPool.md#AddIpsToIpPool) | **Post** /v3/send_ips/pools/{Poolid}/ips:batchAdd | Add a Batch of IPs to an IP Pool



## AddIpsToIpPool

> AddIpsToIpPool200Response AddIpsToIpPool(ctx, Poolidoptional)

Add a Batch of IPs to an IP Pool

This operation appends a batch of IPs to an IP Pool. This operation requires all IP assignments to succeed. If any IP assignments fail, this endpoint will return an error.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Poolid** | **string** | Specifies the unique ID for an IP Pool.

### Other Parameters

Other parameters are passed through a pointer to a AddIpsToIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddIpsToIpPoolRequest** | [**AddIpsToIpPoolRequest**](AddIpsToIpPoolRequest.md) | 

### Return type

[**AddIpsToIpPool200Response**](AddIpsToIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

