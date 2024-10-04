# ListReputation

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListReputation**](ListReputation.md#ListReputation) | **Get** /v3/subusers/reputations | Retrieve Subuser Reputations



## ListReputation

> []ListReputation200ResponseInner ListReputation(ctx, optional)

Retrieve Subuser Reputations

**This endpoint allows you to request the reputations for your subusers.**  Subuser sender reputations give a good idea how well a sender is doing with regards to how recipients and recipient servers react to the mail that is being received. When a bounce, spam report, or other negative action happens on a sent email, it will affect your sender rating.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListReputationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Usernames** | **string** | 

### Return type

[**[]ListReputation200ResponseInner**](ListReputation200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

