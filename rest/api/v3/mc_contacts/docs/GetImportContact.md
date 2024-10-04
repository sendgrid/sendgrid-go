# GetImportContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportContact**](GetImportContact.md#GetImportContact) | **Get** /v3/marketing/contacts/imports/{Id} | Import Contacts Status



## GetImportContact

> ContactImport GetImportContact(ctx, Id)

Import Contacts Status

**This endpoint can be used to check the status of a contact import job**.   Use the `job_id` from the \"Import Contacts,\" \"Add or Update a Contact,\" or \"Delete Contacts\" endpoints as the `id` in the path parameter.  If there is an error with your `PUT` request, download the `errors_url` file and open it to view more details.  The job `status` field indicates whether the job is `pending`, `completed`, `errored`, or `failed`.   Pending means not started. Completed means finished without any errors. Errored means finished with some errors. Failed means finished with all errors, or the job was entirely unprocessable: for example, if you attempt to import file format we do not support.  The `results` object will have fields depending on the job type.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a GetImportContactParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContactImport**](ContactImport.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

