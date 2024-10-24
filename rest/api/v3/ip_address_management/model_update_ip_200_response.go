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

// UpdateIp200Response struct for UpdateIp200Response
type UpdateIp200Response struct {
	// The IP address that was updated.
	Ip *string `json:"ip,omitempty"`
	// Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address).
	IsAutoWarmup *bool `json:"is_auto_warmup,omitempty"`
	// Indicates if a parent on the account is able to send email from the IP address. This parameter is returned only if the IP address is parent assigned.
	IsParentAssigned *bool `json:"is_parent_assigned,omitempty"`
	// An array of Subuser IDs the IP address was assigned to. This parameter is returned only if the IP address is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
}