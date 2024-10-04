# SinglesendsLinkStatsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | This is the URL of the link clicked. If `{{custom_fields}}` are part of the URL, they will be included. |
**UrlLocation** | **int32** | This is the location of the link clicked in each Single Send A/B variation, or in the Single Send itself if there are no variations. Links are numbered from the top down; the topmost link is index `0`. |[optional] 
**AbVariation** | **string** | This is the A/B variation of the Single Send stat returned. It is set to `\"all\"` if the `ab_variation` query parameter was not set in the request and `group_by` doesn't contain `ab_variation`. |
**AbPhase** | [**AbPhase1**](AbPhase1.md) | This is the A/B phase of the Single Send stat returned. If the `ab_phase` query parameter was not provided, it will return `\"all\"`. |
**Clicks** | **int32** | the number of clicks on this particular link |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


