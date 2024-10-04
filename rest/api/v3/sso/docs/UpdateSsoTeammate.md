# UpdateSsoTeammate

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSsoTeammate**](UpdateSsoTeammate.md#UpdateSsoTeammate) | **Patch** /v3/sso/teammates/{Username} | Edit an SSO Teammate



## UpdateSsoTeammate

> PatchSsoTeammates200 UpdateSsoTeammate(ctx, Usernameoptional)

Edit an SSO Teammate

**This endpoint allows you to modify an existing SSO Teammate.**  Only the parent user and Teammates with admin permissions can update another Teammate's permissions.  ### Scopes  When updating a Teammate, you will assign it permissions or scopes. These scopes determine which actions the Teammate can perform and which features they can access. Scopes are provided with one of three properties passed to this endpoint: `is_admin`, `scopes`, and `persona`.  You can make a Teammate an administrator by setting `is_admin` to `true`. Administrators will have all scopes assigned to them. Alternatively, you can assign a `persona` to the teammate, which will assign them a block of permissions commonly required for that type of user. See the \"Persona scopes\" section of [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions#persona-scopes) for a list of permsissions granted by persona. Lastly, you can assign individual permissions with the `scopes` property. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a full list of scopes that can be assigned to a Teammate.  ### Subuser access  SendGrid Teammates may be assigned access to one or more Subusers. Subusers function like SendGrid sub-accounts with their own resources. See [**Subusers**](https://docs.sendgrid.com/ui/account-and-settings/subusers) for more information.  When assigning Subuser access to a Teammate, you may set the `has_restricted_subuser_access` property to `true` to constrain the Teammate so that they can operate only on behalf of the Subusers to which they are assigned. You may further set the level of access the Teammate has to each Subuser with the `subuser_access` property.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Username** | **string** | Set this parameter to the Teammate's email address. This address must be the same address assigned to the Teammate in your IdP.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSsoTeammateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **SsoTeammatesBaseRequestProps** | 

### Return type

[**PatchSsoTeammates200**](PatchSsoTeammates200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

