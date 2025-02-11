# GetSignedEventWebhook

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignedEventWebhook**](GetSignedEventWebhook.md#GetSignedEventWebhook) | **Get** /v3/user/webhooks/event/settings/signed/{Id} | Get Signed Event Webhook&#39;s Public Key



## GetSignedEventWebhook

> GetSignedEventWebhook200Response GetSignedEventWebhook(ctx, Idoptional)

Get Signed Event Webhook's Public Key

**This endpoint allows you to retrieve the public key for a single Event Webhook by ID.**  If you do not pass a webhook ID to this endpoint, it will return the public key for your oldest webhook by `created_date`. This means the default key returned by this endpoint when no ID is provided will be for the first webhook you created. This functionality allows customers who do not have multiple webhooks to use this endpoint to retrieve their only webhook's public key, even if they do not supply an ID. If you have multiple webhooks, you can retrieve their IDs using the [**Get All Event Webhooks**](https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks) endpoint.  Once you have enabled signature verification for a webhook, you will need the public key provided to verify the signatures on requests coming from Twilio SendGrid. You can use the webhook ID to [enable or disable signature verification with the endpoint dedicated for that operation](https://docs.sendgrid.com/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook).  For more information about cryptographically signing the Event Webhook, see [**Getting Started with the Event Webhook Security Features**](https://sendgrid.com/docs/for-developers/tracking-events/getting-started-event-webhook-security-features).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the Event Webhook you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a GetSignedEventWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**GetSignedEventWebhook200Response**](GetSignedEventWebhook200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

