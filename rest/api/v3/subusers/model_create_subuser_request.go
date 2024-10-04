/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Subusers
* The Twilio SendGrid Subusers API allows you to create and manage your Subuser accounts. Subusers are available on [Pro and Premier plans](https://sendgrid.com/pricing), and you can think of them as sub-accounts. Each Subuser can have its own sending domains, IP addresses, and reporting. SendGrid recommends creating Subusers for each of the different types of emails you send—one Subuser for transactional emails and another for marketing emails. Independent Software Vendor (ISV) customers may also create Subusers for each of their customers.  You can also manage Subusers in the [Twilio SendGrid application user interface](https://app.sendgrid.com/settings/subusers). See [**Subusers**](https://docs.sendgrid.com/ui/account-and-settings/subusers) for more information.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// CreateSubuserRequest struct for CreateSubuserRequest
type CreateSubuserRequest struct {
	// The username for this subuser.
	Username string `json:"username"`
	// The email address of the subuser.
	Email string `json:"email"`
	// The password this subuser will use when logging into SendGrid.
	Password string `json:"password"`
	// The IP addresses that should be assigned to this subuser.
	Ips []string `json:"ips"`
	// The region this Subuser should be assigned to. Can be `global` or `eu`. (Regional email is in Public Beta and requires SendGrid Pro plan or above.).
	Region *Region1 `json:"region,omitempty"`
	// A flag that determines if the Subuser's region should be returned in the response. (Regional email is in Public Beta and requires SendGrid Pro plan or above.)
	IncludeRegion *bool `json:"include_region,omitempty"`
}
