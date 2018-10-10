package contactdb

// ContactList holds the type for a list of contacts.
type ContactList struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	RecipientCount int    `json:"recipient_count,omitempty"`
}

// NewContactList creates a new contact list. (POST)
func NewContactList(name, apiKey string) (*ContactList, error) {
	return nil, nil
}

// GetContactList returns a specific contact list. (GET)
func GetContactList(id int, apiKey string) (*ContactList, error) {
	return nil, nil
}

// GetAllContactLists returns all of your contact lists. (GET)
func GetAllContactLists(apiKey string) (*ContactList, error) {
	return nil, nil
}

// DeleteContactList deletes a contact list. (DELETE)
func DeleteContactList(id int, apiKey string) error {
	return nil
}

// DeleteContactLists deletes a list of specific contact lists. (DELETE)
func DeleteContactLists(ids []int, apiKey string) error {
	return nil
}

// UpdateContactList updates a contact list. (PATCH)
func UpdateContactList(listID int, newName, apiKey string) (*ContactList, error) {
	return nil, nil
}

// GetAllRecipientsFromList will get all recipients from a specfic list. (GET)
func GetAllRecipientsFromList(page, pageSize, listID int, apiKey string) ([]*Recipient, error) {
	return nil, nil
}

// AddRecipientToList will add a specfic recipient to a specfic list. (POST)
func AddRecipientToList(listID int, recipientID, apiKey string) error {
	return nil
}

// AddRecipientsToList will add multiple recipients to a specfic list. (POST)
func AddRecipientsToList(listID int, recipientIDs []string, apiKey string) error {
	return nil
}

// DeleteRecipientFromList will remove a specfic recipient from a specific list. (DELETE)
func DeleteRecipientFromList(listID int, recipientID, apiKey string) error {
	return nil
}
