# PostPatchIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your integration. This name can be anything that makes sense for your organization (eg. Twilio SendGrid) |
**Enabled** | **bool** | Indicates if the integration is enabled. |
**SigninUrl** | **string** | The IdP's SAML POST endpoint. This endpoint should receive requests and initiate an SSO login flow. This is called the \"Embed Link\" in the Twilio SendGrid UI. |
**SignoutUrl** | **string** | This URL is relevant only for an IdP-initiated authentication flow. If a user authenticates from their IdP, this URL will return them to their IdP when logging out. |
**EntityId** | **string** | An identifier provided by your IdP to identify Twilio SendGrid in the SAML interaction. This is called the \"SAML Issuer ID\" in the Twilio SendGrid UI. |
**CompletedIntegration** | **bool** | Indicates if the integration is complete. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


