# UpdateSignedEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSignedEventWebhook**](UpdateSignedEventWebhook.md#UpdateSignedEventWebhook) | **Patch** /v3/user/webhooks/event/settings/signed/{Id} | Toggle signature verification for a single Event Webhook by ID



## UpdateSignedEventWebhook

> GetSignedEventWebhook200Response UpdateSignedEventWebhook(ctx, Idoptional)

Toggle signature verification for a single Event Webhook by ID

**This endpoint allows you to enable or disable signature verification for a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will enable signature verification for your oldest webhook by `created_date`. This means the default webhook operated on by this endpoint when no ID is provided will be the first one you created. This functionality allows customers who do not have multiple webhooks to enable or disable signature verifiction for their only webhook, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  This endpoint accepts a single boolean request property, `enabled`, that can be set `true` or `false` to enable or disable signature verification. This endpoint will return the public key required to verify Twilio SendGrid signatures if it is enabled or an empty string if signing is disabled. You can also retrieve your public key using the [**Get an Event Webhook's Public Key**](https://docs.sendgrid.com/api-reference/webhooks/get-signed-event-webhooks-public-key) endpoint.  For more information about cryptographically signing the Event Webhook, see [**Getting Started with the Event Webhook Security Features**](https://sendgrid.com/docs/for-developers/tracking-events/getting-started-event-webhook-security-features).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Event Webhook you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSignedEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateSignedEventWebhookRequest** | [**UpdateSignedEventWebhookRequest**](UpdateSignedEventWebhookRequest.md) | 

### Return type

[**GetSignedEventWebhook200Response**](GetSignedEventWebhook200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

