# ListContactCount

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContactCount**](ListContactCount.md#ListContactCount) | **Get** /v3/marketing/contacts/count | Get Total Contact Count



## ListContactCount

> ListContactCount200Response ListContactCount(ctx, )

Get Total Contact Count

**This endpoint returns the total number of contacts you have stored.**   Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListContactCountParams struct


### Return type

[**ListContactCount200Response**](ListContactCount200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

