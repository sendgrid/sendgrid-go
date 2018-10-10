package contactdb

// Segment holds the struct for the segments
type Segment struct {
	ID             int          `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	ListID         int          `json:"list_id,omitempty"`
	Conditions     []*Condition `json:"conditions,omitempty"`
	AndOr          string       `json:"and_or"` // No omitempty because an empty string is a valid value.
	RecipientCount int          `json:"recipient_count,omitempty"`
}

// Condition is a condition
type Condition struct {
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
	return nil, nil
}

// GetSegment will return a specific segment by ID. (GET)
func GetSegment(segID int, apiKey string) (*Segment, error) {
	return nil, nil
}

// GetAllSegments will return all of your segments. (GET)
func GetAllSegments(apiKey string) ([]*Segment, error) {
	return nil, nil
}

// UpdateSegment will update a specific segment with new values. (PATCH)
func UpdateSegment(segID int, segment *Segment, apiKey string) (*Segment, error) {
	return nil, nil
}

// DeleteSegment will delete a specific segment. (DELETE)
func DeleteSegment(segID int, apiKey string) error {
	return nil
}

// GetSegmentRecipients will return all of the recipients of a segment. (GET)
func GetSegmentRecipients(segID, apiKey string) ([]*Recipient, error) {
	return nil, nil
}
