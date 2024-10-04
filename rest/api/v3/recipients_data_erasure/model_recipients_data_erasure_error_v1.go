/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Recipients' Data Erasure API
* The Recipients' Data Erasure API allows Twilio SendGrid customers to delete their own customers' personal data from the Twilio SendGrid Platform. The use of this API facilitates the self-service removal of email personal data from the Twilio SendGrid platform and will enable customers to comply with various obligatory data privacy regulations.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// RecipientsDataErasureErrorV1 struct for RecipientsDataErasureErrorV1
type RecipientsDataErasureErrorV1 struct {
	// The message representing the error from the API
	Message string `json:"message"`
	// The field associated with the error
	Field string `json:"field"`
}
