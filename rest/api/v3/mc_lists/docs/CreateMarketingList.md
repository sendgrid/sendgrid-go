# CreateMarketingList

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMarketingList**](CreateMarketingList.md#CreateMarketingList) | **Post** /v3/marketing/lists | Create List



## CreateMarketingList

> List CreateMarketingList(ctx, optional)

Create List

**This endpoint creates a new contacts list.**  Once you create a list, you can use the UI to [trigger an automation](https://sendgrid.com/docs/ui/sending-email/getting-started-with-automation/#create-an-automation) every time you add a new contact to the list.  A link to the newly created object is in `_metadata`.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateMarketingListParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateMarketingListRequest** | [**CreateMarketingListRequest**](CreateMarketingListRequest.md) | 

### Return type

[**List**](List.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

