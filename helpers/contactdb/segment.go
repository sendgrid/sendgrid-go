package contactdb

import (
	"encoding/json"
	"strconv"
)

// Segment holds the struct for the segments
type Segment struct {
	ID             int          `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	ListID         int          `json:"list_id,omitempty"`
	Conditions     []*Condition `json:"conditions,omitempty"`
	RecipientCount int          `json:"recipient_count,omitempty"`
}

// segmentList is a list of Segments.
// Used for unmarshalling.
type segmentList struct {
	Segments []*Segment `json:"segments,omitempty"`
}

// Condition is a condition
type Condition struct {
	AndOr    string            `json:"and_or"`
	Field    string            `json:"field,omitempty"`
	Value    string            `json:"value,omitempty"`
	Operator ConditionOperator `json:"operator,omitempty"`
}

// ConditionOperator is the operator for a condition.
type ConditionOperator string

// This is the list of ConditionOperators.
const (
	OperatorEQ       ConditionOperator = "eq"
	OperatorNE       ConditionOperator = "ne"
	OpeartorLT       ConditionOperator = "lt"
	OperatorGT       ConditionOperator = "gt"
	OperatorContains ConditionOperator = "contains"
)

// NewSegment creats a new segment. (POST)
func NewSegment(listID int, conditions []*Condition, name, apiKey string) (*Segment, error) {
	payload, err := json.Marshal(Segment{
		ListID:     listID,
		Conditions: conditions,
		Name:       name,
	})
	if err != nil {
		return nil, err
	}

	response, err := SendPOSTRequest(apiKey, "/contactdb/segments", string(payload))
	if err != nil {
		return nil, err
	}

	segment := &Segment{}
	err = json.Unmarshal(response, segment)
	if err != nil {
		return nil, err
	}

	return segment, nil
}

// GetSegment will return a specific segment by ID. (GET)
func GetSegment(segID, apiKey string) (*Segment, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/segments/"+segID+"?segment_id="+segID)
	if err != nil {
		return nil, err
	}

	segment := &Segment{}
	err = json.Unmarshal(response, segment)
	if err != nil {
		return nil, err
	}

	return segment, nil
}

// GetAllSegments will return all of your segments. (GET)
func GetAllSegments(apiKey string) ([]*Segment, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/segments")
	if err != nil {
		return nil, err
	}

	segments := &segmentList{}
	err = json.Unmarshal(response, segments)
	if err != nil {
		return nil, err
	}

	return segments.Segments, nil
}

// UpdateSegment will update a specific segment with new values. (PATCH)
func UpdateSegment(segment *Segment, segID, apiKey string) (*Segment, error) {
	payload, err := json.Marshal(segment)
	if err != nil {
		return nil, err
	}

	response, err := SendPATCHRequest(apiKey, "/contactdb/segments/"+segID, string(payload))
	if err != nil {
		return nil, err
	}

	newSeg := &Segment{}
	err = json.Unmarshal(response, newSeg)
	if err != nil {
		return nil, err
	}

	return newSeg, nil
}

// DeleteSegment will delete a specific segment. (DELETE)
// DeleteContacts: True to delete all contacts matching the segment in
// addition to deleting the segment.
func DeleteSegment(deleteContacts bool, segID, apiKey string) error {
	path := "/contactdb/segments/" + segID + "?delete_contacts=" + strconv.FormatBool(deleteContacts)
	_, err := SendDELETERequest(apiKey, path, "")
	return err
}

// GetSegmentRecipients will return all of the recipients of a segment. (GET)
func GetSegmentRecipients(segID, apiKey string) ([]*Recipient, error) {
	response, err := SendGETRequest(apiKey, "/contactdb/segments/"+segID+"/recipients")
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
