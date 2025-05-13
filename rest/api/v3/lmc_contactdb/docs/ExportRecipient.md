# ExportRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportRecipient**](ExportRecipient.md#ExportRecipient) | **Post** /v3/contactdb/exports | Export Recipients



## ExportRecipient

> ExportRecipient202Response ExportRecipient(ctx, optional)

Export Recipients

**Use this endpoint to export lists or segments of recipients**.  If you would just like to have a link to the exported list sent to your email set the `notifications.email` option to `true` in the `POST` payload.  If you would like to download the list, take the `id` that is returned and use the \"Export Recipients Status\" endpoint to get the `urls`. Once you have the list of URLs, make a `GET` request to each URL provided to download your CSV file(s).  You specify the segments and or/recipient lists you wish to export by providing the relevant IDs in, respectively, the `segment_ids` and `list_ids` fields in the request body.  The lists will be provided in either JSON or CSV files. To specify which of these you would required, set the request body `file_type` field to `json` or `csv`.  You can also specify a maximum file size (in MB). If the export file is larger than this, it will be split into multiple files.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ExportRecipientParams struct


Name | Type | Description
------------- | ------------- | -------------
**ExportRecipientRequest** | [**ExportRecipientRequest**](ExportRecipientRequest.md) | 

### Return type

[**ExportRecipient202Response**](ExportRecipient202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

