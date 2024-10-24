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

// SubuserCreditsRequest struct for SubuserCreditsRequest
type SubuserCreditsRequest struct {
	// Type determines how credits are reset for a Subuser. `unlimited` indicates that there is no limit to the Subuser's credits. `recurring` indicates that the credits for the Subuser are reset according to the frequency determined by `reset_frequency`. `nonrecurring` indicates that there is no recurring schedule to reset credits and resets must be done on an ad hoc basis.
	Type Type1 `json:"type"`
	// The frequency with which a Subuser's credits are reset if `type` is set to `recurring`. Do _not_ include `reset_frequency` if you choose a reset `type` value of `unlimited` or `nonrecurring`.
	ResetFrequency *ResetFrequency1 `json:"reset_frequency,omitempty"`
	// Total number of credits to which the Subuser is to be reset. If `type` is `nonrecurring` then the Subuser's credits will be reset to `total` on a one-time basis. If `type` is `recurring` then the Subuser's credits will be reset to `total` every time a reset is scheduled in accordance with the `reset_frequency`. Do _not_ include `total` if you choose a reset `type` value of `unlimited`.
	Total *int32 `json:"total,omitempty"`
}