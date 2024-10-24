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

// SsoTeammatesRestrictedSubuserResponseProps struct for SsoTeammatesRestrictedSubuserResponseProps
type SsoTeammatesRestrictedSubuserResponseProps struct {
	// When this property is set to `true`, the Teammate has permissions to operate only on behalf of a Subuser. This property value is `true` when the `subuser_access` property is not empty. The `subuser_access` property determines which Subusers the Teammate may act on behalf of.
	HasRestrictedSubuserAccess *bool `json:"has_restricted_subuser_access,omitempty"`
	// Specifies which Subusers the Teammate may access and act on behalf of. If this property is populated, the `has_restricted_subuser_access` property will be `true`.
	SubuserAccess *[]SsoTeammatesRestrictedSubuserResponsePropsSubuserAccessInner `json:"subuser_access,omitempty"`
}