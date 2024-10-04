# AutomationsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is the ID of the Automation you are requesting stats for. |
**Aggregation** | **string** | This describes the time unit to which the stat is rolled up. It is based on the `aggregated_by` parameter included in the request. It can be \"total\" or the date (in YYYY-MM-DD format) the stats are for. |[default to "total"]
**StepId** | **string** | This is the ID of the step if the stats were requested to be grouped by `step_id`. |[default to "all"]
**Stats** | [**Metrics**](Metrics.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


