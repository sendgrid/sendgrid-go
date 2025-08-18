# ListIpAssignedToIpPool200ResponseResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | An IP address assigned to the IP Pool. |[optional] 
**Region** | [**Region6**](Region6.md) | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 
**Pools** | [**[]ListIp200ResponseResultInnerPoolsInner**](ListIp200ResponseResultInnerPoolsInner.md) | IP Pools the IP address is assigned to. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


