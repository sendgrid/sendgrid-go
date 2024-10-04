# ReverseDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the Reverse DNS. |
**Ip** | **string** | The IP address that this Reverse DNS was created for. |
**Rdns** | **string** | The reverse DNS record for the IP address. This points to the Reverse DNS subdomain. |
**Users** | [**[]ReverseDnsUsersInner**](ReverseDnsUsersInner.md) | The users who are able to send mail from the IP address. |
**Subdomain** | **string** | The subdomain created for this reverse DNS. This is where the rDNS record points. |[optional] 
**Domain** | **string** | The root, or sending, domain. |
**Valid** | **bool** | Indicates if this is a valid Reverse DNS. |
**Legacy** | **bool** | Indicates if this Reverse DNS was created using the legacy whitelabel tool. If it is a legacy whitelabel, it will still function, but you'll need to create a new Reverse DNS if you need to update it. |
**LastValidationAttemptAt** | **int32** | A Unix epoch timestamp representing the last time of a validation attempt. |[optional] 
**ARecord** | [**ReverseDnsARecord**](ReverseDnsARecord.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


