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

// DeleteIpsFromIpPoolRequest struct for DeleteIpsFromIpPoolRequest
type DeleteIpsFromIpPoolRequest struct {
	// An array of IP addresses to remove from the specified IP Pool.
	Ips []string `json:"ips"`
}