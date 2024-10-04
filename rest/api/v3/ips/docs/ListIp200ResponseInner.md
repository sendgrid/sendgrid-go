# ListIp200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | An IP address. |
**Subusers** | **[]string** | The subusers that are able to send email from this IP. |
**Rdns** | **string** | The reverse DNS record for this IP address. |[optional] 
**Pools** | **[]string** | The IP pools that this IP has been added to. |
**Warmup** | **bool** | Indicates if this IP address is currently warming up. |
**StartDate** | **float32** | The date that the IP address was entered into warmup. |
**Whitelabeled** | **bool** | Indicates if this IP address is associated with a reverse DNS record. |
**AssignedAt** | **int32** | The date that the IP address was assigned to the user. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


