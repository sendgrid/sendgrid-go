/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Contacts API
* The Twilio SendGrid Marketing Campaigns Contacts API allows you to manage all of your marketing contacts programmatically. You can also import and export contacts using this API. The Contacts API allows you to associate contacts with lists and segments; however, to manage the lists and segments themselves, see the [Lists API](https://docs.sendgrid.com/api-reference/lists/) and [Segments API](https://docs.sendgrid.com/api-reference/segmenting-contacts-v2/).  You can also manage your marketing contacts with the [Marketing Campaigns application user interface](https://mc.sendgrid.com/contacts). See [**How to Send Email with New Marketing Campaigns**](https://docs.sendgrid.com/ui/sending-email/how-to-send-email-with-marketing-campaigns) for more information.
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
)

type DeleteContactParam struct {
	// Must be set to `\"true\"` to delete all contacts.
	DeleteAllContacts *string `json:"delete_all_contacts,omitempty"`
	// A comma-separated list of contact IDs.
	Ids *string `json:"ids,omitempty"`
}

func (params *DeleteContactParam) SetDeleteAllContacts(DeleteAllContacts string) *DeleteContactParam {
	params.DeleteAllContacts = &DeleteAllContacts
	return params
}
func (params *DeleteContactParam) SetIds(Ids string) *DeleteContactParam {
	params.Ids = &Ids
	return params
}

// **This endpoint can be used to delete one or more contacts**.  The query parameter `ids` must set to a comma-separated list of contact IDs for bulk contact deletion.  The query parameter `delete_all_contacts` must be set to `\"true\"` to delete **all** contacts.   You must set either `ids` or `delete_all_contacts`.  Deletion jobs are processed asynchronously.  Twilio SendGrid recommends exporting your contacts regularly as a backup to avoid issues or lost data.
func (c *ApiService) DeleteContact(params *DeleteContactParam) (interface{}, error) {
	path := "/v3/marketing/contacts"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.DeleteAllContacts != nil {
		data.Set("delete_all_contacts", *params.DeleteAllContacts)
	}
	if params != nil && params.Ids != nil {
		data.Set("ids", *params.Ids)
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 202 {
		ps := &DeleteContact202Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &DeleteContact400Response{}
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
	if resp.StatusCode == 404 {
		ps := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 500 {
		ps := &SearchContact500Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
