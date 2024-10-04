/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Mail API
* The Twilio SendGrid v3 Mail API allows you to send email at scale over HTTP. The Mail Send endpoint supports many levels of functionality, allowing you to send templates, set categories and custom arguments that can be used to analyze your send, and configure which tracking settings to include such as opens and clicks. You can also group mail sends into batches, allowing you to schedule and cancel sends by their batch IDs.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// SendMailRequestAsm An object allowing you to specify how to handle unsubscribes. With SendGrid, an unsubscribe is the action an email recipient takes when they opt-out of receiving your messages. A suppression is the action you take as a sender to filter or suppress an unsubscribed address from your email send. From the perspective of the recipient, your suppression is the result of their unsubscribe. See [**Suppression Groups**](https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/create-a-new-suppression-group) for more information.
type SendMailRequestAsm struct {
	// The unsubscribe group to associate with this email. See the [Suppressions API](https://docs.sendgrid.com/api-reference/suppressions/) to manage unsubscribe group IDs.
	GroupId int32 `json:"group_id"`
	// An array containing the unsubscribe groups that you would like to be displayed to a recipient on the unsubscribe preferences page. This page is displayed in the recipient's browser when they click the unsubscribe link in your message.
	GroupsToDisplay *[]int32 `json:"groups_to_display,omitempty"`
}
