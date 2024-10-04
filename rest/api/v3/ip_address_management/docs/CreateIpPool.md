# CreateIpPool

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpPool**](CreateIpPool.md#CreateIpPool) | **Post** /v3/send_ips/pools | Create an IP Pool with a Name and IP Assignments



## CreateIpPool

> CreateIpPool201Response CreateIpPool(ctx, optional)

Create an IP Pool with a Name and IP Assignments

This operation will create a named IP Pool and associate specified IP addresses with the newly created Pool. This operation requires all IP assignments to succeed. If any IP assignments fail, this endpoint will return an error and the Pool will not be created.  Each IP Pool may have a maximum of 100 assigned IP addresses.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpPoolParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateIpPoolRequest** | [**CreateIpPoolRequest**](CreateIpPoolRequest.md) | 

### Return type

[**CreateIpPool201Response**](CreateIpPool201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

