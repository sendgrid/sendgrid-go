# DeleteContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContact**](DeleteContact.md#DeleteContact) | **Delete** /v3/marketing/contacts | Delete Contacts



## DeleteContact

> DeleteContact202Response DeleteContact(ctx, optional)

Delete Contacts

**This endpoint can be used to delete one or more contacts**.  The query parameter `ids` must set to a comma-separated list of contact IDs for bulk contact deletion.  The query parameter `delete_all_contacts` must be set to `\"true\"` to delete **all** contacts.   You must set either `ids` or `delete_all_contacts`.  Deletion jobs are processed asynchronously.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteContactParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteAllContacts** | **string** | Must be set to `\"true\"` to delete all contacts.
**Ids** | **string** | A comma-separated list of contact IDs.

### Return type

[**DeleteContact202Response**](DeleteContact202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

