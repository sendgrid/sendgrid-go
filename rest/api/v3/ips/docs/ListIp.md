# ListIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIp**](ListIp.md#ListIp) | **Get** /v3/ips | Retrieve all IP addresses



## ListIp

> []ListIp200ResponseInner ListIp(ctx, optional)

Retrieve all IP addresses

**This endpoint allows you to retrieve a paginated list of all assigned and unassigned IPs.**  Response includes warm up status, pools, assigned subusers, and reverse DNS info. The start_date field corresponds to when warmup started for that IP.  A single IP address or a range of IP addresses may be dedicated to an account in order to send email for multiple domains. The reputation of this IP is determined by the aggregate performance of all email traffic sent from it.  You can use the `limit` query parameter to set the page size. If your list contains more items than the page size permits, you can make multiple requests. Use the `offset` query parameter to control the position in the list from which to start retrieving additional items.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ip** | **string** | The IP address to get
**ExcludeWhitelabels** | **bool** | Should we exclude reverse DNS records (whitelabels)?
**Limit** | **int32** | `limit` sets the page size, i.e. maximum number of items from the list to be returned for a single API request. If omitted, the default page size is used.
**Offset** | **int32** | The number of items in the list to skip over before starting to retrieve the items for the requested page. The default `offset` of `0` represents the beginning of the list, i.e. the start of the first page. To request the second page of the list, set the `offset` to the page size as determined by `limit`. Use multiples of the page size as your `offset` to request further consecutive pages. E.g. assume your page size is set to `10`. An `offset` of `10` requests the second page, an `offset` of `20` requests the third page and so on, provided there are sufficiently many items in your list.
**Subuser** | **string** | The subuser you are requesting for.
**SortByDirection** | [**SortByDirection**](SortByDirectionSortByDirection.md) | The direction to sort the results.

### Return type

[**[]ListIp200ResponseInner**](ListIp200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

