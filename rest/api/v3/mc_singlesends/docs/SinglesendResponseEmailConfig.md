# SinglesendResponseEmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The subject line of the Single Send. This property is not used when a `design_id` value is set. |[optional] 
**HtmlContent** | **string** | The HTML content of the Single Send. This property is not used when a `design_id` value is set. |[optional] 
**PlainContent** | **string** | The plain text content of the Single Send. This property is not used when a `design_id` value is set. |[optional] 
**GeneratePlainContent** | **bool** | If this property is set to `true`, `plain_content` is always generated from `html_content`. If it's set to false, `plain_content` is not altered. |[optional] [default to true]
**DesignId** | **string** | A `design_id` can be used in place of `html_content`, `plain_content`, and/or `subject`. You can retrieve a design's ID from the [**List Designs** endpoint](https://docs.sendgrid.com/api-reference/designs-api/list-designs) or by pulling it from the design's detail page URL in the Marketing Campaigns App. |[optional] 
**Editor** | [**Editor1**](Editor1.md) | The editor, `design` or `code`, used to modify the Single Send's design in the Marketing Campaigns application user interface. |[optional] 
**SuppressionGroupId** | **int32** | The ID of the Suppression Group to allow recipients to unsubscribe. You must provide a `suppression_group_id` or the `custom_unsubscribe_url` with your Single Send. |[optional] 
**CustomUnsubscribeUrl** | **string** | The URL allowing recipients to unsubscribe. You must provide a `custom_unsubscribe_url` or a `suppression_group_id` with your Single Send. |[optional] 
**SenderId** | **int32** | The ID of the verified sender from whom the Single Send will be sent. You can retrieve a verified sender's ID from the [**Get Verified Senders** endpoint](https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/get-all-verified-senders) or by pulling it from the sender's detail page URL in the SendGrid App. |[optional] 
**IpPool** | **string** | The name of the IP Pool from which the Single Send emails are sent. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


