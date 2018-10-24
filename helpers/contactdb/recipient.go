package contactdb

import (
	"encoding/json"
	"strconv"
)

// Recipient holds the type of a mail recipient.
type Recipient struct {
	ID           string         `json:"id,omitempty"`
	Email        string         `json:"email,omitempty"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	CreatedAt    int            `json:"created_at,omitempty"`
	UpdatedAt    int            `json:"updated_at,omitempty"`
	LastClicked  int            `json:"last_clicked,omitempty"`
	LastEmailed  int            `json:"last_emailed,omitempty"`
	LastOpened   int            `json:"last_opened,omitempty"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}

// recipientCount just contains a recipient count.
// Used for unmarshalling.
type recipientCount struct {
	RecipientCount int `json:"recipient_count,omitempty"`
}

// RecipientResponse is the response type for a 201 response from NewRecipients.
type RecipientResponse struct {
	ErrorCount          int      `json:"error_count,omitempty"`
	ErrorIndices        []int    `json:"error_indices,omitempty"`
	NewCount            int      `json:"new_count,omitempty"`
	PersistedRecipients []string `json:"persisted_recipients,omitempty"`
	UpdatedCount        int      `json:"updated_count,omitempty"`
	Errors              []struct {
		Message      string `json:"message,omitempty"`
		ErrorIndices []int  `json:"error_indices,omitempty"`
	} `json:"errors,omitempty"`
}

// recipientList contains a slice of Recipient.
// Used for unmarshalling.
type recipientList struct {
	Recipients []*Recipient `json:"recipients,omitempty"`
}

// UploadStatus is the type used for checking the upload status of a recipient.
type UploadStatus struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

// uploadStatusList contains a slice of Recipient.
// Used for unmarshalling.
type uploadStatusList struct {
	Status []*UploadStatus `json:"status,omitempty"`
}

// SearchCondition a list of conditions used for searching.
type SearchCondition struct {
	ListID     int          `json:"list_id,omitempty"`
	Conditions []*Condition `json:"conditions,omitempty"`
}

// NewRecipients will create multiple new recipients. (POST)
func NewRecipients(recipients []*Recipient, apiKey string) (*RecipientResponse, error) {
	payload, err := json.Marshal(recipients)
	if err != nil {
		return nil, err
	}

	response, err := SendPOSTRequest(apiKey, "/contactdb/recipients", string(payload))
	if err != nil {
		return nil, err
	}

	recRes := &RecipientResponse{}
	err = json.Unmarshal(response, recRes)

	return recRes, err
}

// GetRecipient will return a specific recipient. (GET)
func GetRecipient(recipientID, apiKey string) (*Recipient, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/recipients/"+recipientID)
	if err != nil {
		return nil, err
	}

	recipients := &recipientList{}
	err = json.Unmarshal(response, recipients)
	if err != nil {
		return nil, err
	}

	return recipients.Recipients[0], nil
}

// GetAllRecipients will return all of your recipients. (GET)
func GetAllRecipients(page, pageSize int, apiKey string) ([]*Recipient, error) {
	queries := map[string]string{
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	path := QueryBuilder("/contactdb/recipients", queries)

	response, err := SendGETRequest(apiKey, path)
	if err != nil {
		return nil, err
	}

	recipients := &recipientList{}
	err = json.Unmarshal(response, recipients)
	if err != nil {
		return nil, err
	}

	return recipients.Recipients, nil
}

// DeleteRecipient will delete a specfic recipient. (DELETE)
func DeleteRecipient(recipientID, apiKey string) error {
	_, err := SendDELETERequest(apiKey, "/contactdb/recipients/"+recipientID, "")
	return err
}

// DeleteRecipients will delete a list of specific recipients. (DELETE)
func DeleteRecipients(recipientIDs []string, apiKey string) error {
	payload, err := json.Marshal(recipientIDs)
	if err != nil {
		return err
	}

	_, err = SendDELETERequest(apiKey, "/contactdb/recipients", string(payload))
	return err
}

// UpdateRecipients will update multiple recipients. (PATCH)
// Will use the email as the recipient identifier.
func UpdateRecipients(recipients []*Recipient, apiKey string) (*RecipientResponse, error) {
	payload, err := json.Marshal(recipients)
	if err != nil {
		return nil, err
	}

	response, err := SendPATCHRequest(apiKey, "/contactdb/recipients", string(payload))
	if err != nil {
		return nil, err
	}

	res := &RecipientResponse{}
	err = json.Unmarshal(response, res)

	return res, err
}

// GetUploadStatus will return all the upload statuses. (GET)
func GetUploadStatus(apiKey string) ([]*UploadStatus, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/status")
	if err != nil {
		return nil, err
	}

	uploadList := &uploadStatusList{}
	err = json.Unmarshal(response, uploadList)
	if err != nil {
		return nil, err
	}

	return uploadList.Status, nil
}

// GetRecipientLists will return all of the lists which the recipient is part of. (GET)
func GetRecipientLists(recipientID, apiKey string) ([]*ContactList, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/recipients/"+recipientID+"/lists")
	if err != nil {
		return nil, err
	}

	lists := &contactLists{}
	err = json.Unmarshal(response, lists)
	if err != nil {
		return nil, err
	}

	return lists.Lists, nil
}

// GetBillableRecipientCount will return how many recipients you will be billed for. (GET)
func GetBillableRecipientCount(apiKey string) (int, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/recipients/billable_count")
	if err != nil {
		return 0, err
	}

	recipientCount := &recipientCount{}
	err = json.Unmarshal(response, recipientCount)
	if err != nil {
		return 0, err
	}

	return recipientCount.RecipientCount, nil
}

// GetTotalRecipientCount will return how many recipients you have in total. (GET)
func GetTotalRecipientCount(apiKey string) (int, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/recipients/count")
	if err != nil {
		return 0, err
	}

	recipientCount := &recipientCount{}
	err = json.Unmarshal(response, recipientCount)
	if err != nil {
		return 0, err
	}

	return recipientCount.RecipientCount, nil
}

// SearchRecipients will search for specific recipients with the specified field-value pairs. (GET)
func SearchRecipients(fieldValues map[string]string, apiKey string) ([]*Recipient, error) {
	path := QueryBuilder("/contactdb/recipients/search", fieldValues)

	response, err := SendGETRequest(apiKey, path)
	if err != nil {
		return nil, err
	}

	recipients := &recipientList{}
	err = json.Unmarshal(response, recipients)
	if err != nil {
		return nil, err
	}

	return recipients.Recipients, nil
}

// SearchRecipientsConditions will search for specific recipients using a SearchCondition. (GET)
func SearchRecipientsConditions(conditions *SearchCondition, apiKey string) ([]*Recipient, error) {
	payload, err := json.Marshal(conditions)
	if err != nil {
		return nil, err
	}

	response, err := SendPOSTRequest(apiKey, "/contactdb/recipients/search", string(payload))
	if err != nil {
		return nil, err
	}

	recipients := &recipientList{}
	err = json.Unmarshal(response, recipients)
	if err != nil {
		return nil, err
	}

	return recipients.Recipients, nil
}
