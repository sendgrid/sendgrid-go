# ExportSingleSendStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSingleSendStat**](ExportSingleSendStat.md#ExportSingleSendStat) | **Get** /v3/marketing/stats/singlesends/export | Export Single Send Stats



## ExportSingleSendStat

> string ExportSingleSendStat(ctx, optional)

Export Single Send Stats

**This endpoint allows you to export Single Send stats as .CSV data**.  You can specify one Single Send or many: include as many Single Send IDs as you need, separating them with commas, as the value of the `ids` query string parameter.  The data is returned as plain text response but in .CSV format, so your application making the call can present the information in whatever way is most appropriate, or just save the data as a .csv file.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ExportSingleSendStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ids** | **[]string** | The IDs of Single Sends for which to export stats.
**Timezone** | **string** | The [IANA Area/Region](https://en.wikipedia.org/wiki/Tz_database#Names_of_timezones) string representing the timezone in which the stats are to be presented; i.e. `\"America/Chicago\"`. This parameter changes the timezone format only; it does not alter which stats are returned.

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

