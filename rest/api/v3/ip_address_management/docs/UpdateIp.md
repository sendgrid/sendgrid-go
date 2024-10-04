# UpdateIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIp**](UpdateIp.md#UpdateIp) | **Patch** /v3/send_ips/ips/{Ip} | Update Details for an IP Address



## UpdateIp

> UpdateIp200Response UpdateIp(ctx, Ipoptional)

Update Details for an IP Address

This operation updates an IP address's settings, including whether the IP is set to warm up automatically, if the IP is  assigned by a parent account, and whether the IP is enabled or disabled. The request body must include at least one of the `is_auto_warmup`, `is_parent_assigned`, or `is_enabled` fields.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Ip** | **string** | The `ip` path parameter specifies an IP address to make the request against.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateIpRequest** | [**UpdateIpRequest**](UpdateIpRequest.md) | 

### Return type

[**UpdateIp200Response**](UpdateIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

