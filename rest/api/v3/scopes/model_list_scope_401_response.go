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

// ListScope401Response struct for ListScope401Response
type ListScope401Response struct {
	// This 401 response indicates that the user making the call doesn't have the authorization to view the list of scopes.
	Errors []ListScope401ResponseErrorsInner `json:"errors"`
}