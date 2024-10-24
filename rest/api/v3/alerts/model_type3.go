/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Alerts API
* The Twilio SendGrid Alerts API allows you to specify an email address to receive notifications regarding your email usage or statistics. You can set up alerts to be sent to a specific email address on a recurring basis, whether for informational purposes or when specific account actions occur.  For most alerts, you can choose to have the alert sent to you as needed, hourly, daily, weekly, or monthly. The information contained in your alert will be for the last period of the alert. For example, if you choose weekly for the statistics alert, you will receive the statistics for the last week.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// Type3 the model 'Type3'
type Type3 string

// List of Type3
const (
	TYPE3_USAGE_LIMIT        Type3 = "usage_limit"
	TYPE3_STATS_NOTIFICATION Type3 = "stats_notification"
)