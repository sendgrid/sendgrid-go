/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid API Keys API
* The Twilio SendGrid API Keys API allows you manage your API keys and their settings. Your application, mail client, or website can all use API keys to authenticate access to SendGrid services.  To create your initial SendGrid API Key, you should use the [SendGrid application user interface](https://app.sendgrid.com/settings/api_keys). Once you have created a first key with scopes to manage additional API keys, you can use this API for all other key management.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// UpdateApiKeyNameRequest struct for UpdateApiKeyNameRequest
type UpdateApiKeyNameRequest struct {
	// The new name of the API Key.
	Name string `json:"name"`
}