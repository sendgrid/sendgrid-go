# SinglesendResponseShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Name** | **string** | name of the Single Send |
**Abtest** | [**AbTestSummary**](AbTestSummary.md) |  |
**Status** | [**Status3**](Status3.md) | current status of the Single Send |
**Categories** | **[]string** | categories to associate with this Single Send |
**SendAt** | [**time.Time**](time.Time.md) | The ISO 8601 time at which to send the Single Send. This must be in future or the string `now`. SendGrid [Mail Send](https://docs.sendgrid.com/api-reference/mail-send/mail-send) emails can be scheduled up to 72 hours in advance. However, this scheduling constraint does not apply to emails sent via [Marketing Campaigns](https://docs.sendgrid.com/ui/sending-email/how-to-send-email-with-marketing-campaigns/). |[optional] 
**IsAbtest** | **bool** | true if the Single Send's AB Test functionality has been toggled on |
**UpdatedAt** | [**time.Time**](time.Time.md) | the ISO 8601 time at which the Single Send was last updated |
**CreatedAt** | [**time.Time**](time.Time.md) | the ISO 8601 time at which the Single Send was created |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


