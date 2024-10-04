# SsoTeammatesRestrictedSubuserResponsePropsSubuserAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of a Subuser to which the Teammate has access. You can retrieve Subuser IDs from the [Subusers API](https://docs.sendgrid.com/api-reference/subusers-api/). |[optional] 
**Username** | **string** | The username of a Subuser to which the Teammate has access. |[optional] 
**Email** | **string** | The email address of a Subuser to which the Teammate has access. |[optional] 
**Disabled** | **bool** | Indicates if the Subuser is active for the SendGrid account. |[optional] 
**PermissionType** | [**PermissionType1**](PermissionType1.md) | The level of access the Teammate has to the specified Subuser. This property value may be either `admin` or `restricted`. When is property is set to `restricted`, the Teammate has only the permissions assigned in the `scopes` property. |[optional] 
**Scopes** | **[]string** | The permissions or scopes that the Teammate can access on behalf of the Subuser. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of available scopes. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


