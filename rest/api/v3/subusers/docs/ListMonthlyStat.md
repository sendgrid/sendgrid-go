# ListMonthlyStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMonthlyStat**](ListMonthlyStat.md#ListMonthlyStat) | **Get** /v3/subusers/stats/monthly | Retrieve monthly stats for all subusers



## ListMonthlyStat

> SubuserStats ListMonthlyStat(ctx, Dateoptional)

Retrieve monthly stats for all subusers

**This endpoint allows you to retrieve the monthly email statistics for all subusers over the given date range.**  When using the `sort_by_metric` to sort your stats by a specific metric, you can not sort by the following metrics: `bounce_drops`, `deferred`, `invalid_emails`, `processed`, `spam_report_drops`, `spam_reports`, or `unsubscribe_drops`.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMonthlyStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Subuser** | **string** | A substring search of your subusers.
**SortByMetric** | [**SortByMetric**](SortByMetricSortByMetric.md) | The metric that you want to sort by. Metrics that you can sort by are: `blocks`, `bounces`, `clicks`, `delivered`, `opens`, `requests`, `unique_clicks`, `unique_opens`, and `unsubscribes`.'
**SortByDirection** | [**SortByDirection1**](SortByDirection1SortByDirection1.md) | The direction you want to sort.
**Limit** | **int32** | Optional field to limit the number of results returned.
**Offset** | **int32** | Optional beginning point in the list to retrieve from.

### Return type

[**SubuserStats**](SubuserStats.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

