# SeqMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngagementRecency** | **float32** | Indicates whether or not you are sending to an engaged audience. Engagement recency is determined by factors such as how regularly your messages are being opened and clicked. Your score will range from `1` to `5` with a higher score representing better engagement quality. |
**OpenRate** | **float32** | Indicates the degree to which your audience is opening your messages. Your score will range from `1` to `5` with a higher score representing a better open rate and better engagement quality. |
**BounceClassification** | **float32** | Indicates the degree to which mailbox providers are rejecting your mail due to reputation issues or content that looks like spam. Your score is calculated based on a ratio of these specific types of bounces to your total processed email. Your score will range from `1` to `5` with a higher score representing a lower percentage of bounces and better engagement quality. |
**BounceRate** | **float32** | Indicates if you are attempting to send too many messages to addresses that don't exist. This score is determined by calculating your permanent bounces (bounces to invalid mailboxes). Your score will range from `1` to `5` with a higher score representing fewer bounces and better engagement quality. |
**SpamRate** | **float32** | Indicates if your messages are generating excessive spam complaints from recipients. This score is determined by calculating the number of recipients who open your mail and then mark it as spam. Your score will range from `1` to `5` with a higher score representing fewer spam reports and better engagement quality. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


