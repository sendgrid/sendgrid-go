# DownloadCsv

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadCsv**](DownloadCsv.md#DownloadCsv) | **Get** /v3/messages/download/{DownloadUuid} | Download CSV



## DownloadCsv

> DownloadCsv200Response DownloadCsv(ctx, DownloadUuid)

Download CSV

**This endpoint will return a presigned URL that can be used to download the CSV that was requested from the \"Request a CSV\" endpoint.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DownloadUuid** | **string** | UUID used to locate the download csv request entry in the DB.  This is the UUID provided in the email sent to the user when their csv file is ready to download

### Other Parameters

Other parameters are passed through a pointer to a DownloadCsvParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DownloadCsv200Response**](DownloadCsv200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

