# AddSubUsersToIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSubUsersToIp**](AddSubUsersToIp.md#AddSubUsersToIp) | **Post** /v3/send_ips/ips/{Ip}/subusers:batchAdd | Assign a Batch of Subusers to an IP



## AddSubUsersToIp

> AddSubUsersToIp200Response AddSubUsersToIp(ctx, Ipoptional)

Assign a Batch of Subusers to an IP

This operation appends a batch of Subusers to a specified IP address. This endpoint requires all Subuser assignments to succeed. If a Subuser assignment fails, this endpoint will return an error.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Ip** | **string** | The `ip` path parameter specifies an IP address to make the request against.

### Other Parameters

Other parameters are passed through a pointer to a AddSubUsersToIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddSubUsersToIpRequest** | [**AddSubUsersToIpRequest**](AddSubUsersToIpRequest.md) | 

### Return type

[**AddSubUsersToIp200Response**](AddSubUsersToIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

