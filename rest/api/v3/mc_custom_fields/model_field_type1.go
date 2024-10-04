/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Custom Fields API
* The Twilio SendGrid Marketing Campaigns Custom Fields API allows you to add extra information about your marketing contacts that is relevant to your needs. For example, a fashion retailer might create a custom field for customers' shoe sizes, an ice cream shop might store customers' favorite flavors in a custom field, and you can create custom fields that make sense for your business.  With custom fields, you can also create [segments](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2/) based on custom fields values. Your custom fields are completely customizable to the use-cases and user information that you need.  You can also manage your Custom Fields using the SendGrid application user interface. See [**Using Custom Fields**](https://docs.sendgrid.com/ui/managing-contacts/custom-fields) for more information, including a list of Reserved Fields. You can also manage your custom fields in the [Marketing Campaigns application user interface](https://mc.sendgrid.com/custom-fields).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// FieldType1 the model 'FieldType1'
type FieldType1 string

// List of FieldType1
const (
	FIELDTYPE1_TEXT   FieldType1 = "Text"
	FIELDTYPE1_NUMBER FieldType1 = "Number"
	FIELDTYPE1_DATE   FieldType1 = "Date"
)
