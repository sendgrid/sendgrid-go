# UpdateIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIpPool**](UpdateIpPool.md#UpdateIpPool) | **Put** /v3/send_ips/pools/{Poolid} | Update an IP Pool Name



## UpdateIpPool

> UpdateIpPool200Response UpdateIpPool(ctx, Poolidoptional)

Update an IP Pool Name

This operation will rename an IP Pool. An IP Pool name cannot start with a dot/period (.) or space.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Poolid** | **string** | Specifies the unique ID for an IP Pool.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateIpPoolRequest** | [**UpdateIpPoolRequest**](UpdateIpPoolRequest.md) | 

### Return type

[**UpdateIpPool200Response**](UpdateIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

