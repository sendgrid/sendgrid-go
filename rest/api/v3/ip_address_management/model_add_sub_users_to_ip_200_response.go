/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Address Management API
* The Twilio SendGrid IP Address Management API combines functionality that was previously split between the Twilio SendGrid [IP Address API](https://docs.sendgrid.com/api-reference/ip-address) and [IP Pools API](https://docs.sendgrid.com/api-reference/ip-pools). This functionality includes adding IP addresses to your account, assigning IP addresses to IP Pools and Subusers, among other tasks.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// AddSubUsersToIp200Response struct for AddSubUsersToIp200Response
type AddSubUsersToIp200Response struct {
	// The IP address the Subusers were assigned to.
	Ip *string `json:"ip,omitempty"`
	// An array of Subuser IDs that were assigned to the specified IP address.
	Subusers *[]string `json:"subusers,omitempty"`
}
