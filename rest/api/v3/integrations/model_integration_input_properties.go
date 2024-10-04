/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Integrations API
* Integrations allows you to connect your SendGrid applications with third-party services and forward SendGrid email events to those external applications. Currently, Integrations supports event forwarding to [Segment](https://segment.com/docs). You can use this API to create, delete, view, and update your Integrations.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// IntegrationInputProperties The properties of an Integration required to send events to a specific third-party application.
type IntegrationInputProperties struct {
	// The write key provided by the Segment Source you'd like to have events forwarded to. Must be between 6 and 51 characters.
	WriteKey string `json:"write_key"`
	// The workspace region where the Segment Source write key lives.
	DestinationRegion DestinationRegion1 `json:"destination_region"`
}
