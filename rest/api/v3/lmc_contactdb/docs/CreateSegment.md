# CreateSegment

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](CreateSegment.md#CreateSegment) | **Post** /v3/contactdb/segments | Create a Segment



## CreateSegment

> ContactdbSegmentsId200 CreateSegment(ctx, optional)

Create a Segment

**This endpoint allows you to create a new segment.**     Valid operators for create and update depend on the type of the field for which you are searching.  **Dates** - \"eq\", \"ne\", \"lt\" (before), \"gt\" (after)     - You may use MM/DD/YYYY for day granularity or an epoch for second granularity. - \"empty\", \"not_empty\" - \"is within\"     - You may use an [ISO 8601 date format](https://en.wikipedia.org/wiki/ISO_8601) or the # of days.  **Text** - \"contains\" - \"eq\" (is/equals - matches the full field) - \"ne\" (is not/not equals - matches any field where the entire field is not the condition value) - \"empty\" - \"not_empty\"  **Numbers** - \"eq\" (is/equals) - \"lt\" (is less than) - \"gt\" (is greater than) - \"empty\" - \"not_empty\"  **Email Clicks and Opens** - \"eq\" (opened) - \"ne\" (not opened)  All field values must be a string.   Conditions using \"eq\" or \"ne\" for email clicks and opens should provide a \"field\" of either `clicks.campaign_identifier` or `opens.campaign_identifier`. The condition value should be a string containing the id of a completed campaign.   The conditions list may contain multiple conditions, joined by an \"and\" or \"or\" in the \"and_or\" field.  The first condition in the conditions list must have an empty \"and_or\", and subsequent conditions must all specify an \"and_or\".

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSegmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**ContactdbSegments** | [**ContactdbSegments**](ContactdbSegments.md) | 

### Return type

[**ContactdbSegmentsId200**](ContactdbSegmentsId200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

