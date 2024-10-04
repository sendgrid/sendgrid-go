# ScheduleCampaign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduleCampaign**](ScheduleCampaign.md#ScheduleCampaign) | **Post** /v3/campaigns/{CampaignId}/schedules | Schedule a Campaign



## ScheduleCampaign

> ScheduleACampaignResponse ScheduleCampaign(ctx, CampaignIdoptional)

Schedule a Campaign

**This endpoint allows you to schedule a specific date and time for your campaign to be sent.**  If you have the flexibility, it's better to schedule mail for off-peak times. Most emails are scheduled and sent at the top of the hour or half hour. Scheduling email to avoid those times (for example, scheduling at 10:53) can result in lower deferral rates because it won't be going through our servers at the same times as everyone else's mail.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CampaignId** | **int32** | 

### Other Parameters

Other parameters are passed through a pointer to a ScheduleCampaignParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**ScheduleACampaignRequest** | [**ScheduleACampaignRequest**](ScheduleACampaignRequest.md) | 

### Return type

[**ScheduleACampaignResponse**](ScheduleACampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

