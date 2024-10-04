# UpdateTeammate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateTeammate**](UpdateTeammate.md#UpdateTeammate) | **Patch** /v3/teammates/{Username} | Update teammate&#39;s permissions



## UpdateTeammate

> UpdateTeammate200Response UpdateTeammate(ctx, Usernameoptional)

Update teammate's permissions

**This endpoint allows you to update a teammate’s permissions.**  To turn a teammate into an admin, the request body should contain an `is_admin` set to `true`. Otherwise, set `is_admin` to `false` and pass in all the scopes that a teammate should have.  **Only the parent user or other admin teammates can update another teammate’s permissions.**  **Admin users can only update permissions.**

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Username** | **string** | The username of the teammate that you want to retrieve.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTeammateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**UpdateTeammateRequest** | [**UpdateTeammateRequest**](UpdateTeammateRequest.md) | 

### Return type

[**UpdateTeammate200Response**](UpdateTeammate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

