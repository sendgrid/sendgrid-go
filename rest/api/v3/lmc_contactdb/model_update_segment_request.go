/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Legacy Marketing Campaigns Contacts API
* The Twilio SendGrid Legacy Marketing Campaigns Contacts API allows you to manage your marketing contacts programmatically. This API is operational, but we recommend using the current version of Marketing Campaigns' [Contacts API](https://docs.sendgrid.com/api-reference/contacts/), [Lists API](https://docs.sendgrid.com/api-reference/lists/), and [Segments API](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2/) to manage your contacts.  See [**Migrating from Legacy Marketing Campaigns**](https://docs.sendgrid.com/ui/sending-email/migrating-from-legacy-marketing-campaigns) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// UpdateSegmentRequest struct for UpdateSegmentRequest
type UpdateSegmentRequest struct {
	Name string `json:"name"`
	// The list ID you would like this segment to be built from.
	ListId *float32 `json:"list_id,omitempty"`
	// The conditions by which this segment should be created.
	Conditions *[]ContactdbSegmentsConditions `json:"conditions,omitempty"`
}