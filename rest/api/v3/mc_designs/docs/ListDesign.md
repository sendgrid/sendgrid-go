# ListDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDesign**](ListDesign.md#ListDesign) | **Get** /v3/designs | List Designs



## ListDesign

> ListDesign200Response ListDesign(ctx, optional)

List Designs

**This endpoint allows you to retrieve a list of designs already stored in your Design Library**.  A GET request to `/designs` will return a list of your existing designs. This endpoint will not return the pre-built Twilio SendGrid designs. Pre-built designs can be retrieved using the `/designs/pre-builts` endpoint, which is detailed below.  By default, you will receive 100 results per request; however, you can modify the number of results returned by passing an integer to the `page_size` query parameter.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDesignParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | number of results to return
**PageToken** | **string** | token corresponding to a specific page of results, as provided by metadata
**Summary** | **bool** | set to false to return all fields

### Return type

[**ListDesign200Response**](ListDesign200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

