# CreateEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventWebhook**](CreateEventWebhook.md#CreateEventWebhook) | **Post** /v3/user/webhooks/event/settings | Create a new Event Webhook



## CreateEventWebhook

> EventWebhookUnsignedResponse CreateEventWebhook(ctx, optional)

Create a new Event Webhook

**This endpoint allows you to create a new Event Webhook.**  When creating a webhook, you will provide a URL where you want the webhook to send POST requests, and you will select which events you want to receive in those request. See the [**Event Webhook Reference**](https://docs.sendgrid.com/for-developers/tracking-events/event#delivery-events) for details about each event type.  ### Webhook identifiers  When your webhook is succesfully created, you will receive a webhook `id` in the response returned by this endpoint. You can use that ID to [update the webhook's settings](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook), [delete the webhook](https://docs.sendgrid.com/api-reference/webhooks/delete-an-event-webhook), [enable or disable signature verification for the webhook](https://docs.sendgrid.com/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook), and, if signature verification is enabled, [retrieve the webhook's public key](https://docs.sendgrid.com/api-reference/webhooks/get-signed-event-webhooks-public-key).  You may also assign an optional friendly name to each of your webhooks. The friendly name is for convenience only and should not be used to programmatically differentiate your webhooks because it does not need to be unique. Use the webhook ID to reliably differentiate among your webhooks.  ### OAuth  You can optionally configure OAuth verification for your webhook at the time of creation by passing the appropriate values in the `oauth_client_id`, `oauth_client_secret`, and `oauth_token_url` properties. You can enable or disable OAuth for the webhook after creation with the [**Update an Event Webhook**](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) operation.  You may share one OAuth configuration across all your webhooks or create unique credentials for each. See our [webhook security documentation](https://docs.sendgrid.com/for-developers/tracking-events/getting-started-event-webhook-security-features#oauth-20) for details about OAuth and the Event Webhook.  ### Signature verification  Enabling signature verification for your webhook is a separate process and cannot be done at the time of creation with this endpoint. You can use the webhook ID to [enable or disable signature verification with the endpoint dedicated for that operation](https://docs.sendgrid.com/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**EventWebhookRequest** | [**EventWebhookRequest**](EventWebhookRequest.md) | 

### Return type

[**EventWebhookUnsignedResponse**](EventWebhookUnsignedResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

