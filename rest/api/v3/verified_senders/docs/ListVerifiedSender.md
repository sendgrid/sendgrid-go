# ListVerifiedSender

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVerifiedSender**](ListVerifiedSender.md#ListVerifiedSender) | **Get** /v3/verified_senders | Get All Verified Senders



## ListVerifiedSender

> ListVerifiedSender200Response ListVerifiedSender(ctx, optional)

Get All Verified Senders

**This endpoint allows you to retrieve all the Sender Identities associated with an account.**  This endpoint will return both verified and unverified senders.  You can limit the number of results returned using the `limit`, `lastSeenID`, and `id` query string parameters.  * `limit` allows you to specify an exact number of Sender Identities to return. * `lastSeenID` will return senders with an ID number occuring after the passed in ID. In other words, the `lastSeenID` provides a starting point from which SendGrid will iterate to find Sender Identities associated with your account. * `id` will return information about only the Sender Identity passed in the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerifiedSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **float32** | Specifies the number of results to be returned by the API. This parameter can be used to limit the results returned or in combination with the `lastSeenID` parameter to iterate through paginated results.
**LastSeenID** | **float32** | Returns senders with an ID number occurring after the passed in ID. In other words, the `lastSeenID` provides a starting point from which SendGrid will iterate to find Sender Identities associated with your account.
**Id** | **int32** | Returns information about only the Sender Identity passed in the request.

### Return type

[**ListVerifiedSender200Response**](ListVerifiedSender200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

