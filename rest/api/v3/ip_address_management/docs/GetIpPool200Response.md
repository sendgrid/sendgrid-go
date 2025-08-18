# GetIpPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the IP Pool. |[optional] 
**Id** | **string** | The unique ID of the IP Pool. |[optional] 
**IpsPreview** | **[]string** | A sample list of IP addresses associated with the IP Pool. The sample is limited to 10 results. |[optional] 
**TotalIpCount** | **int32** | The total number of IP addresses in the IP Pool. An IP Pool can have a maximum of 100 associated IP addresses. |[optional] 
**IpCountByRegion** | [**[]GetIpPool200ResponseIpCountByRegionInner**](GetIpPool200ResponseIpCountByRegionInner.md) | The total number of IP addresses by region. this object is only returned if the `include_region` parameter is included and set to `true` in the API request. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


