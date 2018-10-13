package contactdb

import (
	"encoding/json"
	"strconv"
)

// CustomField contains the custom field type.
type CustomField struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"` // date, text, number
	Value string `json:"value,omitempty"`
}

// reservedFields is just a list of custom fields.
// It is used for unmarshalling.
type reservedFields struct {
	ReservedFields []*CustomField `json:"reserved_fields,omitempty"`
}

// customFields is just a list of custom fields.
// It is used for unmarshalling.
type customFields struct {
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}

// NewCustomField creates a new custom field. (POST)
func NewCustomField(name, _type, apiKey string) (*CustomField, error) {
	byteArr, err := json.Marshal(CustomField{
		Name: name,
		Type: _type,
	})
	if err != nil {
		return nil, err
	}

	response, err := SendPOSTRequest(apiKey, "/contactdb/custom_fields", string(byteArr))
	if err != nil {
		return nil, err
	}

	customField := &CustomField{}
	err = json.Unmarshal(response, customField)
	if err != nil {
		return nil, err
	}

	return customField, nil
}

// GetAllCustomFields returns all of your custom fields. (GET)
func GetAllCustomFields(apiKey string) ([]*CustomField, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/custom_fields")
	if err != nil {
		return nil, err
	}

	customFields := &customFields{}
	err = json.Unmarshal(response, customFields)
	if err != nil {
		return nil, err
	}

	return customFields.CustomFields, nil
}

// GetCustomField returns a specific custom field. (GET)
func GetCustomField(id int, apiKey string) (*CustomField, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/custom_field/"+strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	customField := &CustomField{}
	err = json.Unmarshal(response, customField)
	if err != nil {
		return nil, err
	}

	return customField, nil
}

// DeleteCustomField deletes a specific custom field. (DELETE)
func DeleteCustomField(id int, apiKey string) error {
	_, err := SendDELETERequest(apiKey, "/contactdb/custom_fields/"+strconv.Itoa(id), "")
	if err != nil {
		return err
	}

	return nil
}

// GetReservedFields returns all of your resereved fields. (GET)
func GetReservedFields(apiKey string) ([]*CustomField, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/reserved_fields")
	if err != nil {
		return nil, err
	}

	fields := &reservedFields{}
	err = json.Unmarshal(response, fields)
	if err != nil {
		return nil, err
	}

	return fields.ReservedFields, nil
}
