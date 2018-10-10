package contactdb

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

// UploadStatus is the type used for checking the upload status of a recipient.
type UploadStatus struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

// NewRecipients will create multiple new recipients. (POST)
func NewRecipients(recipients []*Recipient, apiKey string) (*RecipientResponse, error) {
	return nil, nil
}

// GetRecipient will return a specific recipient. (GET)
func GetRecipient(recipientID, apiKey string) (*Recipient, error) {
	return nil, nil
}

// GetAllRecipients will return all of your recipients. (GET)
func GetAllRecipients(apiKey string) ([]*Recipient, error) {
	return nil, nil
}

// DeleteRecipient will delete a specfic recipient. (DELETE)
func DeleteRecipient(recipientID, apiKey string) error {
	return nil
}

// DeleteRecipients will delete a list of specific recipients. (DELETE)
func DeleteRecipients(recipientIDs []string, apiKey string) error {
	return nil
}

// UpdateRecipients will update multiple recipients. (PATCH)
// Will use the email as the recipient identifier.
func UpdateRecipients(recipients []*Recipient, apiKey string) (*RecipientResponse, error) {
	return nil, nil
}

// GetUploadStatus will return all the upload statuses. (GET)
func GetUploadStatus(apiKey string) ([]*UploadStatus, error) {
	return nil, nil
}

// GetRecipientLists will return all of the lists which the recipient is part of. (GET)
func GetRecipientLists(recipientID, apiKey string) ([]*ContactList, error) {
	return nil, nil
}

// GetBillableRecipientCount will return how many recipients you will be billed for. (GET)
func GetBillableRecipientCount(apiKey string) (int, error) {
	return 0, nil
}

// GetTotalRecipientCount will return how many recipients you have in total. (GET)
func GetTotalRecipientCount(apiKey string) (int, error) {
	return 0, nil
}

// SearchRecipients will search for specific recipients with the specified field-value pairs. (GET)
func SearchRecipients(fieldValues []*CustomField, apiKey string) ([]*Recipient, error) {
	return nil, nil
}

// SearchRecipientsSegment will search for specific recipients using a segment. (GET)
func SearchRecipientsSegment(segment *Segment, apiKey string) ([]*Recipient, error) {
	return nil, nil
}
