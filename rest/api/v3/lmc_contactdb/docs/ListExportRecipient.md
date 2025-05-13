# ListExportRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExportRecipient**](ListExportRecipient.md#ListExportRecipient) | **Get** /v3/contactdb/exports | Get All Existing Exports



## ListExportRecipient

> ListExportRecipient200Response ListExportRecipient(ctx, )

Get All Existing Exports

**Use this endpoint to retrieve details of all current exported jobs**.  It will return an array of objects, each of which records an export job in flight or recently completed.   Each object's `export_type` field will tell you which kind of export it is and its `status` field will indicate what stage of processing it has reached. Exports which are `ready` will be accompanied by a `urls` field which lists the URLs of the export's downloadable files — there will be more than one if you specified a maximum file size in your initial export request.  Use this endpoint if you have exports in flight but do not know their IDs, which are required for the \"Export Recipients Status\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListExportRecipientParams struct


### Return type

[**ListExportRecipient200Response**](ListExportRecipient200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

