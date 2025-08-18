# GetIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpPool**](GetIpPool.md#GetIpPool) | **Get** /v3/send_ips/pools/{Poolid} | Get Details for an IP Pool



## GetIpPool

> GetIpPool200Response GetIpPool(ctx, optional)

Get Details for an IP Pool

This operation will return the details for a specified IP Pool, including the Pool's name, ID, a sample list of the IPs associated with the Pool, and the total number of IPs belonging to the Pool.  A maximum of 10 IPs will be returned per IP Pool by default. To retrieve additional IP addresses associated with a Pool, use the \"Get IPs Assigned to an IP Pool\" operation.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a GetIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP Pool's region information in the response.

### Return type

[**GetIpPool200Response**](GetIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

