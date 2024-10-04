# SearchContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchContact**](SearchContact.md#SearchContact) | **Post** /v3/marketing/contacts/search | Search Contacts



## SearchContact

> SearchContact200Response SearchContact(ctx, optional)

Search Contacts

**Use this endpoint to locate contacts**.  The request body's `query` field accepts valid [SGQL](https://sendgrid.com/docs/for-developers/sending-email/segmentation-query-language/) for searching for a contact.  Because contact emails are stored in lower case, using SGQL to search by email address requires the provided email address to be in lower case. The SGQL `lower()` function can be used for this.  Only the first 50 contacts that meet the search criteria will be returned.  If the query takes longer than 20 seconds, a `408 Request Timeout` status will be returned.  Formatting the `created_at` and `updated_at` values as Unix timestamps is deprecated. Instead, they are returned as ISO format as string.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a SearchContactParams struct


Name | Type | Description
------------- | ------------- | -------------
**SearchContactRequest** | [**SearchContactRequest**](SearchContactRequest.md) | 

### Return type

[**SearchContact200Response**](SearchContact200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

