# ListContactByEmail

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContactByEmail**](ListContactByEmail.md#ListContactByEmail) | **Post** /v3/marketing/contacts/search/emails | Get Contacts by Emails



## ListContactByEmail

> ListContactByEmail200Response ListContactByEmail(ctx, optional)

Get Contacts by Emails

**This endpoint allows you to retrieve up to 100 contacts matching the searched `email` address(es), including any `alternate_emails`.**   Email addresses are unique to a contact, meaning this endpoint can treat an email address as a primary key to search by. The contact object associated with the address, whether it is their `email` or one of their `alternate_emails` will be returned if matched.  Email addresses in the search request do not need to match the case in which they're stored, but the email addresses in the result will be all lower case. Empty strings are excluded from the search and will not be returned.  This endpoint should be used in place of the \"Search Contacts\" endpoint when you can provide exact email addresses and do not need to include other [Segmentation Query Language (SGQL)](https://sendgrid.com/docs/for-developers/sending-email/segmentation-query-language/) filters when searching.  If you need to access a large percentage of your contacts, we recommend exporting your contacts with the \"Export Contacts\" endpoint and filtering the client side results.  This endpoint returns a `200` status code when any contacts match the address(es) you supplied. When searching multiple addresses in a single request, it is possible that some addresses will match a contact while others will not. When a partially successful search like this is made, the matching contacts are returned in an object and an error message is returned for the email address(es) that are not found.   This endpoint returns a `404` status code when no contacts are found for the provided email address(es).  A `400` status code is returned if any searched addresses are invalid.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListContactByEmailParams struct


Name | Type | Description
------------- | ------------- | -------------
**ListContactByEmailRequest** | [**ListContactByEmailRequest**](ListContactByEmailRequest.md) | 

### Return type

[**ListContactByEmail200Response**](ListContactByEmail200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

