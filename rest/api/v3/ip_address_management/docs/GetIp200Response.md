# GetIp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address specified in the request. |[optional] 
**IsParentAssigned** | **bool** | Indicates if a parent on the account is able to send email from the IP address. |[optional] 
**IsAutoWarmup** | **bool** | Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address). |[optional] 
**Pools** | [**[]GetIp200ResponsePoolsInner**](GetIp200ResponsePoolsInner.md) | An array of IP Pools the IP address is assigned to. |[optional] 
**AddedAt** | **int32** | A timestamp indicating when the IP address was added to your account. |[optional] 
**UpdatedAt** | **int32** | A timestamp indicating when the IP was last updated. |[optional] 
**IsEnabled** | **bool** | Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses. |[optional] 
**IsLeased** | **bool** | Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account. |[optional] 
**Region** | **string** | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. Possible values are `us` or `eu`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


