# AddIp201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address that was added to your account. |[optional] 
**IsAutoWarmup** | **bool** | Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address).  This parameter is returned only if the IP address is set to automatically warm up. |[optional] 
**IsParentAssigned** | **bool** | Indicates if a parent on the account is able to send email from the IP address. |[optional] 
**Subusers** | **[]string** | An array of Subuser IDs the IP address was assigned to. |[optional] 
**Region** | [**Region2**](Region2.md) | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


