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

// ListReputation200ResponseInner struct for ListReputation200ResponseInner
type ListReputation200ResponseInner struct {
	// The sender reputation this subuser has attained.
	Reputation float32 `json:"reputation"`
	// The subuser that has this reputation.
	Username string `json:"username"`
}
