# DeleteSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSegment**](DeleteSegment.md#DeleteSegment) | **Delete** /v3/contactdb/segments/{SegmentId} | Delete a segment



## DeleteSegment

> interface{} DeleteSegment(ctx, SegmentIdoptional)

Delete a segment

**This endpoint allows you to delete a segment from your recipients database.**  You also have the option to delete all the contacts from your Marketing Campaigns recipient database who were in this segment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **int32** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**DeleteContacts** | **bool** | True to delete all contacts matching the segment in addition to deleting the segment
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Body** | **interface{}** | 

### Return type

**interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

