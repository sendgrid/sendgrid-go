# AdvancedStatsMailboxProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | **int32** | The number of emails that were not allowed to be delivered by ISPs. |[optional] 
**Bounces** | **int32** | The number of emails that bounced instead of being delivered. |[optional] 
**Deferred** | **int32** | The number of emails that temporarily could not be delivered. |[optional] 
**Delivered** | **int32** | The number of emails SendGrid was able to confirm were actually delivered to a recipient. |[optional] 
**Drops** | **int32** | The number of emails that were not delivered due to the recipient email address being on a suppression list. |[optional] 
**Requests** | **int32** | The number of emails that were requested to be delivered. |[optional] 
**Processed** | **int32** | Requests from your website, application, or mail client via SMTP Relay or the Web API that SendGrid processed. |[optional] 
**SpamReports** | **int32** | The number of recipients who marked your email as spam. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


