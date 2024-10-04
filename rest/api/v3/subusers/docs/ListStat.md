# ListStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStat**](ListStat.md#ListStat) | **Get** /v3/subusers/stats | Retrieve email statistics for your subusers.



## ListStat

> []CategoryStats ListStat(ctx, SubusersStartDateoptional)

Retrieve email statistics for your subusers.

**This endpoint allows you to retrieve the email statistics for the given subusers.**  You may retrieve statistics for up to 10 different subusers by including an additional _subusers_ parameter for each additional subuser.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int32** | Limits the number of results returned per page.
**Offset** | **int32** | The point in the list to begin retrieving results from.
**AggregatedBy** | [**AggregatedBy**](AggregatedByAggregatedBy.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today.

### Return type

[**[]CategoryStats**](CategoryStats.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

