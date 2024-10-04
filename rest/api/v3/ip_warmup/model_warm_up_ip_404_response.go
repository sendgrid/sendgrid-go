/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid IP Warmup API
* The Twilio SendGrid IP Warm Up API allows you to gradually increasing the volume of mail sent with a dedicated IP address according to a predetermined schedule. This gradual process helps to establish a reputation with ISPs (Internet Service Providers) as a legitimate email sender.  SendGrid can automatically warm up dedicated IP addresses by limiting the amount of mail that can be sent through them per hour. The limit is determined by how long the IP address has been warming up.  See the [warmup schedule](https://sendgrid.com/docs/ui/sending-email/warming-up-an-ip-address/#automated-ip-warmup-hourly-send-schedule) to learn how SendGrid limits your email traffic for IPs in warmup.  You can also choose to use Twilio SendGrid's automated IP warmup for any of your IPs from the **IP Addresses** settings menu in the [Twilio SendGrid application user interface](https://app.sendgrid.com/settings/ip_addresses).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

// WarmUpIp404Response struct for WarmUpIp404Response
type WarmUpIp404Response struct {
	// The errors that were encountered.
	Errors *[]WarmUpIp404ResponseErrorsInner `json:"errors,omitempty"`
}
