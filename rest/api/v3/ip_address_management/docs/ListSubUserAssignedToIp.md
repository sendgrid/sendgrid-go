# ListSubUserAssignedToIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubUserAssignedToIp**](ListSubUserAssignedToIp.md#ListSubUserAssignedToIp) | **Get** /v3/send_ips/ips/{Ip}/subusers | Get a List of Subusers Assigned to an IP



## ListSubUserAssignedToIp

> ListSubUserAssignedToIp200Response ListSubUserAssignedToIp(ctx, optional)

Get a List of Subusers Assigned to an IP

This operation returns a list of Subuser IDs that have been assigned the specified IP address. To retrieve more information about the returned Subusers, use the [Subusers API](https://docs.sendgrid.com/api-reference/subusers-api/list-all-subusers).  You can use the `after_key` and `limit` query parameters to iterate through paginated results. The maximum limit is 100, meaning you may retrieve up to 100 Subusers per request. If the `after_key` in the API response is not null, there are more Subusers assigned to the IP address than those returned in the request. You can repeat the request with the non-null `after_key` value and the same limit to retrieve the next group of Subusers.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubUserAssignedToIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**AfterKey** | **int32** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
**Limit** | **int32** | Specifies the number of results to be returned by the API. This parameter can be used in combination with the `after_key` parameters to iterate through paginated results. The maximum limit is 100.

### Return type

[**ListSubUserAssignedToIp200Response**](ListSubUserAssignedToIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

