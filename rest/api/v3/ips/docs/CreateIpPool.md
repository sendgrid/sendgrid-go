# CreateIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpPool**](CreateIpPool.md#CreateIpPool) | **Post** /v3/ips/pools | Create an IP pool



## CreateIpPool

> IpPools200 CreateIpPool(ctx, optional)

Create an IP pool

**This endpoint allows you to create an IP pool.**  Before you can create an IP pool, you need to activate the IP in your SendGrid account:   1. Log into your SendGrid account.   1. Navigate to **Settings** and then select **IP Addresses**.   1. Find the IP address you want to activate and then click **Edit**.   1. Check **Allow my account to send mail using this IP address**. 1. Click **Save**.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateIpPoolRequest** | [**CreateIpPoolRequest**](CreateIpPoolRequest.md) | 

### Return type

[**IpPools200**](IpPools200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

