# UpdateSubuserIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubuserIp**](UpdateSubuserIp.md#UpdateSubuserIp) | **Put** /v3/subusers/{SubuserName}/ips | Update IPs assigned to a subuser



## UpdateSubuserIp

> UpdateSubuserIp200Response UpdateSubuserIp(ctx, SubuserNameoptional)

Update IPs assigned to a subuser

**This endpoint allows you update your subusers' assigned IP.**  Each subuser should be assigned to an IP address from which all of this subuser's mail will be sent. Often, this is the same IP as the parent account, but each subuser can have one or more of their own IP addresses as well.   More information:  * [How to request more IPs](https://sendgrid.com/docs/ui/account-and-settings/dedicated-ip-addresses/) * [Setup Reverse DNS](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-reverse-dns/)

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubuserIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**RequestBody** | **[]string** | 

### Return type

[**UpdateSubuserIp200Response**](UpdateSubuserIp200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

