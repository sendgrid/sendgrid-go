# SubuserCreditsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Type1**](Type1.md) | Type determines how credits are reset for a Subuser. `unlimited` indicates that there is no limit to the Subuser's credits. `recurring` indicates that the credits for the Subuser are reset according to the frequency determined by `reset_frequency`. `nonrecurring` indicates that there is no recurring schedule to reset credits and resets must be done on an ad hoc basis. |
**ResetFrequency** | [**ResetFrequency1**](ResetFrequency1.md) | The frequency with which a Subuser's credits are reset if `type` is set to `recurring`. Do _not_ include `reset_frequency` if you choose a reset `type` value of `unlimited` or `nonrecurring`. |[optional] 
**Total** | **int32** | Total number of credits to which the Subuser is to be reset. If `type` is `nonrecurring` then the Subuser's credits will be reset to `total` on a one-time basis. If `type` is `recurring` then the Subuser's credits will be reset to `total` every time a reset is scheduled in accordance with the `reset_frequency`. Do _not_ include `total` if you choose a reset `type` value of `unlimited`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


