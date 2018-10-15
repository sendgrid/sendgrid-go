package contactdb

import (
	"encoding/json"
	"strconv"
)

// ContactList holds the type for a list of contacts.
type ContactList struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	RecipientCount int    `json:"recipient_count,omitempty"`
}

// contactLists contains a slice of ContactList.
// Used for unmarshalling.
type contactLists struct {
	Lists []*ContactList `json:"lists,omitempty"`
}

// recipientList contains a slice of Recipient.
// Used for unmarshalling.
type recipientList struct {
	Recipients []*Recipient `json:"recipients,omitempty"`
}

// NewContactList creates a new contact list. (POST)
func NewContactList(name, apiKey string) (*ContactList, error) {
	byteArr, err := json.Marshal(ContactList{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	response, err := SendPOSTRequest(apiKey, "/contactdb/lists", string(byteArr))
	if err != nil {
		return nil, err
	}

	contactList := &ContactList{}
	err = json.Unmarshal(response, contactList)
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

// GetContactList returns a specific contact list. (GET)
func GetContactList(id int, apiKey string) (*ContactList, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/lists/"+strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	contactList := &ContactList{}
	err = json.Unmarshal(response, contactList)
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

// GetAllContactLists returns all of your contact lists. (GET)
func GetAllContactLists(apiKey string) ([]*ContactList, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/lists")
	if err != nil {
		return nil, err
	}

	contactLists := &contactLists{}
	err = json.Unmarshal(response, contactLists)
	if err != nil {
		return nil, err
	}

	return contactLists.Lists, nil
}

// DeleteContactList deletes a contact list. (DELETE)
func DeleteContactList(id int, deleteContacts bool, apiKey string) error {
	path := "/contactdb/lists/" + strconv.Itoa(id) + "?delete_contacts=" + strconv.FormatBool(deleteContacts)
	_, err := SendDELETERequest(apiKey, path, "")
	return err
}

// DeleteContactLists deletes a list of specific contact lists. (DELETE)
func DeleteContactLists(ids []int, apiKey string) error {
	byteArr, err := json.Marshal(ids)
	if err != nil {
		return err
	}

	_, err = SendDELETERequest(apiKey, "/contactdb/lists", string(byteArr))
	if err != nil {
		return err
	}

	return nil
}

// UpdateContactList updates a contact list. (PATCH)
func UpdateContactList(listID int, newName, apiKey string) (*ContactList, error) {
	byteArr, err := json.Marshal(ContactList{
		Name: newName,
	})
	if err != nil {
		return nil, err
	}

	queries := map[string]string{
		"list_id": strconv.Itoa(listID),
	}

	path := QueryBuilder("/contactdb/lists/"+strconv.Itoa(listID), queries)
	response, err := SendPATCHRequest(apiKey, path, string(byteArr))
	if err != nil {
		return nil, err
	}

	contactList := &ContactList{}
	err = json.Unmarshal(response, contactList)
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

// GetAllRecipientsFromList will get all recipients from a specfic list. (GET)
// Default Values: page = 1, pageSize = 20
func GetAllRecipientsFromList(page, pageSize, listID int, apiKey string) ([]*Recipient, error) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	queries := map[string]string{
		"list_id":   strconv.Itoa(listID),
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}

	path := QueryBuilder("/contactdb/lists/"+strconv.Itoa(listID)+"/recipients", queries)
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

// AddRecipientToList will add a specfic recipient to a specfic list. (POST)
func AddRecipientToList(listID int, recipientID, apiKey string) error {
	_, err := SendPOSTRequest(apiKey, "/contactdb/lists/"+strconv.Itoa(listID)+"/recipients/"+recipientID, "")
	return err
}

// AddRecipientsToList will add multiple recipients to a specfic list. (POST)
func AddRecipientsToList(listID int, recipientIDs []string, apiKey string) error {
	recipients, err := json.Marshal(recipientIDs)
	if err != nil {
		return err
	}

	_, err = SendPOSTRequest(apiKey, "/contactdb/lists/"+strconv.Itoa(listID)+"recipients", string(recipients))
	return err
}

// DeleteRecipientFromList will remove a specfic recipient from a specific list. (DELETE)
func DeleteRecipientFromList(listID int, recipientID, apiKey string) error {
	queries := map[string]string{
		"list_id":      strconv.Itoa(listID),
		"recipient_id": recipientID,
	}

	path := QueryBuilder("/contactdb/lists/"+strconv.Itoa(listID)+"/recipients/"+recipientID, queries)
	_, err := SendDELETERequest(apiKey, path, "")
	return err
}
