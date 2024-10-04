/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Email Address Validation API
* The Twilio SendGrid Email Address Validation API provides real-time detailed information on the validity of email addresses. You can integrate this validation process into your platform's signup form and customize the best use of email address validation for your use case.  Email Address Validation is available to [Email API Pro and Premier level accounts](https://sendgrid.com/pricing) only. Prior to upgrading your account to Pro or Premier, you will not see the option to create an Email Validation API key. An Email Validation API key is separate from and in addition to your other keys, including Full Access API keys.  You can use this API to: - Indicate to users that the address they have entered into a form is invalid. - Drop invalid email addresses from your database. - Suppress invalid email addresses from your sending to decrease your bounce rate.  See [**Email Address Validation**](https://docs.sendgrid.com/ui/managing-contacts/email-address-validation) for more information.  You can also view your Email Validation results and metrics in the Validation section of the [Twilio SendGrid application user interface](https://docs.sendgrid.com/ui/managing-contacts/email-address-validation).
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"net/url"

	"strings"
)

type GetEmailJobForVerificationParam struct {
	// The ID of the Bulk Email Address Validation Job you wish to retrieve.
	JobId *string `json:"job_id"`
}

func (params *GetEmailJobForVerificationParam) SetJobId(JobId string) *GetEmailJobForVerificationParam {
	params.JobId = &JobId
	return params
}

// **This endpoint returns a specific Bulk Email Validation Job. You can use this endpoint to check on the progress of a Job.**
func (c *ApiService) GetEmailJobForVerification(params *GetEmailJobForVerificationParam) (interface{}, error) {
	path := "/v3/validations/email/jobs/{JobId}"
	if params != nil && params.JobId != nil {
		path = strings.Replace(path, "{"+"JobId"+"}", *params.JobId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &GetValidationsEmailJobsJobId200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
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
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
