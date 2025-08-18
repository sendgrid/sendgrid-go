# ListIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIp**](ListIp.md#ListIp) | **Get** /v3/send_ips/ips | Get a List of all IP Addresses on your Account



## ListIp

> ListIp200Response ListIp(ctx, optional)

Get a List of all IP Addresses on your Account

This operation returns a list of all IP addresses associated with your account. A sample of IP details is returned with each IP, including which Pools the IP is associated with, whether the IP is set to warm up automatically, and when the IP was last updated.  ### Limitations  The `is_parent_assigned` parameter and `pool` parameter cannot be used at the same time. By definition, an IP cannot be assigned to a Pool if it is not first enabled. You can use either the `before_key` or `after_key` in combination with the `limit` parameter to iterate through paginated results but not both.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ip** | **string** | Specifies an IP address. The `ip` query parameter can be used to filter results by IP address.
**Limit** | **int32** | Specifies the number of results to be returned by the API. This parameter can be used in combination with the `before_key` or `after_key` parameters to iterate through paginated results.
**AfterKey** | **int32** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
**BeforeKey** | **string** | Specifies which items to be returned by the API. When the `before_key` is specified, the API will return items beginning from the first item before the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.
**IsLeased** | **bool** | Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account.
**IsEnabled** | **bool** | Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses.
**IsParentAssigned** | **bool** | A parent must be assigned to an IP address before the parent can send mail from the IP and before the address can be assigned to an IP pool. Set this parameter value to true to allow the parent to send mail from the IP and make the IP eligible for IP pool assignment using the IP pool endpoints.
**Pool** | **string** | Specifies the unique ID for an IP Pool. When included, only IP addresses belonging to the specified Pool will be returned.
**StartAddedAt** | **int32** | The `start_added_at` and `end_added_at` parameters are used to set a time window. IP addresses that were added to your account in the specified time window will be returned. The `start_added_at` parameter sets the beginning of the time window.
**EndAddedAt** | **int32** | The `start_added_at` and `end_added_at` parameters are used to set a time window. IP addresses that were added to your account in the specified time window will be returned. The `end_added_at` parameter sets the end of the time window.
**Region** | [**Region7**](Region7Region7.md) | Allowed values are `all`, `eu`, and `us`. If you provide a specific region, results will include all pools that have at least one IP in the filtered region. If `all`, pools with at least one IP (regardless of region) will be returned. If the `region` filter is not provided, the query returns all pools, including empty ones.
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP Pool's region information in the response.

### Return type

[**ListIp200Response**](ListIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

