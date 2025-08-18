# ListIpAssignedToIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIpAssignedToIpPool**](ListIpAssignedToIpPool.md#ListIpAssignedToIpPool) | **Get** /v3/send_ips/pools/{Poolid}/ips | Get IPs Assigned to an IP Pool



## ListIpAssignedToIpPool

> ListIpAssignedToIpPool200Response ListIpAssignedToIpPool(ctx, optional)

Get IPs Assigned to an IP Pool

This operation returns the IP addresses that are assigned to the specified IP pool.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpAssignedToIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int32** | Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
**AfterKey** | **int32** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP Pool's region information in the response.

### Return type

[**ListIpAssignedToIpPool200Response**](ListIpAssignedToIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

