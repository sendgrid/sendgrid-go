# LinkBranding200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the branded link. |
**Domain** | **string** | The root domain of the branded link. |
**Subdomain** | **string** | The subdomain used to generate the DNS records for this link branding. This subdomain must be different from the subdomain used for your authenticated domain. |[optional] 
**Username** | **string** | The username of the account that this link branding is associated with. |
**UserId** | **int32** | The ID of the user that this link branding is associated with. |
**Default** | [**Default2**](Default2.md) | Indicates if this is the default link branding. |
**Valid** | [**Valid3**](Valid3.md) | Indicates if this link branding is valid. |
**Legacy** | [**Legacy**](Legacy.md) | Indicates if this link branding was created using the legacy whitelabel tool. If it is a legacy whitelabel, it will still function, but you'll need to create new link branding if you need to update it. |
**Dns** | [**LinkBranding200Dns**](LinkBranding200Dns.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


