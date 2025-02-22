# DeleteReverseDns

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReverseDns**](DeleteReverseDns.md#DeleteReverseDns) | **Delete** /v3/whitelabel/ips/{Id} | Delete a reverse DNS record



## DeleteReverseDns

> DeleteReverseDns(ctx, Idoptional)

Delete a reverse DNS record

**This endpoint allows you to delete a reverse DNS record.**  A call to this endpoint will respond with a 204 status code if the deletion was successful.  You can retrieve the IDs associated with all your reverse DNS records using the \"Retrieve all reverse DNS records\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the reverse DNS record that you would like to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a DeleteReverseDnsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

