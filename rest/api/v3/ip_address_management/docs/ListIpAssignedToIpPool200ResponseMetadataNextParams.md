# ListIpAssignedToIpPool200ResponseMetadataNextParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterKey** | Pointer to **string** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results. The `after_key` cannot be used in combination with the `before_key` parameter. |
**Limit** | **string** | The `limit` specified in the request. This parameter will be included only if it was specified in the request. This is not the default limit enforced by the API. |[optional] 
**IncludeRegion** | **string** | Indicates whether or not to include the IP address's region. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


