# ListRemainingIpCount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRemainingIpCount**](ListRemainingIpCount.md#ListRemainingIpCount) | **Get** /v3/ips/remaining | Get remaining IPs count



## ListRemainingIpCount

> ListRemainingIpCount200Response ListRemainingIpCount(ctx, )

Get remaining IPs count

**This endpoint gets amount of IP Addresses that can still be created during a given period and the price of those IPs.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRemainingIpCountParams struct


### Return type

[**ListRemainingIpCount200Response**](ListRemainingIpCount200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

