# ListAutomationStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAutomationStat**](ListAutomationStat.md#ListAutomationStat) | **Get** /v3/marketing/stats/automations | Get All Automation Stats



## ListAutomationStat

> AutomationsResponse ListAutomationStat(ctx, optional)

Get All Automation Stats

**This endpoint allows you to retrieve stats for all your Automations.**  By default, all of your Automations will be returned, but you can specify a selection by passing in a comma-separated list of Automation IDs as the value of the query string parameter `automation_ids`.  Responses are paginated. You can limit the number of responses returned per batch using the `page_size` query string parameter. The default is 25, but you can specify a value between 1 and 50.  You can retrieve a specific page of responses with the `page_token` query string parameter.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAutomationStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**AutomationIds** | **[]string** | This endpoint returns all automation IDs if no `automation_ids` are specified.
**PageSize** | **int32** | The number of elements you want returned on each page.
**PageToken** | **string** | The stats endpoints are paginated. To get the next page, call the passed `_metadata.next` URL. If `_metadata.prev` doesn't exist, you're at the first page. Similarly, if `_metadata.next` is not present, you're at the last page.

### Return type

[**AutomationsResponse**](AutomationsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

