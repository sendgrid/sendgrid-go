# GetClientStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientStat**](GetClientStat.md#GetClientStat) | **Get** /v3/clients/{ClientType}/stats | Retrieve stats by a specific client type.



## GetClientStat

> []ListClientStat200ResponseInner GetClientStat(ctx, StartDateClientTypeoptional)

Retrieve stats by a specific client type.

**This endpoint allows you to retrieve your email statistics segmented by a specific client type.**  **We only store up to 7 days of email activity in our database.** By default, 500 items will be returned per request via the Advanced Stats API endpoints.  ### Available Client Types - phone - tablet - webmail - desktop  Advanced Stats provide a more in-depth view of your email statistics and the actions taken by your recipients. You can segment these statistics by geographic location, device type, client type, browser, and mailbox provider. For more information about statistics, please see our [Statistics Overview](https://sendgrid.com/docs/ui/analytics-and-reporting/stats-overview/).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ClientType** | [**ClientType**](ClientType.md) | Specifies the type of client to retrieve stats for. Must be either \"phone\", \"tablet\", \"webmail\", or \"desktop\".

### Other Parameters

Other parameters are passed through a pointer to a GetClientStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.
**AggregatedBy** | [**AggregatedBy2**](AggregatedBy2AggregatedBy2.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".

### Return type

[**[]ListClientStat200ResponseInner**](ListClientStat200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
