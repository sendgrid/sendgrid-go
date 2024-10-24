# UpdateGoogleAnalyticsTrackingSetting

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateGoogleAnalyticsTrackingSetting**](UpdateGoogleAnalyticsTrackingSetting.md#UpdateGoogleAnalyticsTrackingSetting) | **Patch** /v3/tracking_settings/google_analytics | Update Google Analytics Settings



## UpdateGoogleAnalyticsTrackingSetting

> GoogleAnalyticsSettings UpdateGoogleAnalyticsTrackingSetting(ctx, optional)

Update Google Analytics Settings

**This endpoint allows you to update your current setting for Google Analytics.**  Google Analytics helps you understand how users got to your site and what they're doing there. For more information about using Google Analytics, please refer to [Google’s URL Builder](https://support.google.com/analytics/answer/1033867?hl=en) and their article on [\"Best Practices for Campaign Building\"](https://support.google.com/analytics/answer/1037445).  We default the settings to Google’s recommendations. For more information, see [Google Analytics Demystified](https://sendgrid.com/docs/ui/analytics-and-reporting/google-analytics/).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateGoogleAnalyticsTrackingSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**GoogleAnalyticsSettings** | [**GoogleAnalyticsSettings**](GoogleAnalyticsSettings.md) | 

### Return type

[**GoogleAnalyticsSettings**](GoogleAnalyticsSettings.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
