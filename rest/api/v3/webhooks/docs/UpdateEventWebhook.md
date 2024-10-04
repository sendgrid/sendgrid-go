# UpdateEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateEventWebhook**](UpdateEventWebhook.md#UpdateEventWebhook) | **Patch** /v3/user/webhooks/event/settings/{Id} | Update a single Event Webhook by ID.



## UpdateEventWebhook

> EventWebhookUnsignedResponse UpdateEventWebhook(ctx, Idoptional)

Update a single Event Webhook by ID.

**This endpoint allows you to update a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will update and return your oldest webhook by `created_date`. This means the default webhook updated by this endpoint when no ID is provided will be the first one you created. This functionality allows customers who do not have multiple webhooks to use this endpoint to update their only webhook, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  ### Enable or disable the webhook  You can set the `enabled` property to `true` to enable the webhook or `false` to disable it. Disabling a webhook will not delete it from your account, but it will prevent the webhook from sending events to your designated URL.  ### URL  A webhook's URL is the endpoint where you want the webhook to send POST requests containing event data. No more than one webhook may be configured to send to the same URL. SendGrid will return an error if you attempt to set a URL for a webhook that is already in use by the user on another webhook.  ### Event settings  If an event type is marked as `true`, the event webhook will send information about that event type. See the [**Event Webhook Reference**](https://docs.sendgrid.com/for-developers/tracking-events/event#delivery-events) for details about each event type.  ### Webhook identifiers  You may assign an optional friendly name to each of your webhooks. The friendly name is for convenience only and should not be used to programmatically differentiate your webhooks because it does not need to be unique.  ### OAuth  You can configure OAuth for your webhook by passing the required values to this endpoint in the `oauth_client_id`, `oauth_client_secret`, and `oauth_token_url` properties. To disable OAuth, pass an empty string to this endpoint for each of the OAuth properties. You may share one OAuth configuration across all your webhooks or create unique credentials for each. See our [webhook security documentation](https://docs.sendgrid.com/for-developers/tracking-events/getting-started-event-webhook-security-features#oauth-20) for more detailed information about OAuth and the Event Webhook.  ### Signature verification  Enabling signature verification for your webhook is a separate process and cannot be done with this endpoint. You can use the webhook ID to [enable or disable signature verification with the endpoint dedicated for that operation](https://docs.sendgrid.com/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Event Webhook you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a UpdateEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Include** | **string** | Use this to include optional fields in the response payload. When this is set to `include=account_status_change`, the `account_status_change` field will be part of the response payload and set to `false` by default. See [Update an event webhook](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) for enabling this webhook notification which lets you subscribe to account status change events related to compliance action taken by SendGrid.
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

