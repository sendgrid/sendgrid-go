# ListClickTrackingSetting

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClickTrackingSetting**](ListClickTrackingSetting.md#ListClickTrackingSetting) | **Get** /v3/tracking_settings/click | Retrieve Click Track Settings



## ListClickTrackingSetting

> ClickTracking ListClickTrackingSetting(ctx, optional)

Retrieve Click Track Settings

**This endpoint allows you to retrieve your current click tracking setting.**  Click Tracking overrides all the links and URLs in your emails and points them to either SendGrid’s servers or the domain with which you branded your link. When a customer clicks a link, SendGrid tracks those [clicks](https://sendgrid.com/docs/glossary/clicks/).  Click tracking helps you understand how users are engaging with your communications. SendGrid can track up to 1000 links per email

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListClickTrackingSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ClickTracking**](ClickTracking.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

