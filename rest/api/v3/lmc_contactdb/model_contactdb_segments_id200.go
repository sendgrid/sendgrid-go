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

// ContactdbSegmentsId200 struct for ContactdbSegmentsId200
type ContactdbSegmentsId200 struct {
	// The ID of the segment.
	Id float32 `json:"id"`
	// The name of this segment.
	Name string `json:"name"`
	// The list id from which to make this segment. Not including this ID will mean your segment is created from the main contactdb rather than a list.
	ListId *int32 `json:"list_id,omitempty"`
	// The conditions for a recipient to be included in this segment.
	Conditions []ContactdbSegmentsConditions `json:"conditions"`
	// The count of recipients in this list. This is not included on creation of segments.
	RecipientCount *float32 `json:"recipient_count,omitempty"`
}
