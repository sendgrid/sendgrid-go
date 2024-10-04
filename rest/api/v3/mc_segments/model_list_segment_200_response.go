/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Segments API
* This API was deprecated on December 31, 2022. Following deprecation, all segments created in the Marketing Campaigns user interface began using the [Segmentation v2 API](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2).  To enable manual migration and data retrieval, this API's GET and DELETE operations will remain available. The POST (create) and PATCH (update) endpoints were removed on January 31, 2023 because it is no longer possible to create new v1 segments or modify existing ones. See our [Segmentation v1 to v2 upgrade instructions](https://docs.sendgrid.com/for-developers/sending-email/getting-started-the-marketing-campaigns-v2-segmentation-api#upgrade-a-v1-segment-to-v2) to manually migrate your segments to the v2 API.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ListSegment200Response struct for ListSegment200Response
type ListSegment200Response struct {
	Results []SegmentSummary `json:"results"`
}
