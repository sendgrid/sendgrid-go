# ListBrowserStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBrowserStat**](ListBrowserStat.md#ListBrowserStat) | **Get** /v3/browsers/stats | Retrieve email statistics by browser.



## ListBrowserStat

> []ListBrowserStat200ResponseInner ListBrowserStat(ctx, StartDateoptional)

Retrieve email statistics by browser.

**This endpoint allows you to retrieve your email statistics segmented by browser type.**  **We only store up to 7 days of email activity in our database.** By default, 500 items will be returned per request via the Advanced Stats API endpoints.  Advanced Stats provide a more in-depth view of your email statistics and the actions taken by your recipients. You can segment these statistics by geographic location, device type, client type, browser, and mailbox provider. For more information about statistics, please see our [Statistics Overview](https://sendgrid.com/docs/ui/analytics-and-reporting/stats-overview/).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBrowserStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Browsers** | **string** | The browsers to get statistics for. You can include up to 10 different browsers by including this parameter multiple times.
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Limit** | **int32** | The number of results to return.
**Offset** | **int32** | The point in the list to begin retrieving results.
**AggregatedBy** | [**AggregatedBy3**](AggregatedBy3AggregatedBy3.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.

### Return type

[**[]ListBrowserStat200ResponseInner**](ListBrowserStat200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
