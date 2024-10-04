# DeleteSubUsersFromIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubUsersFromIp**](DeleteSubUsersFromIp.md#DeleteSubUsersFromIp) | **Post** /v3/send_ips/ips/{Ip}/subusers:batchDelete | Delete a Batch of Subusers from an IP



## DeleteSubUsersFromIp

> DeleteSubUsersFromIp(ctx, Ipoptional)

Delete a Batch of Subusers from an IP

This operation removes a batch of Subusers from a specified IP address.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Ip** | **string** | The `ip` path parameter specifies an IP address to make the request against.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubUsersFromIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteSubUsersFromIpRequest** | [**DeleteSubUsersFromIpRequest**](DeleteSubUsersFromIpRequest.md) | 

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

