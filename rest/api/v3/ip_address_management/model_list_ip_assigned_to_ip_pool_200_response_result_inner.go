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

// ListIpAssignedToIpPool200ResponseResultInner struct for ListIpAssignedToIpPool200ResponseResultInner
type ListIpAssignedToIpPool200ResponseResultInner struct {
	// An IP address assigned to the IP Pool.
	Ip *string `json:"ip,omitempty"`
	// The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request.
	Region *Region6 `json:"region,omitempty"`
	// IP Pools the IP address is assigned to.
	Pools *[]ListIp200ResponseResultInnerPoolsInner `json:"pools,omitempty"`
}