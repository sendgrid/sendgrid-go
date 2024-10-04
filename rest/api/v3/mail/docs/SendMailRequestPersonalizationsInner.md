# SendMailRequestPersonalizationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**MailFrom**](MailFrom.md) |  |[optional] 
**To** | [**[]MailTo**](MailTo.md) | An array of recipients to whom the email will be sent. Each object in this array must contain a recipient's email address. Each object in the array may optionally contain a recipient's name. |
**Cc** | [**[]MailTo**](MailTo.md) | An array of recipients to whom a copy of the email will be sent. Each object in this array must contain a recipient's email address. Each object in the array may optionally contain a recipient's name. |[optional] 
**Bcc** | [**[]MailTo**](MailTo.md) | An array of recipients to whom a blind carbon copy of your email will be sent. Each object in this array must contain a recipient's email address. Each object in the array may optionally contain a recipient's name. |[optional] 
**Subject** | **string** | The subject of your email. See line length limits specified in [RFC 2822](https://www.rfc-editor.org/rfc/rfc2822#section-2.1.1) for guidance on subject line character limits. |[optional] 
**Headers** | **map[string]interface{}** | A collection of JSON property name and property value pairs allowing you to specify handling instructions for your email. You may not override the following headers: `x-sg-id`, `x-sg-eid`, `received`, `dkim-signature`, `Content-Type`, `Content-Transfer-Encoding`, `To`, `From`, `Subject`, `Reply-To`, `CC`, `BCC`. |[optional] 
**Substitutions** | **map[string]string** | A collection property names that will be substituted by their corresponding property values in the `subject`, `reply-to` and `content` portions of your message. Name and value pairs follow the pattern `\"substitution_tag\":\"value to substitute\"`. The total collective size of your substitutions may not exceed 10,000 bytes per personalization object. Substitutions allow you to insert data without using SendGrid's Dynamic Templates. This property should not be used in combination with a Dynamic Template, which can be identified by a `template_id` starting with `d-`. See [**Substitution Tags**](https://docs.sendgrid.com/for-developers/sending-email/substitution-tags) for more information. |[optional] 
**DynamicTemplateData** | **map[string]interface{}** | A collection property names that will be substituted by their corresponding property values in the `subject`, `reply-to` and `content` portions of a SendGrid Dynamic Template. Dynamic template data is available in a template using [Handlebars](https://docs.sendgrid.com/for-developers/sending-email/using-handlebars) syntax. This property should be used in combination with a Dynamic Template, which can be identified by a `template_id` starting with `d-`. See [**How to Send an Email with Dynamic Templates**](https://docs.sendgrid.com/ui/sending-email/how-to-send-an-email-with-dynamic-templates) for more information. |[optional] 
**CustomArgs** | **map[string]interface{}** | Values that are specific to this personalization that will be carried along with the email and its activity data. Substitutions will not be made on custom arguments, so any string that is assigned to this property will be assumed to be the custom argument that you would like to be used. This field may not exceed 10,000 bytes. |[optional] 
**SendAt** | **int32** | A [unix timestamp](https://en.wikipedia.org/wiki/Unix_time) allowing you to specify when your email should be sent. A send cannot be scheduled more than 72 hours in advance. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


