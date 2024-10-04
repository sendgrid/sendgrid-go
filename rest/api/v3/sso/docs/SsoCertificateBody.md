# SsoCertificateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicCertificate** | **string** | This certificate is used by Twilio SendGrid to verify that SAML requests are coming from Okta. This is called the X509 certificate in the Twilio SendGrid UI. |[optional] 
**Id** | **float32** | A unique ID assigned to the certificate by SendGrid. |[optional] 
**NotBefore** | **float32** | A unix timestamp (e.g., 1603915954) that indicates the time before which the certificate is not valid. |[optional] 
**NotAfter** | **float32** | A unix timestamp (e.g., 1603915954) that indicates the time after which the certificate is no longer valid. |[optional] 
**IntergrationId** | **string** | An ID that matches a certificate to a specific IdP integration. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


