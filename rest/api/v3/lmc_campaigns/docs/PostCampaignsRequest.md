# PostCampaignsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The display title of your campaign. This will be viewable by you in the Marketing Campaigns UI. |
**Subject** | **string** | The subject of your campaign that your recipients will see. |[optional] 
**SenderId** | **int32** | The ID of the \"sender\" identity that you have created. Your recipients will see this as the \"from\" on your marketing emails. |[optional] 
**ListIds** | **[]int32** | The IDs of the lists you are sending this campaign to. You can have both segment IDs and list IDs |[optional] 
**SegmentIds** | **[]int32** | The segment IDs that you are sending this list to. You can have both segment IDs and list IDs. Segments are limited to 10 segment IDs. |[optional] 
**Categories** | **[]string** | The categories you would like associated to this campaign. |[optional] 
**SuppressionGroupId** | **int32** | The suppression group that this marketing email belongs to, allowing recipients to opt-out of emails of this type. |[optional] 
**CustomUnsubscribeUrl** | **string** | This is the url of the custom unsubscribe page that you provide for customers to unsubscribe from your suppression groups. |[optional] 
**IpPool** | **string** | The pool of IPs that you would like to send this email from. |[optional] 
**HtmlContent** | **string** | The HTML of your marketing email. |[optional] 
**PlainContent** | **string** | The plain text content of your emails. |[optional] 
**Editor** | [**Editor**](Editor.md) | The editor used in the UI. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


