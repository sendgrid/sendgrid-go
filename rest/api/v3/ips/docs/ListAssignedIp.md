# ListAssignedIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAssignedIp**](ListAssignedIp.md#ListAssignedIp) | **Get** /v3/ips/assigned | Retrieve all assigned IPs



## ListAssignedIp

> []ListAssignedIp200ResponseInner ListAssignedIp(ctx, )

Retrieve all assigned IPs

**This endpoint allows you to retrieve only assigned IP addresses.**  A single IP address or a range of IP addresses may be dedicated to an account in order to send email for multiple domains. The reputation of this IP is based on the aggregate performance of all the senders who use it.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAssignedIpParams struct


### Return type

[**[]ListAssignedIp200ResponseInner**](ListAssignedIp200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

