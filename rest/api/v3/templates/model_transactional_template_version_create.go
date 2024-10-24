/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Templates API
* The Twilio SendGrid Templates API allows you to create and manage email templates to be delivered with SendGrid's sending APIs. The templates you create will be available using a template ID that is passed to our sending APIs as part of the request. Each template may then have multiple versions associated with it. Whichever version is set as \"active\" at the time of the request will be sent to your recipients. This system allows you to update a single template's look and feel entirely without modifying your requests to our Mail Send API. For example, you could have a single template for welcome emails. That welcome template could then have a version for each season of the year that's themed appropriately and marked as active during the appropriate season. The template ID passed to our sending APIs never needs to change; you can just modify which version is active.  This API provides operations to create and manage your templates as well as their versions.  Each user can create up to 300 different templates. Templates are specific to accounts and Subusers. Templates created on a parent account will not be accessible from the Subusers' accounts.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// TransactionalTemplateVersionCreate struct for TransactionalTemplateVersionCreate
type TransactionalTemplateVersionCreate struct {
	// Set the version as the active version associated with the template (0 is inactive, 1 is active). Only one version of a template can be active. The first version created for a template will automatically be set to Active.
	Active *Active `json:"active,omitempty"`
	// Name of the transactional template version.
	Name string `json:"name"`
	// The HTML content of the version. Maximum of 1048576 bytes allowed.
	HtmlContent *string `json:"html_content,omitempty"`
	// Text/plain content of the transactional template version. Maximum of 1048576 bytes allowed.
	PlainContent *string `json:"plain_content,omitempty"`
	// If true, plain_content is always generated from html_content. If false, plain_content is not altered.
	GeneratePlainContent *bool `json:"generate_plain_content,omitempty"`
	// Subject of the new transactional template version.
	Subject string `json:"subject"`
	// The editor used in the UI.
	Editor *Editor `json:"editor,omitempty"`
	// For dynamic templates only, the mock json data that will be used for template preview and test sends.
	TestData *string `json:"test_data,omitempty"`
}