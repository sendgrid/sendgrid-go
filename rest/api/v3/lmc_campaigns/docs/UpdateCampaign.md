# UpdateCampaign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCampaign**](UpdateCampaign.md#UpdateCampaign) | **Patch** /v3/campaigns/{CampaignId} | Update a Campaign



## UpdateCampaign

> Campaigns2xx UpdateCampaign(ctx, CampaignIdoptional)

Update a Campaign

**This endpoint allows you to update a specific campaign.**  This is especially useful if you only set up the campaign using POST /campaigns, but didn't set many of the parameters.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CampaignId** | **int32** | The id of the campaign you would like to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCampaignParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateACampaignRequest** | [**UpdateACampaignRequest**](UpdateACampaignRequest.md) | 

### Return type

[**Campaigns2xx**](Campaigns2xx.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

