# RequestCsv

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestCsv**](RequestCsv.md#RequestCsv) | **Post** /v3/messages/download | Request CSV



## RequestCsv

> RequestCsv202Response RequestCsv(ctx, optional)

Request CSV

This request will kick off a backend process to generate a CSV file. Once generated, the worker will then send an email for the user download the file. The link will expire in 3 days.  The CSV will contain the events from the last 30 days, limited to the last 1 million events maximum. This endpoint will be rate limited to 1 request every 12 hours (rate limit may change).  This endpoint is similar to the GET Single Message endpoint - the only difference is that /download is added to indicate that this is a CSV download requests but the same query is used to determine what the CSV should contain.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a RequestCsvParams struct


Name | Type | Description
------------- | ------------- | -------------
**Query** | **string** | Uses a SQL like syntax to indicate which messages to include in the CSV

### Return type

[**RequestCsv202Response**](RequestCsv202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

