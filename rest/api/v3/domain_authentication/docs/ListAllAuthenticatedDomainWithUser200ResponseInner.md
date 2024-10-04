# ListAllAuthenticatedDomainWithUser200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the authenticated domain. |
**UserId** | **float32** | The ID of the user that this domain is associated with. |
**Subdomain** | **string** | The subdomain to use for this authenticated domain. |
**Domain** | **string** | The domain to be authenticated. |
**Username** | **string** | The username that this domain will be associated with. |
**Ips** | **[]string** | The IPs to be included in the custom SPF record for this authenticated domain. |
**CustomSpf** | **bool** | Indicates whether this authenticated domain uses custom SPF. |
**Default** | **bool** | Indicates if this is the default authenticated domain. |
**Legacy** | **bool** | Indicates if this authenticated domain was created using the legacy whitelabel tool. If it is a legacy whitelabel, it will still function, but you'll need to create a new authenticated domain if you need to update it. |
**AutomaticSecurity** | **bool** | Indicates if this authenticated domain uses automated security. |
**Valid** | **bool** | Indicates if this is a valid authenticated domain. |
**Dns** | [**ListAllAuthenticatedDomainWithUser200ResponseInnerDns**](ListAllAuthenticatedDomainWithUser200ResponseInnerDns.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


