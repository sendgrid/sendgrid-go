# AutmoationsLinkStatsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | This is the URL of the link clicked. If `{{custom_fields}}` are part of the URL, they will be included. |
**UrlLocation** | **int32** | This is the location of the link clicked in each Automation step. Links are located according to their position within the message; the topmost link has index `0`. |[optional] 
**StepId** | **string** | This is the ID of the step if the stats were requested to be grouped by `step_id`. |
**Clicks** | **int32** | The number of clicks on this particular link. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


