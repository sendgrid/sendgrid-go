# TestEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestEventWebhook**](TestEventWebhook.md#TestEventWebhook) | **Post** /v3/user/webhooks/event/test | Test an Event Webhook&#39;s settings



## TestEventWebhook

> TestEventWebhook(ctx, optional)

Test an Event Webhook's settings

**This endpoint allows you to test an Event Webhook.**  Retry logic for this endpoint differs from other endpoints, which use a rolling 24-hour retry.  This endpoint will make a POST request with a fake event notification to a URL you provide. This allows you to verify that you have properly configured the webhook before sending real data to your URL.  ### Test OAuth configuration  To test your OAuth configuration, you must include the necessary OAuth properties: `oauth_client_id`, `oauth_client_secret`, and `oauth_token_url`.  If the webhook you are testing already has OAuth credentials saved, you will provide only the `oauth_client_id` and `oauth_token_url`—we will pull the secret for you. If you are testing a new set of OAuth credentials that have not been saved with SendGrid, you must provide all three property values.  You can retrieve a previously saved `oauth_client_id` and `oauth_token_url` from the [**Get an Event Webhook**](https://docs.sendgrid.com/api-reference/webhooks/get-an-event-webhook) endpoint; however, for security reasons, SendGrid will not provide your `oauth_client_secret`.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a TestEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**EventWebhookTestRequest** | [**EventWebhookTestRequest**](EventWebhookTestRequest.md) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

