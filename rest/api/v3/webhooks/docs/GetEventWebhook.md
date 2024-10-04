# GetEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventWebhook**](GetEventWebhook.md#GetEventWebhook) | **Get** /v3/user/webhooks/event/settings/{Id} | Get the settings for a single Event Webhook.



## GetEventWebhook

> EventWebhookNoDatesResponse GetEventWebhook(ctx, Idoptional)

Get the settings for a single Event Webhook.

**This endpoint allows you to retrieve a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will return your oldest webhook by `created_date`. This means the default webhook returned by this endpoint when no ID is provided will be the first one you created. This functionality allows customers who do not have multiple webhooks to use this endpoint to retrieve their only webhook, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  ### Event settings  Your webhook will be returned with all of its settings, which include the events that will be included in the POST request by the webhook and the URL where they will be sent. If an event type is marked as `true`, the event webhook will send information about that event type. See the [**Event Webhook Reference**](https://docs.sendgrid.com/for-developers/tracking-events/event#delivery-events) for details about each event type.  ### Signature verification  The `public_key` property will be returned only for webhooks with signature verification enabled.  ### OAuth  You may share one OAuth configuration across all your webhooks or create unique credentials for each. The OAuth properties will be returned only for webhooks with OAuth configured.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Event Webhook you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a GetEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Include** | **string** | Use this to include optional fields in the response payload. When this is set to `include=account_status_change`, the `account_status_change` field will be part of the response payload and set to `false` by default. See [Update an event webhook](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) for enabling this webhook notification which lets you subscribe to account status change events related to compliance action taken by SendGrid.

### Return type

[**EventWebhookNoDatesResponse**](EventWebhookNoDatesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

