# ValidateEmail200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email being validated |
**Verdict** | [**Verdict**](Verdict.md) | A generic classification of whether or not the email address is valid. |
**Score** | **float32** | A numeric representation of the email validity. |
**Local** | **string** | The local part of the email address. |
**Host** | **string** | The domain of the email address. |
**Suggestion** | **string** | A suggested correction in the event of domain name typos (e.g., gmial.com) |[optional] 
**Checks** | [**ValidateEmail200ResponseResultChecks**](ValidateEmail200ResponseResultChecks.md) |  |
**Source** | **string** | The source of the validation, as per the API request. |[optional] 
**IpAddress** | **string** | The IP address associated with this email. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


