# ListSubuserByTemplate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasRestrictedSubuserAccess** | **bool** | When this property is set to `true`, the Teammate has permissions to operate only on behalf of a Subuser. This property value is `true` when the `subuser_access` property is not empty. The `subuser_access` property determines which Subusers the Teammate may act on behalf of. |[optional] 
**SubuserAccess** | [**[]ListSubuserByTemplate200ResponseSubuserAccessInner**](ListSubuserByTemplate200ResponseSubuserAccessInner.md) | Specifies which Subusers the Teammate may access and act on behalf of. If this property is populated, the `has_restricted_subuser_access` property will be `true`. |[optional] 
**Metadata** | [**ListSubuserByTemplate200ResponseMetadata**](ListSubuserByTemplate200ResponseMetadata.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


