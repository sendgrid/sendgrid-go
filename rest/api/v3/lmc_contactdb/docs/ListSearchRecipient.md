# ListSearchRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSearchRecipient**](ListSearchRecipient.md#ListSearchRecipient) | **Get** /v3/contactdb/recipients/search | Search recipients



## ListSearchRecipient

> ListRecipientsFromContactDbList200Response ListSearchRecipient(ctx, optional)

Search recipients

**This endpoint allows you to perform a search on all of your Marketing Campaigns recipients.**  field_name:  * is a variable that is substituted for your actual custom field name from your recipient. * Text fields must be url-encoded. Date fields are searchable only by unix timestamp (e.g. 2/2/2015 becomes 1422835200) * If field_name is a 'reserved' date field, such as created_at or updated_at, the system will internally convert your epoch time to a date range encompassing the entire day. For example, an epoch time of 1422835600 converts to Mon, 02 Feb 2015 00:06:40 GMT, but internally the system will search from Mon, 02 Feb 2015 00:00:00 GMT through Mon, 02 Feb 2015 23:59:59 GMT.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSearchRecipientParams struct


Name | Type | Description
------------- | ------------- | -------------
**FieldName** | **string** | 
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ListRecipientsFromContactDbList200Response**](ListRecipientsFromContactDbList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

