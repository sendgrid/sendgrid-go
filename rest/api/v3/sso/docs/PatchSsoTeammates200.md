# PatchSsoTeammates200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The Teammate's street address. |[optional] 
**Address2** | **string** | The Teammate's apartment number, suite number, or other secondary address information that is not part of the physical street address. |[optional] 
**City** | **string** | The Teammate's city. |[optional] 
**Company** | **string** | The Teammate's company name. |[optional] 
**Country** | **string** | The Teammate's country of residence. |[optional] 
**Username** | **string** | The Teammate's username. This property is set to the Teammate's email address. |[optional] 
**Phone** | **string** | The Teammate's phone number. |[optional] 
**State** | **string** | The Teammate's state or province. |[optional] 
**UserType** | [**UserType**](UserType.md) | A Teammate can be an `admin`, `owner`, or `teammate`. Each role is associated with the scope of the Teammate's permissions. |[optional] 
**Website** | **string** | A website associated with the Teammate. |[optional] 
**Zip** | **string** | The Teammate's zip code. |[optional] 
**FirstName** | **string** | The Teammate's first name. |[optional] 
**LastName** | **string** | The Teammate's last name. |[optional] 
**Email** | **string** | Teammate's email address. This email address also functions as the Teammate's username and must match the address assigned to the user in your IdP. This address cannot be changed after the Teammate is created. |[optional] 
**IsAdmin** | **bool** | Indicates if the Teammate has administrator permissions. When set to `true`, the Teammate is an admin. |[optional] 
**IsSso** | **bool** | Indicates how the Teammate authenticates with SendGrid. When set to `true`, the Teammate will access SendGrid via SSO and their IdP. When set to `false`, the Teammate will authenticate directly with SendGrid via a username and password. |[optional] 
**Scopes** | **[]string** | The permissions or scopes currently assigned to the Teammate. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of available scopes. |[optional] 
**HasRestrictedSubuserAccess** | **bool** | When this property is set to `true`, the Teammate has permissions to operate only on behalf of a Subuser. This property value is `true` when the `subuser_access` property is not empty. The `subuser_access` property determines which Subusers the Teammate may act on behalf of. |[optional] 
**SubuserAccess** | [**[]SsoTeammatesRestrictedSubuserResponsePropsSubuserAccessInner**](SsoTeammatesRestrictedSubuserResponsePropsSubuserAccessInner.md) | Specifies which Subusers the Teammate may access and act on behalf of. If this property is populated, the `has_restricted_subuser_access` property will be `true`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


