# ListContact

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContact**](ListContact.md#ListContact) | **Get** /v3/marketing/contacts | Get Sample Contacts



## ListContact

> ListContact200Response ListContact(ctx, )

Get Sample Contacts

**This endpoint will return up to 50 of the most recent contacts uploaded or attached to a list**.   This list will then be sorted by email address.  The full contact count is also returned.  Please note that pagination of the contacts has been deprecated.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListContactParams struct


### Return type

[**ListContact200Response**](ListContact200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

