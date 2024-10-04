# DuplicatePreBuiltDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DuplicatePreBuiltDesign**](DuplicatePreBuiltDesign.md#DuplicatePreBuiltDesign) | **Post** /v3/designs/pre-builts/{Id} | Duplicate SendGrid Pre-built Design



## DuplicatePreBuiltDesign

> DesignOutput DuplicatePreBuiltDesign(ctx, Idoptional)

Duplicate SendGrid Pre-built Design

**This endpoint allows you to duplicate one of the pre-built Twilio SendGrid designs**.  Like duplicating one of your existing designs, you are not required to pass any data in the body of a request to this endpoint. If you choose to leave the `name` field blank, your duplicate will be assigned the name of the design it was copied from with the text \"Duplicate: \" prepended to it. This name change is only a convenience, as the duplicate design will be assigned a unique ID that differentiates it from your other designs. You can retrieve the IDs for Twilio SendGrid pre-built designs using the \"List SendGrid Pre-built Designs\" endpoint.  You can modify your duplicate’s name at the time of creation by passing an updated value to the `name` field when making the initial request. More on retrieving design IDs can be found above.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the pre-built Design you want to duplicate.

### Other Parameters

Other parameters are passed through a pointer to a DuplicatePreBuiltDesignParams struct


Name | Type | Description
------------- | ------------- | -------------
**DesignDuplicateInput** | [**DesignDuplicateInput**](DesignDuplicateInput.md) | 

### Return type

[**DesignOutput**](DesignOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

