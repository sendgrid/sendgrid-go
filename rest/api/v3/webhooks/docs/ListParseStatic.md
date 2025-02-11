# ListParseStatic

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListParseStatic**](ListParseStatic.md#ListParseStatic) | **Get** /v3/user/webhooks/parse/stats | Retrieves Inbound Parse Webhook statistics.



## ListParseStatic

> []ListParseStatic200ResponseInner ListParseStatic(ctx, StartDateoptional)

Retrieves Inbound Parse Webhook statistics.

**This endpoint allows you to retrieve the statistics for your Parse Webhook usage.**  SendGrid's Inbound Parse Webhook allows you to parse the contents and attachments of incoming emails. The Parse API can then POST the parsed emails to a URL that you specify. The Inbound Parse Webhook cannot parse messages greater than 30MB in size, including all attachments.  There are a number of pre-made integrations for the SendGrid Parse Webhook which make processing events easy. You can find these integrations in the [Library Index](https://docs.sendgrid.com/for-developers/sending-email/libraries#webhook-libraries).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListParseStaticParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **string** | The number of statistics to return on each page.
**Offset** | **string** | The number of statistics to skip.
**AggregatedBy** | [**AggregatedBy**](AggregatedByAggregatedBy.md) | How you would like the statistics to by grouped. 
**EndDate** | **string** | The end date of the statistics you want to retrieve. Must be in the format YYYY-MM-DD
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**[]ListParseStatic200ResponseInner**](ListParseStatic200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

