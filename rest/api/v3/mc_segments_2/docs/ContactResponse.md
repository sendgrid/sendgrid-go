# ContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID assigned to a contact when added to the system. |
**Email** | **string** | Email of the contact. This is a reserved field. |[optional] 
**PhoneNumberId** | **string** | The contact's Phone Number ID. This must be a valid phone number. |[optional] 
**ExternalId** | **string** | The contact's External ID. |[optional] 
**AnonymousId** | **string** | The contact's Anonymous ID. |[optional] 
**AlternateEmails** | **[]string** | Alternate emails of the contact. This is a reserved field. |
**FirstName** | **string** | First name of the contact. This is a reserved field. |
**LastName** | **string** | Last name of the contact. This is a reserved field. |
**AddressLine1** | **string** | First line of address of the contact. This is a reserved field. |
**AddressLine2** | **string** | Second line of address of the contact. This is a reserved field. |
**City** | **string** | City associated with the contact. This is a reserved field. |
**StateProvinceRegion** | **string** | State associated with the contact. This is a reserved field. |
**PostalCode** | **int32** | Zipcode associated with the address of the contact. This is a reserved field. |
**Country** | **string** | Country associated with the address of the contact. This is a reserved field. |
**ListIds** | **[]string** | IDs of all lists the contact is part of |[optional] 
**CustomFields** | [**ContactResponseCustomFields**](ContactResponseCustomFields.md) |  |
**SegmentIds** | **[]string** | IDs of all segments the contact is part of |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


