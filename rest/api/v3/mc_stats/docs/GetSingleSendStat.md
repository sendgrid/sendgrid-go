# GetSingleSendStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSingleSendStat**](GetSingleSendStat.md#GetSingleSendStat) | **Get** /v3/marketing/stats/singlesends/{Id} | Get Single Send Stats by ID



## GetSingleSendStat

> SinglesendsResponse GetSingleSendStat(ctx, Idoptional)

Get Single Send Stats by ID

**This endpoint allows you to retrieve stats for an individual Single Send using a Single Send ID.**  Multiple Single Send IDs can be retrieved using the \"Get All Single Sends Stats\" endpoint. Once you have an ID, this endpoint will return detailed stats for the Single Send specified.  You may constrain the stats returned using the `start_date` and `end_date` query string parameters. You can also use the `group_by` and `aggregated_by` query string parameters to further refine the stats returned.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of Single Send for which you want to retrieve stats.

### Other Parameters

Other parameters are passed through a pointer to a GetSingleSendStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**AggregatedBy** | [**AggregatedBy**](AggregatedByAggregatedBy.md) | Dictates how the stats are time-sliced. Currently, `\"total\"` and `\"day\"` are supported.
**StartDate** | **string** | Format: `YYYY-MM-DD`. If this parameter is included, the stats' start date is included in the search.
**EndDate** | **string** | Format: `YYYY-MM-DD`.If this parameter is included, the stats' end date is included in the search.
**Timezone** | **string** | [IANA Area/Region](https://en.wikipedia.org/wiki/Tz_database#Names_of_timezones) string representing the timezone in which the stats are to be presented, e.g., \"America/Chicago\".
**PageSize** | **int32** | The number of elements you want returned on each page.
**PageToken** | **string** | The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.
**GroupBy** | [**[]Items1**](Items1.md) | A/B Single Sends have multiple variation IDs and phase IDs. Including these additional fields allows further granularity of stats by these fields.

### Return type

[**SinglesendsResponse**](SinglesendsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

