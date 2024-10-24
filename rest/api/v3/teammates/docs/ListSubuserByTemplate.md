# ListSubuserByTemplate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubuserByTemplate**](ListSubuserByTemplate.md#ListSubuserByTemplate) | **Get** /v3/teammates/{TeammateName}/subuser_access | Get Teammate Subuser Access



## ListSubuserByTemplate

> ListSubuserByTemplate200Response ListSubuserByTemplate(ctx, TeammateNameoptional)

Get Teammate Subuser Access

**This operation allows you to retrieve the Subusers that can be accessed by a specified Teammate.**  This operation will return the Subusers available to a Teammate, including the scopes available. If the Teammate is an administrator, all Subusers will be returned.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TeammateName** | **string** | The username of the Teammate for whom you want to retrieve Subuser access information.

### Other Parameters

Other parameters are passed through a pointer to a ListSubuserByTemplateParams struct


Name | Type | Description
------------- | ------------- | -------------
**AfterSubuserId** | **int32** | The Subuser ID from which the API request will begin retrieving Subusers. This query parameter can be used in successive API calls to retrieve additional Subusers.
**Limit** | **int32** | Limit the number of Subusers to be returned. The default `limit` is `100` per request.
**Username** | **string** | A Subuser's username that will be used to filter the returned result.

### Return type

[**ListSubuserByTemplate200Response**](ListSubuserByTemplate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

