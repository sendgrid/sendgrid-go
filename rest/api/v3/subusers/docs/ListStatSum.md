# ListStatSum

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStatSum**](ListStatSum.md#ListStatSum) | **Get** /v3/subusers/stats/sums | Retrieve the totals for each email statistic metric for all subusers.



## ListStatSum

> CategoryStats ListStatSum(ctx, StartDateoptional)

Retrieve the totals for each email statistic metric for all subusers.

**This endpoint allows you to retrieve the total sums of each email statistic metric for all subusers over the given date range.**

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListStatSumParams struct


Name | Type | Description
------------- | ------------- | -------------
**SortByDirection** | [**SortByDirection2**](SortByDirection2SortByDirection2.md) | The direction you want to sort. 
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.
**Limit** | **int32** | Limits the number of results returned per page.
**Offset** | **int32** | The point in the list to begin retrieving results from.
**AggregatedBy** | **string** | How to group the statistics. Defaults to today. Must follow format YYYY-MM-DD.
**SortByMetric** | **string** | The metric that you want to sort by.  Must be a single metric.

### Return type

[**CategoryStats**](CategoryStats.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

