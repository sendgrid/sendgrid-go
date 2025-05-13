# GetSuppressionBouncesClassifications

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSuppressionBouncesClassifications**](GetSuppressionBouncesClassifications.md#GetSuppressionBouncesClassifications) | **Get** /v3/suppression/bounces/classifications/{Classification} | Retrieve bounce classification over time by domain stats



## GetSuppressionBouncesClassifications

> GetSuppressionBouncesClassifications200Response GetSuppressionBouncesClassifications(ctx, AcceptClassificationoptional)

Retrieve bounce classification over time by domain stats

This endpoint will return the number of bounces for the classification specified in descending order for each day. You can retrieve the bounce classification totals in CSV format by specifying `\"text/csv\"` in the Accept header.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Classification** | [**Classification1**](Classification1.md) | The classification you want to filter by. Possible values are: `Content`, `Frequency or Volume Too High`, `Invalid Address`, `Mailbox Unavailable`, `Reputation`, `Technical Failure`, `Unclassified`.

### Other Parameters

Other parameters are passed through a pointer to a GetSuppressionBouncesClassificationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**StartDate** | **string** | The start of the time range, in YYYY-MM-DD format, when a bounce was created (inclusive).
**EndDate** | **string** | The end of the time range, in YYYY-MM-DD format, when a bounce was created (inclusive).
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**GetSuppressionBouncesClassifications200Response**](GetSuppressionBouncesClassifications200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

