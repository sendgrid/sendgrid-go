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

import (
	"encoding/json"
	"net/url"
)

type WarmUpIpParam struct {
	//
	WarmUpIpRequest *WarmUpIpRequest `json:"WarmUpIpRequest,omitempty"`
}

func (params *WarmUpIpParam) SetWarmUpIpRequest(WarmUpIpRequest WarmUpIpRequest) *WarmUpIpParam {
	params.WarmUpIpRequest = &WarmUpIpRequest
	return params
}

// **This endpoint allows you to put an IP address into warmup mode.**
func (c *ApiService) WarmUpIp(params *WarmUpIpParam) (interface{}, error) {
	path := "/v3/ips/warmup"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.WarmUpIpRequest != nil {
		b, err := json.Marshal(*params.WarmUpIpRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &[]IpWarmup200Inner{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &WarmUpIp404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
