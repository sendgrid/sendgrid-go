# ListSingleSendTrackingStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSingleSendTrackingStat**](ListSingleSendTrackingStat.md#ListSingleSendTrackingStat) | **Get** /v3/marketing/stats/singlesends/{Id}/links | Get Single Send Click Tracking Stats by ID



## ListSingleSendTrackingStat

> SinglesendsLinkStatsResponse ListSingleSendTrackingStat(ctx, Idoptional)

Get Single Send Click Tracking Stats by ID

**This endpoint lets you retrieve click-tracking stats for one Single Send**.  The stats returned list the URLs embedded in the specified Single Send and the number of clicks each one received.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of Single Send for which you want to retrieve link stats.

### Other Parameters

Other parameters are passed through a pointer to a ListSingleSendTrackingStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | The number of elements you want returned on each page.
**PageToken** | **string** | The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.
**GroupBy** | [**[]Items2**](Items2.md) | A/B Single Sends have multiple variation IDs and phase IDs. Including these additional fields allows further granularity of stats by these fields.
**AbVariationId** | **string** | 
**AbPhaseId** | [**AbPhaseId**](AbPhaseIdAbPhaseId.md) | 

### Return type

[**SinglesendsLinkStatsResponse**](SinglesendsLinkStatsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

