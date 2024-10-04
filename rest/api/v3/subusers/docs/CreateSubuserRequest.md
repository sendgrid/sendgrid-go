# CreateSubuserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username for this subuser. |
**Email** | **string** | The email address of the subuser. |
**Password** | **string** | The password this subuser will use when logging into SendGrid. |
**Ips** | **[]string** | The IP addresses that should be assigned to this subuser. |
**Region** | [**Region1**](Region1.md) | The region this Subuser should be assigned to. Can be `global` or `eu`. (Regional email is in Public Beta and requires SendGrid Pro plan or above.). |[optional] 
**IncludeRegion** | **bool** | A flag that determines if the Subuser's region should be returned in the response. (Regional email is in Public Beta and requires SendGrid Pro plan or above.) |[optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


