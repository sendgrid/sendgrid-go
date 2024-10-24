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

// ListIp200ResponseResultInner struct for ListIp200ResponseResultInner
type ListIp200ResponseResultInner struct {
	// An IP address on your account.
	Ip *string `json:"ip,omitempty"`
	// An array of IP Pools the IP address is assigned to.
	Pools *[]ListIp200ResponseResultInnerPoolsInner `json:"pools,omitempty"`
	// Indicates if the IP address is set to automatically [warmup](https://docs.sendgrid.com/ui/sending-email/warming-up-an-ip-address).
	IsAutoWarmup *bool `json:"is_auto_warmup,omitempty"`
	// Indicates if a parent on the account is able to send email from the IP address. This parameter will be returned only if the request was made by the parent account.
	IsParentAssigned *bool `json:"is_parent_assigned,omitempty"`
	// A timestamp indicating when the IP was last updated.
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// Indicates if the IP address is billed and able to send email. This parameter applies to non-Twilio SendGrid APIs that been added to your Twilio SendGrid account. This parameter's value is `null` for Twilio SendGrid IP addresses.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Indicates whether an IP address is leased from Twilio SendGrid. If `false`, the IP address is not a Twilio SendGrid IP; it is a customer's own IP that has been added to their Twilio SendGrid account.
	IsLeased *bool `json:"is_leased,omitempty"`
	// A timestamp representing when the IP address was added to your account.
	AddedAt *int32 `json:"added_at,omitempty"`
	// The region to which the IP address is assigned. This property will only be returned if the `include_region` query parameter is included and set to `true` as part of the API request.
	Region *Region `json:"region,omitempty"`
}