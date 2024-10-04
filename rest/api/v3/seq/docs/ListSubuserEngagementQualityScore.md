# ListSubuserEngagementQualityScore

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubuserEngagementQualityScore**](ListSubuserEngagementQualityScore.md#ListSubuserEngagementQualityScore) | **Get** /v3/engagementquality/subusers/scores | Get Subusers&#39; Engagement Quality Scores



## ListSubuserEngagementQualityScore

> interface{} ListSubuserEngagementQualityScore(ctx, Dateoptional)

Get Subusers' Engagement Quality Scores

**This operation allows you to retrieve SendGrid Engagement Quality (SEQ) scores for your Subusers or customer accounts for a specific date.** A successful request with this API operation will return either a `200` or `202` response. ### 202 This operation returns a `202` response when SendGrid does not yet have scores available for the specified date range. Scores are calculated asynchronously from requests to this endpoint. This means a score may be available for the specified date at a later time, but a score is not available at the time of your API request. ### 200 A `200` response will include scores for all Subusers or customer accounts belonging to the requesting parent or reseller account. The `score` and `metrics` properties will be omitted from the response if a Subuser or customer account is not eligible to receive a score for the specified date. The `score` property represents a Subuser or customer account's overall engagement quality. The `metrics` property provides additional scores for the input categories that contribute to that overall score. All scores range from `1` to `5` with a higher number representing better engagement quality. See [**SendGrid Engagement Quality Overview**](https://docs.sendgrid.com/api-reference/sendgrid-engagement-quality-api/overview) for more information

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubuserEngagementQualityScoreParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Limit** | **int32** | Specifies the number of results to be returned by the API. This parameter can be used to limit the results returned or in combination with the `after_key` parameter to iterate through paginated results.
**AfterKey** | **int32** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results.

### Return type

**interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

