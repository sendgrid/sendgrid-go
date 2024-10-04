# GetAllowedIp

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllowedIp**](GetAllowedIp.md#GetAllowedIp) | **Get** /v3/access_settings/whitelist/{RuleId} | Retrieve a specific allowed IP



## GetAllowedIp

> IpAccessManagement2xx GetAllowedIp(ctx, RuleIdoptional)

Retrieve a specific allowed IP

**This endpoint allows you to retreive a specific IP address that has been allowed to access your account.**  You must include the ID for the specific IP address you want to retrieve in your call. You can retrieve the IDs associated with your allowed IP addresses using the \"Retrieve a  list of currently allowed IPs\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RuleId** | **string** | The ID of the allowed IP address that you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a GetAllowedIpParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**IpAccessManagement2xx**](IpAccessManagement2xx.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

