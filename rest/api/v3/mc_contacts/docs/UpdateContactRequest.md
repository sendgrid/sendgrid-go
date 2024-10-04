# UpdateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListIds** | **[]string** | An array of List ID strings that this contact will be added to. |[optional] 
**Contacts** | [**[]ContactRequest**](ContactRequest.md) | One or more contacts objects that you intend to upsert. Each contact needs to include at least one of `email`, `phone_number_id`, `external_id`, or `anonymous_id` as an identifier. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


