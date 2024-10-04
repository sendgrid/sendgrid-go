/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Scopes API
* The Twilio SendGrid Scopes API allows you to retrieve the scopes or permissions available to a user, see the user's attempts to access your SendGrid account, and, if necessary, deny an access request.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ListScope200Response struct for ListScope200Response
type ListScope200Response struct {
	// The list of scopes for which this user has access.
	Scopes []string `json:"scopes"`
}
