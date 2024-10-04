# SsoTeammatesBaseRequestProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Set this property to the Teammate's first name. |
**LastName** | **string** | Set this property to the Teammate's last name. |
**IsAdmin** | **bool** | Set this property to `true` if the Teammate has admin permissions. You should not include the `scopes` or `persona` properties when setting the `is_admin` property to `true`—an admin will be allocated all scopes. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of scopes. |[optional] 
**Persona** | [**Persona**](Persona.md) | A persona represents a group of permissions often required by a type of Teammate such as a developer or marketer. Assigning a persona allows you to allocate a group of pre-defined permissions rather than assigning each scope individually. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a full list of the scopes assigned to each persona. |[optional] 
**Scopes** | **[]string** | Add or remove permissions from a Teammate using this `scopes` property. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of available scopes. You should not include this propety in the request when using the `persona` property or when setting the `is_admin` property to `true`—assigning a `persona` or setting `is_admin` to `true` will allocate a group of permissions to the Teammate. |[optional] 
**HasRestrictedSubuserAccess** | **bool** | Set this property to `true` to give the Teammate permissions to operate only on behalf of a Subuser. This property value must be `true` if the `subuser_access` property is not empty. The `subuser_access` property determines which Subusers the Teammate may act on behalf of. If this property is set to `true`, you cannot specify individual `scopes`, assign a `persona`, or set `is_admin` to `true`—a Teammate cannot specify scopes for the parent account and have restricted Subuser access. |[optional] 
**SubuserAccess** | [**[]SsoTeammatesBaseRequestPropsSubuserAccessInner**](SsoTeammatesBaseRequestPropsSubuserAccessInner.md) | Specify which Subusers the Teammate may access and act on behalf of with this property. If this property is populated, you must set the `has_restricted_subuser_access` property to `true`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


