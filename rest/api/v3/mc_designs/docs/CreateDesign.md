# CreateDesign

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDesign**](CreateDesign.md#CreateDesign) | **Post** /v3/designs | Create Design



## CreateDesign

> DesignOutput CreateDesign(ctx, optional)

Create Design

**This endpoint allows you to create a new design**.  You can add a new design by passing data, including a string of HTML email content, to `/designs`. When creating designs from scratch, be aware of the styling constraints inherent to many email clients. For a list of best practices, see our guide to [Cross-Platform Email Design](https://sendgrid.com/docs/ui/sending-email/cross-platform-html-design/).  The Design Library can also convert your design’s HTML elements into drag and drop modules that are editable in the Designs Library user interface. For more, visit the [Design and Code Editor documentation](https://sendgrid.com/docs/ui/sending-email/editor/#drag--drop-markup).  Because the `/designs` endpoint makes it easy to add designs, you can create a design with your preferred tooling or migrate designs you already own without relying on the Design Library UI.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateDesignParams struct


Name | Type | Description
------------- | ------------- | -------------
**DesignInput** | [**DesignInput**](DesignInput.md) | 

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

