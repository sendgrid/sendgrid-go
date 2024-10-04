# SsoTeammatesBaseRequestPropsSubuserAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Set this property to the ID of a Subuser to which the Teammate should have access. You can retrieve Subuser IDs from the [Subusers API](https://docs.sendgrid.com/api-reference/subusers-api/). |
**PermissionType** | [**PermissionType**](PermissionType.md) | Grant the level of access the Teammate should have to the specified Subuser with this property. This property value may be either `admin` or `restricted`. When set to `restricted`, the Teammate has only the permissions assigned in the `scopes` property. |
**Scopes** | **[]string** | Add or remove permissions that the Teammate can access on behalf of the Subuser. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of available scopes. You should not include this property in the request when the `permission_type` property is set to `admin`—administrators have full access to the specified Subuser. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


