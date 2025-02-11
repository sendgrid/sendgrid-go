# ListIp200ResponseMetadataNextParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterKey** | **string** | Specifies which items to be returned by the API. When the `after_key` is specified, the API will return items beginning from the first item after the item specified. This parameter can be used in combination with `limit` to iterate through paginated results. The `after_key` cannot be used in combination with the `before_key` parameter. |[optional] 
**BeforeKey** | **string** | Specifies which items to be returned by the API. When the `before_key` is specified, the API will return items beginning from the first item before the item specified. This parameter can be used in combination with `limit` to iterate through paginated results. The `before_key` cannot be used in combination with the `after_key` parameter. |[optional] 
**Ip** | **string** | The IP address specified in the request with the `ip` query parameter. This parameter is returned only when an IP is included in the request. |[optional] 
**IsLeased** | **bool** | Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account. This parameter is returned only if the IP address is leased. |[optional] 
**IsEnabled** | **bool** | Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses. This parameter is returned only if the IP address is enabled. |[optional] 
**IsParentAssigned** | **bool** | Indicates if a parent on the account is able to send email from the IP address. This parameter is returned only if the IP address is parent assigned. |[optional] 
**Pool** | **string** | The IP Pool ID specified in the request with the `pool` query parameter. This parameter is returned only when an IP Pool is included in the request. |[optional] 
**StartAddedAt** | **string** | The beginning of the time window specified in the request with the `start_added_at` query parameter. This parameter is returned only when the `start_added_at` parameter is included in the request. |[optional] 
**EndAddedAt** | **string** | The end of the time window specified in the request with the `end_added_at` query parameter. This parameter is returned only when the `end_added_at` parameter is included in the request. |[optional] 
**Limit** | **string** | The number of items returned in the request. This parameter is returned only when a `limit` is set using the `limit` query parameter in the request. |[optional] 
**Region** | [**Region1**](Region1.md) | The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request, |[optional] 
**IncludeRegion** | **string** | Indicates whether or not to include the IP address's region. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


