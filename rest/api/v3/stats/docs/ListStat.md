# ListStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStat**](ListStat.md#ListStat) | **Get** /v3/stats | Retrieve global email statistics



## ListStat

> []ListStat200ResponseInner ListStat(ctx, StartDateoptional)

Retrieve global email statistics

**This endpoint allows you to retrieve all of your global email statistics between a given date range.**  Parent accounts can see either aggregated stats for the parent account or aggregated stats for a subuser specified in the `on-behalf-of` header. Subuser accounts will see only their own stats.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Limit** | **int32** | The number of results to return.
**Offset** | **int32** | The point in the list to begin retrieving results.
**AggregatedBy** | [**AggregatedBy3**](AggregatedBy3AggregatedBy3.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.

### Return type

[**[]ListStat200ResponseInner**](ListStat200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

