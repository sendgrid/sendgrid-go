# GetIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIp**](GetIp.md#GetIp) | **Get** /v3/ips/{IpAddress} | Retrieve all IP pools an IP address belongs to



## GetIp

> GetIp200Response GetIp(ctx, IpAddress)

Retrieve all IP pools an IP address belongs to

**This endpoint allows you to see which IP pools a particular IP address has been added to.**  The same IP address can be added to multiple IP pools.  A single IP address or a range of IP addresses may be dedicated to an account in order to send email for multiple domains. The reputation of this IP is based on the aggregate performance of all the senders who use it.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAddress** | **string** | The IP address you are retrieving the IP pools for.

### Other Parameters

Other parameters are passed through a pointer to a GetIpParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**GetIp200Response**](GetIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

