/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Legacy Marketing Campaigns Sender Identities API
* The Twilio SendGrid Legacy Marketing Campaigns Sender Identites API allows you to manage the email addresses used to send your marketing email. This API is operational, but we recommend using the current version of Marketing Campaigns to manage your [senders](https://docs.sendgrid.com/api-reference/senders/).  See [**Migrating from Legacy Marketing Campaigns**](https://docs.sendgrid.com/ui/sending-email/migrating-from-legacy-marketing-campaigns) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// SenderIdRequest struct for SenderIdRequest
type SenderIdRequest struct {
	// A nickname for the sender identity. Not used for sending.
	Nickname *string                 `json:"nickname,omitempty"`
	From     *SenderIdRequestFrom    `json:"from,omitempty"`
	ReplyTo  *SenderIdRequestReplyTo `json:"reply_to,omitempty"`
	// The physical address of the sender identity.
	Address *string `json:"address,omitempty"`
	// Additional sender identity address information.
	Address2 *string `json:"address_2,omitempty"`
	// The city of the sender identity.
	City *string `json:"city,omitempty"`
	// The state of the sender identity.
	State *string `json:"state,omitempty"`
	// The zipcode of the sender identity.
	Zip *string `json:"zip,omitempty"`
	// The country of the sender identity.
	Country *string `json:"country,omitempty"`
}
