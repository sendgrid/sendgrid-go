# ListIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIpPool**](ListIpPool.md#ListIpPool) | **Get** /v3/send_ips/pools | GET all IP Pools that have Associated IPs



## ListIpPool

> ListIpPool200Response ListIpPool(ctx, optional)

GET all IP Pools that have Associated IPs

This operation returns a list of your IP Pools and a sample of each Pools' associated IP addresses.  A maximum of 10 IPs will be returned per IP Pool by default. To retrieve additional IP addresses associated with a Pool, use the \"Get IPs Assigned to an IP Pool\" operation. Each user may have a maximum of 100 IP Pools.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int32** | Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
**AfterKey** | **int32** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
**Ip** | **string** | Specifies an IP address. The `ip` query parameter can be used to filter results by IP address.
**Region** | [**Region7**](Region7Region7.md) | Allowed values are `all`, `eu`, and `us`. If you provide a specific region, results will include all pools that have at least one IP in the filtered region. If `all`, pools with at least one IP (regardless of region) will be returned. If the `region` filter is not provided, the query returns all pools, including empty ones.
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP Pool's region information in the response.

### Return type

[**ListIpPool200Response**](ListIpPool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

