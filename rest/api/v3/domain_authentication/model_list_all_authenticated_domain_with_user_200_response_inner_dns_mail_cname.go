/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Domain Authentication API
* The Twilio SendGrid Domain Authentication API allows you to manage your authenticated domains and their settings.  Domain Authentication is a required step when setting up your Twilio SendGrid account because it's essential to ensuring the deliverability of your email. Domain Authentication signals trustworthiness to email inbox providers and your recipients by approving SendGrid to send email on behalf of your domain. For more information, see [**How to Set Up Domain Authentication**](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-domain-authentication/).  Each user may have a maximum of 3,000 authenticated domains and 3,000 link brandings. This limit is at the user level, meaning each Subuser belonging to a parent account may have its own 3,000 authenticated domains and 3,000 link brandings.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// ListAllAuthenticatedDomainWithUser200ResponseInnerDnsMailCname The CNAME for your sending domain that points to sendgrid.net.
type ListAllAuthenticatedDomainWithUser200ResponseInnerDnsMailCname struct {
	// Indicates if this is a valid CNAME.
	Valid bool `json:"valid"`
	// The type of DNS record.
	Type string `json:"type"`
	// The domain that this CNAME is created for.
	Host string `json:"host"`
	// The CNAME record.
	Data string `json:"data"`
}