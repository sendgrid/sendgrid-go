/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Senders API
* The Twilio SendGrid Marketing Campaigns Senders API allows you to create a verified sender from which your marketing emails will be sent. To ensure our customers maintain the best possible sender reputations and to uphold legitimate sending behavior, we require customers to verify their Senders. A Sender represents your “From” email address—the address your recipients will see as the sender of your emails.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// Sender struct for Sender
type Sender struct {
	// The unique identifier of the Sender.
	Id int32 `json:"id"`
	// A nickname for the Sender. Not used for sending.
	Nickname string                     `json:"nickname"`
	From     CreateSenderRequestFrom    `json:"from"`
	ReplyTo  CreateSenderRequestReplyTo `json:"reply_to"`
	// The physical address of the Sender.
	Address string `json:"address"`
	// Additional Sender address information.
	Address2 *string `json:"address_2,omitempty"`
	// The city of the Sender.
	City string `json:"city"`
	// The state of the Sender.
	State *string `json:"state,omitempty"`
	// The zipcode of the Sender.
	Zip *string `json:"zip,omitempty"`
	// The country of the Sender.
	Country string `json:"country"`
	// A boolean flag indicating whether the Sender is verified or not. Only verified Senders can be used to send email.
	Verified bool `json:"verified"`
	// A boolean flag that is `true` when the Sender is associated with a campaign in Draft, Scheduled, or In Progress status. You cannot update or delete a locked Sender.
	Locked bool `json:"locked"`
	// The time the Sender was last updated.
	UpdatedAt int32 `json:"updated_at"`
	// The time the Sender was created.
	CreatedAt int32 `json:"created_at"`
}