# ListSubuser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubuser**](ListSubuser.md#ListSubuser) | **Get** /v3/subusers | List all Subusers



## ListSubuser

> []Subuser ListSubuser(ctx, optional)

List all Subusers

**This endpoint allows you to retrieve a paginated list of all your subusers.**  You can use the `username` query parameter to filter the list for specific subusers.  You can use the `limit` query parameter to set the page size. If your list contains more items than the page size permits, you can make multiple requests. Use the `offset` query parameter to control the position in the list from which to start retrieving additional items.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubuserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Username** | **string** | The username of this subuser.
**Limit** | **int32** | `limit` sets the page size, i.e. maximum number of items from the list to be returned for a single API request. If omitted, the default page size is used.
**Region** | [**Region**](RegionRegion.md) | Filter for Subusers in this region. If not provided, all Subusers will be returned. All users can also be explicitly requested by using the filter `all`. Users who are not pinned to a region will be displayed as `global`.
**IncludeRegion** | **bool** | Optional flag to include the regions of the Subusers in the response. If not provided, the region will be omitted from the response.
**Offset** | **int32** | The number of items in the list to skip over before starting to retrieve the items for the requested page. The default `offset` of `0` represents the beginning of the list, i.e. the start of the first page. To request the second page of the list, set the `offset` to the page size as determined by `limit`. Use multiples of the page size as your `offset` to request further consecutive pages. E.g. assume your page size is set to `10`. An `offset` of `10` requests the second page, an `offset` of `20` requests the third page and so on, provided there are sufficiently many items in your list.

### Return type

[**[]Subuser**](Subuser.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

