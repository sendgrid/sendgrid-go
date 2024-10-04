# AuthenticatedDomainSpf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the authenticated domain. |
**Domain** | **string** | The domain authenticated. |
**Subdomain** | **string** | The subdomain that was used to create this authenticated domain. |[optional] 
**Username** | **string** | The username of the account that this authenticated domain is associated with. |
**UserId** | **int32** | The user_id of the account that this authenticated domain is associated with. |
**Ips** | **[]interface{}** | The IP addresses that are included in the SPF record for this authenticated domain. |
**CustomSpf** | **bool** | Indicates if this authenticated domain uses custom SPF. |
**Default** | **bool** | Indicates if this is the default domain. |
**Legacy** | **bool** | Indicates if this authenticated domain was created using the legacy whitelabel tool. If it is a legacy whitelabel, it will still function, but you'll need to create a new authenticated domain if you need to update it. |
**AutomaticSecurity** | **bool** | Indicates if this authenticated domain uses automated security. |
**Valid** | **bool** | Indicates if this is a valid authenticated domain . |
**Dns** | [**AuthenticatedDomainSpfDns**](AuthenticatedDomainSpfDns.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


