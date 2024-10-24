/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Teammates API
* The Twilio SendGrid Teammates API allows you to add, manage, and remove Teammates, or user accounts, from your SendGrid account. Teammates function like user accounts on the SendGrid account, allowing you to invite additional users to your account with scoped access. You can think of Teammates as SendGrid's approach to enabling [role-based access control](https://en.wikipedia.org/wiki/Role-based_access_control) for your SendGrid account. For even more control over the access to your account, see the [Twilio SendGrid SSO API](https://docs.sendgrid.com/api-reference/single-sign-on-teammates/), which enables SSO-authenticated Teammates to be managed through a SAML 2.0 identity provider.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// GetTeammate200Response struct for GetTeammate200Response
type GetTeammate200Response struct {
	// Teammate's username
	Username *string `json:"username,omitempty"`
	// Teammate's first name
	FirstName *string `json:"first_name,omitempty"`
	// Teammate's last name
	LastName *string `json:"last_name,omitempty"`
	// Teammate's email
	Email *string `json:"email,omitempty"`
	// Scopes associated to teammate
	Scopes *[]interface{} `json:"scopes,omitempty"`
	// Indicate the type of user: account owner, teammate admin user, or normal teammate
	UserType *UserType1 `json:"user_type,omitempty"`
	// Set to true if teammate has admin privileges
	IsAdmin *bool `json:"is_admin,omitempty"`
	// (optional) Teammate's phone number
	Phone *string `json:"phone,omitempty"`
	// (optional) Teammate's website
	Website *string `json:"website,omitempty"`
	// (optional) Teammate's address
	Address *string `json:"address,omitempty"`
	// (optional) Teammate's address
	Address2 *string `json:"address2,omitempty"`
	// (optional) Teammate's city
	City *string `json:"city,omitempty"`
	// (optional) Teammate's state
	State *string `json:"state,omitempty"`
	// (optional) Teammate's zip
	Zip *string `json:"zip,omitempty"`
	// (optional) Teammate's country
	Country *string `json:"country,omitempty"`
}