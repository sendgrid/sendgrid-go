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

// GetIpPool200ResponseIpCountByRegionInner struct for GetIpPool200ResponseIpCountByRegionInner
type GetIpPool200ResponseIpCountByRegionInner struct {
	// The region associated with the number of IPs listed in the `count` property.
	Region *Region5 `json:"region,omitempty"`
	// The number of IP addresses in the pool that are assigned to the region specified in the `region` property for this count.
	Count *int32 `json:"count,omitempty"`
}