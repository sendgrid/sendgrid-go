# ListSubuserMonthlyStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubuserMonthlyStat**](ListSubuserMonthlyStat.md#ListSubuserMonthlyStat) | **Get** /v3/subusers/{SubuserName}/stats/monthly | Retrieve the monthly email statistics for a single subuser



## ListSubuserMonthlyStat

> SubuserStats ListSubuserMonthlyStat(ctx, DateSubuserNameoptional)

Retrieve the monthly email statistics for a single subuser

**This endpoint allows you to retrive the monthly email statistics for a specific subuser.**  When using the `sort_by_metric` to sort your stats by a specific metric, you can not sort by the following metrics: `bounce_drops`, `deferred`, `invalid_emails`, `processed`, `spam_report_drops`, `spam_reports`, or `unsubscribe_drops`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubuserName** | **string** | The username of the Subuser.

### Other Parameters

Other parameters are passed through a pointer to a ListSubuserMonthlyStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**SortByMetric** | **string** | The metric that you want to sort by. Metrics that you can sort by are: `blocks`, `bounces`, `clicks`, `delivered`, `opens`, `requests`, `unique_clicks`, `unique_opens`, and `unsubscribes`.'
**SortByDirection** | [**SortByDirection**](SortByDirectionSortByDirection.md) | The direction you want to sort.
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

