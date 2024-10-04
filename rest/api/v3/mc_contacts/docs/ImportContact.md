# ImportContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportContact**](ImportContact.md#ImportContact) | **Put** /v3/marketing/contacts/imports | Import Contacts



## ImportContact

> ImportContact200Response ImportContact(ctx, optional)

Import Contacts

**This endpoint allows a CSV upload containing up to one million contacts or 5GB of data, whichever is smaller. At least one identifier is required for a successful import.**  Imports take place asynchronously: the endpoint returns a URL (`upload_uri`) and HTTP headers (`upload_headers`) which can subsequently be used to `PUT` a file of contacts to be imported into our system.  Uploaded CSV files may also be [gzip-compressed](https://en.wikipedia.org/wiki/Gzip).  In either case, you must include the field `file_type` with the value `csv` in your request body.  The `field_mappings` parameter is a respective list of field definition IDs to map the uploaded CSV columns to. It allows you to use CSVs where one or more columns are skipped (`null`) or remapped to the contact field.  For example, if `field_mappings` is set to `[null, \"w1\", \"_rf1\"]`, this means skip column 0, map column 1 to the custom field with the ID `w1`, and map column 2 to the reserved field with the ID `_rf1`. See the \"Get All Field Definitions\" endpoint to fetch your custom and reserved field IDs to use with `field_mappings`.  Once you receive the response body you can then initiate a **second** API call where you use the supplied URL and HTTP header to upload your file. For example:  `curl --upload-file \"file/path.csv\" \"URL_GIVEN\" -H \"HEADER_GIVEN\"`  If you would like to monitor the status of your import job, use the `job_id` and the \"Import Contacts Status\" endpoint.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ImportContactParams struct


Name | Type | Description
------------- | ------------- | -------------
**ImportContactRequest** | [**ImportContactRequest**](ImportContactRequest.md) | 

### Return type

[**ImportContact200Response**](ImportContact200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

