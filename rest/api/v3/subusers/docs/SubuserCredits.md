# SubuserCredits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Type**](Type.md) | Type determines how credits are reset for a Subuser. `unlimited` indicates that there is no limit to the Subuser's credits. `recurring` indicates that the credits for the Subuser are reset according to the frequency determined by `reset_frequency`. `nonrecurring` indicates that there is no recurring schedule to reset credits and resets must be done on an ad hoc basis. |
**ResetFrequency** | [**ResetFrequency**](ResetFrequency.md) | The frequency with which a Subuser's credits are reset if `type` is set to `recurring`, otherwise `null`. |
**Remain** | **int32** | Total number of remaining credits. `remain` is `null` if the reset `type` for the Subuser's credits is set to `unlimited`. |
**Total** | **int32** | Total number of allowable credits. `total` is `null` if the reset `type` for the Subuser's credits is set to `unlimited` or `nonrecurring`. |
**Used** | **int32** | Total number of used credits. `used` is `null` if the reset `type` for the Subuser's credits is set to `unlimited` or `nonrecurring`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


