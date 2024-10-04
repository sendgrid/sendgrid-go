# AuthenticateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain being authenticated. |
**Subdomain** | **string** | The subdomain to use for this authenticated domain. |[optional] 
**Username** | **string** | The username associated with this domain. |[optional] 
**Ips** | **[]string** | The IP addresses that will be included in the custom SPF record for this authenticated domain. |[optional] 
**CustomSpf** | **bool** | Specify whether to use a custom SPF or allow SendGrid to manage your SPF. This option is only available to authenticated domains set up for manual security. |[optional] 
**Default** | **bool** | Whether to use this authenticated domain as the fallback if no authenticated domains match the sender's domain. |[optional] 
**AutomaticSecurity** | **bool** | Whether to allow SendGrid to manage your SPF records, DKIM keys, and DKIM key rotation. |[optional] 
**CustomDkimSelector** | **string** | Add a custom DKIM selector. Accepts three letters or numbers. |[optional] 
**Region** | **string** | The region of the domain. Allowed values are `global` and `eu`. The default value is `global`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


