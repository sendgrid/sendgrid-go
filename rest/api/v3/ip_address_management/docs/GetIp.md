# GetIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIp**](GetIp.md#GetIp) | **Get** /v3/send_ips/ips/{Ip} | Get Details for an IP Address



## GetIp

> GetIp200Response GetIp(ctx, optional)

Get Details for an IP Address

This operation returns details for a specified IP address. Details include whether the IP is assigned to a parent account, set to warm up automatically, which Pools the IP is associated with, when the IP was added and modified, whether the IP is leased, and whether the IP is enabled. Note that this operation will not return Subuser information associated with the IP. To retrieve Subuser information, use the \"Get a List of Subusers Assigned to an IP\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a GetIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP Pool's region information in the response.

### Return type

[**GetIp200Response**](GetIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

