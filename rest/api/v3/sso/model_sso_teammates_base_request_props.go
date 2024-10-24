/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Single Sign-On API
* The Single Sign-On API allows you to manage your SAML 2.0 SSO configurations. You can also work with your SSO integrations using the SSO section of the [Twilio SendGrid application user interface](https://app.sendgrid.com/settings/sso).  The Single Sign-On Settings operations allow you to create, retrieve, modify, and delete SSO integrations for your Twilio SendGrid account. Each integration will correspond to a specific IdP such as Okta, Duo, or Microsoft Azure Active Directory.  The Single Sign-On Certificates operations allow you to create, modify, and delete SSO certificates. A SAML certificate allows your IdP and Twilio SendGrid to verify requests are coming from one another using the `public_certificate` and `integration_id` parameters.  The Single Sign-On Teammates operations allow you to add and modify SSO Teammates. SSO Teammates are the individual user accounts who will access your Twilio SendGrid account with SSO credentials. To retrieve or delete an SSO Teammate, you will use the separate [Teammates API](https://docs.sendgrid.com/api-reference/teammates/).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// SsoTeammatesBaseRequestProps struct for SsoTeammatesBaseRequestProps
type SsoTeammatesBaseRequestProps struct {
	// Set this property to the Teammate's first name.
	FirstName string `json:"first_name"`
	// Set this property to the Teammate's last name.
	LastName string `json:"last_name"`
	// Set this property to `true` if the Teammate has admin permissions. You should not include the `scopes` or `persona` properties when setting the `is_admin` property to `true`—an admin will be allocated all scopes. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of scopes.
	IsAdmin *bool `json:"is_admin,omitempty"`
	// A persona represents a group of permissions often required by a type of Teammate such as a developer or marketer. Assigning a persona allows you to allocate a group of pre-defined permissions rather than assigning each scope individually. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a full list of the scopes assigned to each persona.
	Persona *Persona `json:"persona,omitempty"`
	// Add or remove permissions from a Teammate using this `scopes` property. See [**Teammate Permissions**](https://docs.sendgrid.com/ui/account-and-settings/teammate-permissions) for a complete list of available scopes. You should not include this propety in the request when using the `persona` property or when setting the `is_admin` property to `true`—assigning a `persona` or setting `is_admin` to `true` will allocate a group of permissions to the Teammate.
	Scopes *[]string `json:"scopes,omitempty"`
	// Set this property to `true` to give the Teammate permissions to operate only on behalf of a Subuser. This property value must be `true` if the `subuser_access` property is not empty. The `subuser_access` property determines which Subusers the Teammate may act on behalf of. If this property is set to `true`, you cannot specify individual `scopes`, assign a `persona`, or set `is_admin` to `true`—a Teammate cannot specify scopes for the parent account and have restricted Subuser access.
	HasRestrictedSubuserAccess *bool `json:"has_restricted_subuser_access,omitempty"`
	// Specify which Subusers the Teammate may access and act on behalf of with this property. If this property is populated, you must set the `has_restricted_subuser_access` property to `true`.
	SubuserAccess *[]SsoTeammatesBaseRequestPropsSubuserAccessInner `json:"subuser_access,omitempty"`
}