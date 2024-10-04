# AddIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoWarmup** | **bool** | Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address). |
**IsParentAssigned** | **bool** | Indicates if a parent on the account is able to send email from the IP address. |
**Subusers** | **[]string** | An array of Subuser IDs the IP address will be assigned to. |[optional] 
**Region** | [**Region3**](Region3.md) | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 
**IncludeRegion** | **bool** | Boolean indicating whether or not to return the IP address's region information in the response. |[optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


