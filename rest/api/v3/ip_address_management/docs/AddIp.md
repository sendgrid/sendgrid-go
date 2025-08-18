# AddIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIp**](AddIp.md#AddIp) | **Post** /v3/send_ips/ips | Add a Twilio SendGrid IP Address



## AddIp

> AddIp201Response AddIp(ctx, optional)

Add a Twilio SendGrid IP Address

This operation adds a Twilio SendGrid IP address to your account. You can also assign up to 100 Subusers to the IP address at creation.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a AddIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddIpRequest** | [**AddIpRequest**](AddIpRequest.md) | 

### Return type

[**AddIp201Response**](AddIp201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

