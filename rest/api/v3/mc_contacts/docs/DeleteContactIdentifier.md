# DeleteContactIdentifier

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContactIdentifier**](DeleteContactIdentifier.md#DeleteContactIdentifier) | **Delete** /v3/marketing/contacts/{ContactId}/identifiers | Delete a Contact Identifier



## DeleteContactIdentifier

> DeleteContactIdentifier202Response DeleteContactIdentifier(ctx, ContactIdoptional)

Delete a Contact Identifier

**This endpoint can be used to delete one identifier from a contact.**  Deletion jobs are processed asynchronously.  Note this is different from deleting a contact. If the contact has only one identifier, the asynchronous request will fail. All contacts are required to have at least one identifier.  The request body field `identifier_type` must have a valid value of \"EMAIL\", \"PHONENUMBERID\", \"EXTERNALID\", or \"ANONYMOUSID\". 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ContactId** | **string** | Must be set to the `contact_id` of the contact you want to remove the identifier from.

### Other Parameters

Other parameters are passed through a pointer to a DeleteContactIdentifierParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteContactIdentifierRequest** | [**DeleteContactIdentifierRequest**](DeleteContactIdentifierRequest.md) | 

### Return type

[**DeleteContactIdentifier202Response**](DeleteContactIdentifier202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

