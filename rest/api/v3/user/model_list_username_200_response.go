/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid User API
* The Twilio SendGrid User API allows you to modify the settings of a SendGrid user account such as the user's email address or username. Keeping your user profile up to date helps SendGrid verify who you are and share important communications with you.  See [**Account Details**](https://docs.sendgrid.com/ui/account-and-settings/account) for more information. You can also manage your user settings in the [SendGrid application user interface](https://app.sendgrid.com/account/details).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ListUsername200Response struct for ListUsername200Response
type ListUsername200Response struct {
	// Your account username.
	Username string `json:"username"`
	// The user ID for your account.
	UserId int32 `json:"user_id"`
}