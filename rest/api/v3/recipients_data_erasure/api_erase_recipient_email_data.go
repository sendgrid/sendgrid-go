/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Recipients' Data Erasure API
* The Recipients' Data Erasure API allows Twilio SendGrid customers to delete their own customers' personal data from the Twilio SendGrid Platform. The use of this API facilitates the self-service removal of email personal data from the Twilio SendGrid platform and will enable customers to comply with various obligatory data privacy regulations.
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

type EraseRecipientEmailDataParam struct {
	//
	RecipientsDataErasureEraseRecipientsRequest *RecipientsDataErasureEraseRecipientsRequest `json:"RecipientsDataErasureEraseRecipientsRequest"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *EraseRecipientEmailDataParam) SetRecipientsDataErasureEraseRecipientsRequest(RecipientsDataErasureEraseRecipientsRequest RecipientsDataErasureEraseRecipientsRequest) *EraseRecipientEmailDataParam {
	params.RecipientsDataErasureEraseRecipientsRequest = &RecipientsDataErasureEraseRecipientsRequest
	return params
}
func (params *EraseRecipientEmailDataParam) SetOnbehalfof(Onbehalfof string) *EraseRecipientEmailDataParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This operation allows you to delete your recipients' personal email data**  The Delete Recipients' Email Data operation accepts a list of 5,000 `email_addresses` or a total payload size of 256Kb per request, whichever comes first. Upon a successful request with this operation, SendGrid will run a search on the email addresses provided against the SendGrid system to identify matches. SendGrid will then delete all personal data associated with the matched users such as the recipients' names, email addresses, subject lines, categories, and IP addresses.  All email addresses are filtered for uniqueness and tested for structural validity—any invalid addresses will be returned in an error response.  Please note that recipient data is deleted for the account making the request only—deletions do not cascade from a parent account to its Subusers' recipients. To delete a Subuser's recipients' data, you can use the `on-behalf-of` header.
func (c *ApiService) EraseRecipientEmailData(params *EraseRecipientEmailDataParam) (interface{}, error) {
	path := "/v3/recipients/erasejob"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.RecipientsDataErasureEraseRecipientsRequest != nil {
		b, err := json.Marshal(*params.RecipientsDataErasureEraseRecipientsRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 202 {
		ps := &RecipientsDataErasureJobId{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
