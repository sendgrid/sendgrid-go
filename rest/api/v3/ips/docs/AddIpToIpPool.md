# AddIpToIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIpToIpPool**](AddIpToIpPool.md#AddIpToIpPool) | **Post** /v3/ips/pools/{PoolName}/ips | Add an IP address to a pool



## AddIpToIpPool

> AddIpToIpPool201Response AddIpToIpPool(ctx, PoolNameoptional)

Add an IP address to a pool

**This endpoint allows you to add an IP address to an IP pool.**  You can add the same IP address to multiple pools. It may take up to 60 seconds for your IP address to be added to a pool after your request is made.  Before you can add an IP to a pool, you need to activate it in your SendGrid account:  1. Log into your SendGrid account.   1. Navigate to **Settings** and then select **IP Addresses**.   1. Find the IP address you want to activate and then click **Edit**.   1. Check **Allow my account to send mail using this IP address**. 1. Click **Save**.  You can retrieve all of your available IP addresses from the \"Retrieve all IP addresses\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PoolName** | **string** | The name of the IP pool you want to add the address to. If the name contains spaces, they must be URL encoded (e.g., \"Test Pool\" becomes \"Test%20Pool\").

### Other Parameters

Other parameters are passed through a pointer to a AddIpToIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddIpToIpPoolRequest** | [**AddIpToIpPoolRequest**](AddIpToIpPoolRequest.md) | 

### Return type

[**AddIpToIpPool201Response**](AddIpToIpPool201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

