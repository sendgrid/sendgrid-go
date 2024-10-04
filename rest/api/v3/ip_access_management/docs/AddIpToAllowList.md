# AddIpToAllowList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIpToAllowList**](AddIpToAllowList.md#AddIpToAllowList) | **Post** /v3/access_settings/whitelist | Add one or more IPs to the allow list



## AddIpToAllowList

> IpAccessManagement2xx AddIpToAllowList(ctx, optional)

Add one or more IPs to the allow list

**This endpoint allows you to add one or more allowed IP addresses.**  To allow one or more IP addresses, pass them to this endpoint in an array. Once an IP address is added to your allow list, it will be assigned an `id` that can be used to remove the address. You can retrieve the ID associated with an IP using the \"Retrieve a list of currently allowed IPs\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a AddIpToAllowListParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**AddIpToAllowListRequest** | [**AddIpToAllowListRequest**](AddIpToAllowListRequest.md) | 

### Return type

[**IpAccessManagement2xx**](IpAccessManagement2xx.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

