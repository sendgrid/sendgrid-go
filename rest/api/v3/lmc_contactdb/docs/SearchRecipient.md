# SearchRecipient

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchRecipient**](SearchRecipient.md#SearchRecipient) | **Post** /v3/contactdb/recipients/search | Search recipients



## SearchRecipient

> SearchRecipient200Response SearchRecipient(ctx, optional)

Search recipients

Search using segment conditions without actually creating a segment. Body contains a JSON object with `conditions`, a list of conditions as described below, and an optional `list_id`, which is a valid list ID for a list to limit the search on.  Valid operators for create and update depend on the type of the field for which you are searching.  - Dates:   - `\"eq\"`, `\"ne\"`, `\"lt\"` (before), `\"gt\"` (after)     - You may use MM/DD/YYYY for day granularity or an epoch for second granularity.   - `\"empty\"`, `\"not_empty\"`   - `\"is within\"`     - You may use an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date format or the # of days. - Text: `\"contains\"`, `\"eq\"` (is - matches the full field), `\"ne\"` (is not - matches any field where the entire field is not the condition value), `\"empty\"`, `\"not_empty\"` - Numbers: `\"eq\"`, `\"lt\"`, `\"gt\"`, `\"empty\"`, `\"not_empty\"` - Email Clicks and Opens: `\"eq\"` (opened), `\"ne\"` (not opened)  Field values must all be a string.  Search conditions using `\"eq\"` or `\"ne\"` for email clicks and opens should provide a \"field\" of either `clicks.campaign_identifier` or `opens.campaign_identifier`.  The condition value should be a string containing the id of a completed campaign.  Search conditions list may contain multiple conditions, joined by an `\"and\"` or `\"or\"` in the `\"and_or\"` field.  The first condition in the conditions list must have an empty `\"and_or\"`, and subsequent conditions must all specify an `\"and_or\"`.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a SearchRecipientParams struct


Name | Type | Description
------------- | ------------- | -------------
**SearchRecipientRequest** | [**SearchRecipientRequest**](SearchRecipientRequest.md) | 

### Return type

[**SearchRecipient200Response**](SearchRecipient200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

