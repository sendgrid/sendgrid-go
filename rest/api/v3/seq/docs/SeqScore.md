# SeqScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The user identifier associated with this score. |
**Username** | **string** | The username associated with this score. |[optional] 
**Date** | **string** | The date (UTC) for which this score was calculated. |[optional] 
**Score** | **float32** | Your SendGrid Engagement Quality Score. The `score` property will be omitted when a score is unknown. A score may be unknown if you have not turned on [open tracking](https://docs.sendgrid.com/ui/account-and-settings/tracking) and/or your account or Subuser has not sent at least 1,000 messages during the previous 30 days. |[optional] 
**Metrics** | [**SeqMetrics**](SeqMetrics.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


