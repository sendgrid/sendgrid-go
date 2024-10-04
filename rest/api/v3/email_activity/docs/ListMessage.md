# ListMessage

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMessage**](ListMessage.md#ListMessage) | **Get** /v3/messages | Filter all messages



## ListMessage

> ListMessage200Response ListMessage(ctx, Queryoptional)

Filter all messages

Filter all messages to search your Email Activity. All queries must be [URL encoded](https://meyerweb.com/eric/tools/dencoder/), and use the following format:  `query={query_type}=\"{query_content}\"`   Once URL encoded, the previous query will look like this:  `query=type%3D%22query_content%22`  For example, to filter by a specific email, use the following query:  `query=to_email%3D%22example%40example.com%22`  Visit our [Query Reference section](https://docs.sendgrid.com/for-developers/sending-email/getting-started-email-activity-api#query-reference) to see a full list of basic query types and examples.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **float32** | The number of messages returned. This parameter must be greater than 0 and less than or equal to 1000

### Return type

[**ListMessage200Response**](ListMessage200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

