# DeleteEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventWebhook**](DeleteEventWebhook.md#DeleteEventWebhook) | **Delete** /v3/user/webhooks/event/settings/{Id} | Delete a single Event Webhook by ID.



## DeleteEventWebhook

> DeleteEventWebhook(ctx, Idoptional)

Delete a single Event Webhook by ID.

**This endpoint allows you to delete a single Event Webhook by ID.**  Unlike the [**Get an Event Webhook**](https://docs.sendgrid.com/api-reference/webhooks/get-an-event-webhook) and [**Update an Event Webhook**](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) endpoints, which will operate on your oldest webhook by `created_date` when you don't provide an ID, this endpoint will return an error if you do not pass it an ID. This behavior prevents customers from unintentionally deleting a webhook. You can retrieve your webhooks' IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  ### Enable or disable the webhook  This endpoint will permanently delete the webhook specified. If you instead want to disable a webhook, you can set the `enabled` property to `false` with the [**Update an Event Webhook**](https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook) endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Event Webhook you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

