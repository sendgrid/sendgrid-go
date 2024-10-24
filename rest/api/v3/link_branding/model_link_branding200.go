/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Link Branding API
* The Twilio SendGrid Link Branding API allows you to configure your domain settings so that all of the click-tracked links, opens, and images in your emails are served from your domain rather than `sendgrid.net`. Spam filters and recipient servers look at the links within emails to determine whether the email appear trustworthy. They use the reputation of the root domain to determine whether the links can be trusted.  You can also manage Link Branding in the **Sender Authentication** section of the Twilio SendGrid application user interface.   See [**How to Set Up Link Branding**](https: //sendgrid.com/docs/ui/account-and-settings/how-to-set-up-link-branding/) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// LinkBranding200 struct for LinkBranding200
type LinkBranding200 struct {
	// The ID of the branded link.
	Id int32 `json:"id"`
	// The root domain of the branded link.
	Domain string `json:"domain"`
	// The subdomain used to generate the DNS records for this link branding. This subdomain must be different from the subdomain used for your authenticated domain.
	Subdomain *string `json:"subdomain,omitempty"`
	// The username of the account that this link branding is associated with.
	Username string `json:"username"`
	// The ID of the user that this link branding is associated with.
	UserId int32 `json:"user_id"`
	// Indicates if this is the default link branding.
	Default Default2 `json:"default"`
	// Indicates if this link branding is valid.
	Valid Valid3 `json:"valid"`
	// Indicates if this link branding was created using the legacy whitelabel tool. If it is a legacy whitelabel, it will still function, but you'll need to create new link branding if you need to update it.
	Legacy Legacy             `json:"legacy"`
	Dns    LinkBranding200Dns `json:"dns"`
}