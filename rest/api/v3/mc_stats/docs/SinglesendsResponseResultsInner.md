# SinglesendsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is the ID of the Single Send you require stats for. |
**AbVariation** | **string** | This is the A/B variation of the Single Send stat returned. If the `group_by` parameter doesn't include `ab_variation` in the request, then the value is \"all\". |[default to "a14dcc63-d651-4c57-9826-4a3705f5c78d"]
**AbPhase** | [**AbPhase**](AbPhase.md) | This is the A/B phase of the Single Send stat returned. If the `group_by` parameter doesn't include `ab_phase` in the request, then the value is \"all\". |
**Aggregation** | **string** | This describes the time unit to which the stat is rolled up. It is based on the `aggregated_by` parameter included in the request. It can be \"total\" or the date (in YYYY-MM-DD format) the stats are for. |[optional] [default to "total"]
**Stats** | [**Metrics**](Metrics.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


