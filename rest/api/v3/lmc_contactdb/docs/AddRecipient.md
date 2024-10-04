# AddRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecipient**](AddRecipient.md#AddRecipient) | **Post** /v3/contactdb/recipients | Add recipients



## AddRecipient

> ContactdbRecipientResponse201 AddRecipient(ctx, optional)

Add recipients

**This endpoint allows you to add a Marketing Campaigns recipient.**  You can add custom field data as a parameter on this endpoint. We have provided an example using some of the default custom fields SendGrid provides.  The rate limit is three requests every 2 seconds. You can upload 1000  contacts per request. So the maximum upload rate is 1500 recipients per second.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a AddRecipientParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**AddRecipientRequestInner** | [**[]AddRecipientRequestInner**](AddRecipientRequestInner.md) | 

### Return type

[**ContactdbRecipientResponse201**](ContactdbRecipientResponse201.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

