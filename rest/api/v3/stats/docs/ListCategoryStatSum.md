# ListCategoryStatSum

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCategoryStatSum**](ListCategoryStatSum.md#ListCategoryStatSum) | **Get** /v3/categories/stats/sums | Retrieve sums of email stats for each category.



## ListCategoryStatSum

> CategoryStats ListCategoryStatSum(ctx, StartDateoptional)

Retrieve sums of email stats for each category.

**This endpoint allows you to retrieve the total sum of each email statistic for every category over the given date range.**  If you do not define any query parameters, this endpoint will return a sum for each category in groups of 10.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCategoryStatSumParams struct


Name | Type | Description
------------- | ------------- | -------------
**SortByMetric** | **string** | The metric that you want to sort by.  Must be a single metric.
**SortByDirection** | [**SortByDirection**](SortByDirectionSortByDirection.md) | The direction you want to sort.
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.
**Limit** | **int32** | Limits the number of results returned.
**Offset** | **int32** | The point in the list to begin retrieving results.
**AggregatedBy** | [**AggregatedBy1**](AggregatedBy1AggregatedBy1.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

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

