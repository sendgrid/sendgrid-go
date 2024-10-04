/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Mail API
* The Twilio SendGrid v3 Mail API allows you to send email at scale over HTTP. The Mail Send endpoint supports many levels of functionality, allowing you to send templates, set categories and custom arguments that can be used to analyze your send, and configure which tracking settings to include such as opens and clicks. You can also group mail sends into batches, allowing you to schedule and cancel sends by their batch IDs.
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

type GetMailBatchParam struct {
	// Use the `on-behalf-of` header to make API calls for a particular Subuser through the parent account. You can use this header to automate bulk updates or to administer a Subuser without changing the authentication in your code. You will use the parent account's API key when using this header.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *GetMailBatchParam) SetOnbehalfof(Onbehalfof string) *GetMailBatchParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This operation allows you to validate a mail batch ID.**  If you provide a valid batch ID, this operation will return a `200` status code and the batch ID itself. If you provide an invalid batch ID, you will receive a `400` level status code and an error message. A batch ID does not need to be assigned to a send to be considered valid. A successful response means only that the batch ID has been created, but it does not indicate that the ID has been assigned to a send.
func (c *ApiService) GetMailBatch(params *GetMailBatchParam) (interface{}, error) {
	path := "/v3/mail/batch/{BatchId}"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &MailBatchResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 401 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 403 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 405 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 500 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
