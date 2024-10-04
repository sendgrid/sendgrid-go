# UpdateContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateContact**](UpdateContact.md#UpdateContact) | **Put** /v3/marketing/contacts | Add or Update a Contact



## UpdateContact

> UpdateContact202Response UpdateContact(ctx, optional)

Add or Update a Contact

**This endpoint allows the [upsert](https://en.wiktionary.org/wiki/upsert) (insert or update) of up to 30,000 contacts, or 6MB of data, whichever is lower**. Because the creation and update of contacts is an asynchronous process, the response will not contain immediate feedback on the processing of your upserted contacts. Rather, it will contain an HTTP 202 response indicating the contacts are queued for processing or an HTTP 4XX error containing validation errors. Should you wish to get the resulting contact's ID or confirm that your contacts have been updated or added, you can use the [Get Contacts by Identifiers operation](https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-contacts-by-identifiers). Please note that custom fields need to have been already created if you wish to set their values for the contacts being upserted. To do this, please use the [Create Custom Field Definition endpoint](https://www.twilio.com/docs/sendgrid/api-reference/custom-fields/create-custom-field-definition). You will see a `job_id` in the response to your request. This can be used to check the status of your upsert job. To do so, please use the [Import Contacts Status endpoint](https://www.twilio.com/docs/sendgrid/api-reference/contacts/import-contacts-status). If the contact already exists in the system, any entries submitted via this endpoint will update the existing contact. In order to update a contact, all of its existing identifiers must be present. Any fields omitted from the request will remain as they were. A contact's ID cannot be used to update the contact. The email field will be changed to all lower-case. If a contact is added with an email that exists but contains capital letters, the existing contact with the all lower-case email will be updated.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateContactParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateContactRequest** | [**UpdateContactRequest**](UpdateContactRequest.md) | 

### Return type

[**UpdateContact202Response**](UpdateContact202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

