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

// ListIpAssignedToIpPool200Response struct for ListIpAssignedToIpPool200Response
type ListIpAssignedToIpPool200Response struct {
	Result   *[]ListIpAssignedToIpPool200ResponseResultInner `json:"result,omitempty"`
	Metadata *ListIpAssignedToIpPool200ResponseMetadata      `json:"_metadata,omitempty"`
}