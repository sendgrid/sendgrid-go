/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Segments 2.0 API
* The Twilio SendGrid Marketing Campaigns Segments V2 API allows you to create, edit, and delete segments as well as retrieve a list of segments or an individual segment by ID.  Segments are similar to contact lists, except they update dynamically over time as information stored about your contacts or the criteria used to define your segments changes. When you segment your audience, you are able to create personalized Automation emails and Single Sends that directly address the wants and needs of your particular audience.  Note that Twilio SendGrid checks for newly added or modified contacts who meet a segment's criteria on an hourly schedule. Only existing contacts who meet a segment's criteria will be included in the segment searches within 15 minutes.  Segments built using engagement data such as \"was sent\" or \"clicked\" will take approximately 30 minutes to begin populating.  Segment samples and counts are refreshed approximately once per hour; they do not update immediately. If no contacts are added to or removed from a segment since the last refresh, the sample and UI count displayed will be refreshed at increasing time intervals with a maximum sample and count refresh delay of 24 hours.  You can also manage your Segments with the [Marketing Campaigns application user interface](https://mc.sendgrid.com/contacts). See [**Segmenting Your Contacts**](https://docs.sendgrid.com/ui/managing-contacts/segmenting-your-contacts) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// SegmentError struct for SegmentError
type SegmentError struct {
	// A description of the error.
	Error string `json:"error"`
}
