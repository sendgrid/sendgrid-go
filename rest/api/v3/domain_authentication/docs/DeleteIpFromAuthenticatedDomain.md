# DeleteIpFromAuthenticatedDomain

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpFromAuthenticatedDomain**](DeleteIpFromAuthenticatedDomain.md#DeleteIpFromAuthenticatedDomain) | **Delete** /v3/whitelabel/domains/{Id}/ips/{Ip} | Remove an IP from an authenticated domain.



## DeleteIpFromAuthenticatedDomain

> AuthenticatedDomainSpf DeleteIpFromAuthenticatedDomain(ctx, IdIpoptional)

Remove an IP from an authenticated domain.

**This endpoint allows you to remove an IP address from that domain's authentication.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **int32** | ID of the domain to delete the IP from.
**Ip** | **string** | IP to remove from the domain.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpFromAuthenticatedDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**AuthenticatedDomainSpf**](AuthenticatedDomainSpf.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

