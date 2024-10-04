# UpdateOpenTrackingSetting

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateOpenTrackingSetting**](UpdateOpenTrackingSetting.md#UpdateOpenTrackingSetting) | **Patch** /v3/tracking_settings/open | Update Open Tracking Settings



## UpdateOpenTrackingSetting

> ListOpenTrackingSetting200Response UpdateOpenTrackingSetting(ctx, optional)

Update Open Tracking Settings

**This endpoint allows you to update your current settings for open tracking.**  Open Tracking adds an invisible image at the end of the email which can track email opens.  If the email recipient has images enabled on their email client, a request to SendGrid’s server for the invisible image is executed and an open event is logged.  These events are logged in the Statistics portal, Email Activity interface, and are reported by the Event Webhook.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOpenTrackingSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateOpenTrackingSettingRequest** | [**UpdateOpenTrackingSettingRequest**](UpdateOpenTrackingSettingRequest.md) | 

### Return type

[**ListOpenTrackingSetting200Response**](ListOpenTrackingSetting200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

