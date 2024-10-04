# ListIp200ResponseResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | An IP address on your account. |[optional] 
**Pools** | [**[]ListIp200ResponseResultInnerPoolsInner**](ListIp200ResponseResultInnerPoolsInner.md) | An array of IP Pools the IP address is assigned to. |[optional] 
**IsAutoWarmup** | **bool** | Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address). |[optional] 
**IsParentAssigned** | Pointer to **bool** | Indicates if a parent on the account is able to send email from the IP address. This parameter will be returned only if the request was made by the parent account. |
**UpdatedAt** | Pointer to **int32** | A timestamp indicating when the IP was last updated. |
**IsEnabled** | Pointer to **bool** | Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses. |
**IsLeased** | **bool** | Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account. |[optional] 
**AddedAt** | **int32** | A timestamp representing when the IP address was added to your account. |[optional] 
**Region** | [**Region**](Region.md) | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


