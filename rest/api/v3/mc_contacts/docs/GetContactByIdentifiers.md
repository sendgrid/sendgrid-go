# GetContactByIdentifiers

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContactByIdentifiers**](GetContactByIdentifiers.md#GetContactByIdentifiers) | **Post** /v3/marketing/contacts/search/identifiers/{IdentifierType} | Get Contacts by Identifiers



## GetContactByIdentifiers

> GetContactByIdentifiers200Response GetContactByIdentifiers(ctx, IdentifierTypeoptional)

Get Contacts by Identifiers

**This endpoint allows you to retrieve up to 100 contacts matching the searched identifier values for one type of specified identifier.**  `identifier_type` must be a valid identifier type: `email`, `phone_number_id`, `external_id`, or `anonymous_id`.  This endpoint should be used in place of the [Search Contacts endpoint](https://www.twilio.com/docs/sendgrid/api-reference/contacts/search-contacts) when you can provide exact identifiers and do not need to include other [Segmentation Query Language (SGQL)](https://www.twilio.com/docs/sendgrid/for-developers/sending-email/segmentation-query-language/) filters when searching.  This endpoint returns a `200` status code when any contacts match the identifiers you supplied. When searching multiple identifiers in a single request, it is possible that some will match a contact while others will not. When a partially successful search like this is made, the matching contacts are returned in an object and an error message is returned for the identifiers that are not found.  This endpoint returns a `404` status code when no contacts are found for the provided identifiers.  A `400` status code is returned if any searched addresses are invalid.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdentifierType** | [**IdentifierType1**](IdentifierType1.md) | The type of identifier to search for.

### Other Parameters

Other parameters are passed through a pointer to a GetContactByIdentifiersParams struct


Name | Type | Description
------------- | ------------- | -------------
**GetContactByIdentifiersRequest** | [**GetContactByIdentifiersRequest**](GetContactByIdentifiersRequest.md) | 

### Return type

[**GetContactByIdentifiers200Response**](GetContactByIdentifiers200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

