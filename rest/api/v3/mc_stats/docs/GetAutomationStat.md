# GetAutomationStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationStat**](GetAutomationStat.md#GetAutomationStat) | **Get** /v3/marketing/stats/automations/{Id} | Get Automation Stats by ID



## GetAutomationStat

> AutomationsResponse GetAutomationStat(ctx, Idoptional)

Get Automation Stats by ID

**This endpoint allows you to retrieve stats for a single Automation using its ID.**  Multiple Automation IDs can be retrieved using the \"Get All Automation Stats\" endpoint. Once you have an ID, this endpoint will return detailed stats for the single automation specified.  You may constrain the stats returned using the `start_date` and `end_date` query string parameters. You can also use the `group_by` and `aggregated_by` query string parameters to further refine the stats returned.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Automation for which you want to retrieve statistics. 

### Other Parameters

Other parameters are passed through a pointer to a GetAutomationStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**GroupBy** | [**[]Items**](Items.md) | Automations can have multiple steps. Including `step_id` as a `group_by` metric allows further granularity of stats.
**StepIds** | **[]string** | Comma-separated list of `step_ids` that you want the link stats for.
**AggregatedBy** | [**AggregatedBy**](AggregatedByAggregatedBy.md) | Dictates how the stats are time-sliced. Currently, `\"total\"` and `\"day\"` are supported.
**StartDate** | **string** | Format: `YYYY-MM-DD`. If this parameter is included, the stats' start date is included in the search.
**EndDate** | **string** | Format: `YYYY-MM-DD`.If this parameter is included, the stats' end date is included in the search.
**Timezone** | **string** | [IANA Area/Region](https://en.wikipedia.org/wiki/Tz_database#Names_of_timezones) string representing the timezone in which the stats are to be presented, e.g., \"America/Chicago\".
**PageSize** | **int32** | The number of elements you want returned on each page.
**PageToken** | **string** | The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.

### Return type

[**AutomationsResponse**](AutomationsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

