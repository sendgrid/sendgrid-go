package contactdb

// CustomField contains the custom field type.
type CustomField struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"` // date, text, number
	Value string `json:"value,omitempty"`
}

// NewCustomField creates a new custom field. (POST)
func NewCustomField(name, _type, apiKey string) (*CustomField, error) {
	return nil, nil
}

// GetAllCustomFields returns all of your custom fields. (GET)
func GetAllCustomFields(apiKey string) ([]*CustomField, error) {
	return nil, nil
}

// GetCustomField returns a specific custom field. (GET)
func GetCustomField(id int, apiKey string) (*CustomField, error) {
	return nil, nil
}

// DeleteCustomField deletes a specific custom field. (DELETE)
func DeleteCustomField(id int, apiKey string) (*CustomField, error) {
	return nil, nil
}

// GetReservedFields returns all of your resereved fields. (GET)
func GetReservedFields(apiKey string) ([]*CustomField, error) {
	return nil, nil
}
