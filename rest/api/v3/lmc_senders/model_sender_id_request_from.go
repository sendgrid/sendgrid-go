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

// SenderIdRequestFrom struct for SenderIdRequestFrom
type SenderIdRequestFrom struct {
	// The email address from which your recipient will receive emails.
	Email *string `json:"email,omitempty"`
	// The name appended to the from email field. Typically your name or company name.
	Name *string `json:"name,omitempty"`
}
