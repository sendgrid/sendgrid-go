# ListPreBuiltDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPreBuiltDesign**](ListPreBuiltDesign.md#ListPreBuiltDesign) | **Get** /v3/designs/pre-builts | List SendGrid Pre-built Designs



## ListPreBuiltDesign

> ListDesign200Response ListPreBuiltDesign(ctx, optional)

List SendGrid Pre-built Designs

**This endpoint allows you to retrieve a list of pre-built designs provided by Twilio SendGrid**.  Unlike the `/designs` endpoint where *your* designs are stored, a GET request made to `designs/pre-builts` will retrieve a list of the pre-built Twilio SendGrid designs. This endpoint will not return the designs stored in your Design Library.  By default, you will receive 100 results per request; however, you can modify the number of results returned by passing an integer to the `page_size` query parameter.  This endpoint is useful for retrieving the IDs of Twilio SendGrid designs that you want to duplicate and modify.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPreBuiltDesignParams struct


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

