# GetExportContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExportContact**](GetExportContact.md#GetExportContact) | **Get** /v3/marketing/contacts/exports/{Id} | Export Contacts Status



## GetExportContact

> ContactExport GetExportContact(ctx, Id)

Export Contacts Status

**This endpoint can be used to check the status of a contact export job**.   To use this call, you will need the `id` from the \"Export Contacts\" call.  If you would like to download a list, take the `id` that is returned from the \"Export Contacts\" endpoint and make an API request here to get the `urls`. Once you have the list of URLs, make a `GET` request on each URL to download your CSV file(s).  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetExportContactParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContactExport**](ContactExport.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

