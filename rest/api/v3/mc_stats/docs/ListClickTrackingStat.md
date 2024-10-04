# ListClickTrackingStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClickTrackingStat**](ListClickTrackingStat.md#ListClickTrackingStat) | **Get** /v3/marketing/stats/automations/{Id}/links | Get Automation Click Tracking Stats by ID



## ListClickTrackingStat

> AutmoationsLinkStatsResponse ListClickTrackingStat(ctx, Idoptional)

Get Automation Click Tracking Stats by ID

**This endpoint lets you retrieve click-tracking stats for a single Automation**.  The stats returned list the URLs embedded in your Automation and the number of clicks each one received.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Automation you want to get click tracking stats for. 

### Other Parameters

Other parameters are passed through a pointer to a ListClickTrackingStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**GroupBy** | [**[]Items**](Items.md) | Automations can have multiple steps. Including `step_id` as a `group_by` metric allows further granularity of stats.
**StepIds** | **[]string** | Comma-separated list of `step_ids` that you want the link stats for.
**PageSize** | **int32** | The number of elements you want returned on each page.
**PageToken** | **string** | The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.

### Return type

[**AutmoationsLinkStatsResponse**](AutmoationsLinkStatsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

