# SinglesendRequestEmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The subject line of the Single Send. Do not include this field when using a `design_id`. |[optional] 
**HtmlContent** | **string** | The HTML content of the Single Send. Do not include this field when using a `design_id`. |[optional] 
**PlainContent** | **string** | The plain text content of the Single Send. Do not include this field when using a `design_id`. |[optional] 
**GeneratePlainContent** | **bool** | If set to `true`, `plain_content` is always generated from `html_content`. If set to false, `plain_content` is not altered. |[optional] [default to true]
**DesignId** | **string** | A `design_id` can be used in place of `html_content`, `plain_content`, and/or `subject`. You can retrieve a design's ID from the [\"List Designs\" endpoint](https://docs.sendgrid.com/api-reference/designs-api/list-designs) or by pulling it from the design's detail page URL in the Marketing Campaigns App. |[optional] 
**Editor** | [**Editor**](Editor.md) | The editor — `\"design\"` or `\"code\"` — used to modify the Single Send's design in the Marketing Campaigns App. |[optional] 
**SuppressionGroupId** | **int32** | The ID of the Suppression Group to allow recipients to unsubscribe — you must provide this or the `custom_unsubscribe_url`. |[optional] 
**CustomUnsubscribeUrl** | **string** | The URL allowing recipients to unsubscribe — you must provide this or the `suppression_group_id`. |[optional] 
**SenderId** | **int32** | The ID of the verified Sender. You can retrieve a verified Sender's ID from the [\"Get Verified Senders\" endpoint](https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/get-all-verified-senders) or by pulling it from the Sender's detail page URL in the SendGrid App. |[optional] 
**IpPool** | **string** | The name of the IP Pool from which the Single Send emails are sent. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


