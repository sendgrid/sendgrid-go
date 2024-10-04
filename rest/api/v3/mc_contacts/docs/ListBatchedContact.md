# ListBatchedContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBatchedContact**](ListBatchedContact.md#ListBatchedContact) | **Post** /v3/marketing/contacts/batch | Get Batched Contacts by IDs



## ListBatchedContact

> ListBatchedContact200Response ListBatchedContact(ctx, optional)

Get Batched Contacts by IDs

**This endpoint is used to retrieve a set of contacts identified by their IDs.**  This can be more efficient endpoint to get contacts than making a series of individual `GET` requests to the \"Get a Contact by ID\" endpoint.  You can supply up to 100 IDs. Pass them into the `ids` field in your request body as an array or one or more strings.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBatchedContactParams struct


Name | Type | Description
------------- | ------------- | -------------
**ListBatchedContactRequest** | [**ListBatchedContactRequest**](ListBatchedContactRequest.md) | 

### Return type

[**ListBatchedContact200Response**](ListBatchedContact200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

