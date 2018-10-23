package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// CreateCustomField : Create a Custom Field
// POST /contactdb/custom_fields
func CreateCustomField() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/custom_fields", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "pet",
  "type": "text"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetAllCustomFields : Retrieve all custom fields
// GET /contactdb/custom_fields
func GetAllCustomFields() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/custom_fields", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetCustomField : Retrieve a Custom Field
// GET /contactdb/custom_fields/{custom_field_id}
func GetCustomField() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/custom_fields/{custom_field_id}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteCustomField : Delete a Custom Field
// DELETE /contactdb/custom_fields/{custom_field_id}
func DeleteCustomField() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/custom_fields/{custom_field_id}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// CreateList : Create a List
// POST /contactdb/lists
func CreateList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "your list name"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetAllLists : Retrieve all lists
// GET /contactdb/lists
func GetAllLists() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteLists : Delete Multiple lists
// DELETE /contactdb/lists
func DeleteLists() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists", host)
	request.Method = "DELETE"
	request.Body = []byte(` [
  1,
  2,
  3,
  4
]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateList : Update a List
// PATCH /contactdb/lists/{list_id}
func UpdateList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "newlistname"
}`)
	queryParams := make(map[string]string)
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetList : Retrieve a single list
// GET /contactdb/lists/{list_id}
func GetList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteList : Delete a List
// DELETE /contactdb/lists/{list_id}
func DeleteList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}", host)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// AddRecipientsToList : Add Multiple Recipients to a List
// POST /contactdb/lists/{list_id}/recipients
func AddRecipientsToList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}/recipients", host)
	request.Method = "POST"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetAllRecipientsByList : Retrieve all recipients on a List
// GET /contactdb/lists/{list_id}/recipients
func GetAllRecipientsByList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}/recipients", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// AddRecipientToList : Add a Single Recipient to a List
// POST /contactdb/lists/{list_id}/recipients/{recipient_id}
func AddRecipientToList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}/recipients/{recipient_id}", host)
	request.Method = "POST"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteRecipientfromList : Delete a Single Recipient from a Single List
// DELETE /contactdb/lists/{list_id}/recipients/{recipient_id}
func DeleteRecipientfromList() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/lists/{list_id}/recipients/{recipient_id}", host)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["recipient_id"] = "1"
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateRecipient : Update a Recipient
// PATCH /contactdb/recipients
func UpdateRecipient() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients", host)
	request.Method = "PATCH"
	request.Body = []byte(` [
  {
    "email": "jones@example.com",
    "first_name": "Guy",
    "last_name": "Jones"
  }
]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// AddRecipients : Add recipients
// POST /contactdb/recipients
func AddRecipients() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients", host)
	request.Method = "POST"
	request.Body = []byte(` [
  {
    "age": 25,
    "email": "example@example.com",
    "first_name": "",
    "last_name": "User"
  },
  {
    "age": 25,
    "email": "example2@example.com",
    "first_name": "Example",
    "last_name": "User"
  }
]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetRecipients : Retrieve recipients
// GET /contactdb/recipients
func GetRecipients() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteRecipients : Delete multiple recipients
// DELETE /contactdb/recipients
func DeleteRecipients() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients", host)
	request.Method = "DELETE"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetBillableRecipientsCount : Retrieve the count of billable recipients
// GET /contactdb/recipients/billable_count
func GetBillableRecipientsCount() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/billable_count", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetRecipientsCount : Retrieve the count of recipients
// GET /contactdb/recipients/count
func GetRecipientsCount() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/count", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// SearchRecipients : Retrieve recipients matching search criteria
// GET /contactdb/recipients/search
func SearchRecipients() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/search", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["{field_name}"] = "test_string"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetRecipient : Retrieve a single recipient
// GET /contactdb/recipients/{recipient_id}
func GetRecipient() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/{recipient_id}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteRecipient : Delete a Recipient
// DELETE /contactdb/recipients/{recipient_id}
func DeleteRecipient() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/{recipient_id}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetListsByRecipient : Retrieve the lists that a recipient is on
// GET /contactdb/recipients/{recipient_id}/lists
func GetListsByRecipient() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients/{recipient_id}/lists", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetReservedFields : Retrieve reserved fields
// GET /contactdb/reserved_fields
func GetReservedFields() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/reserved_fields", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// CreateSegment : Create a Segment
// POST /contactdb/segments
func CreateSegment() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    },
    {
      "and_or": "and",
      "field": "last_clicked",
      "operator": "gt",
      "value": "01/02/2015"
    },
    {
      "and_or": "or",
      "field": "clicks.campaign_identifier",
      "operator": "eq",
      "value": "513"
    }
  ],
  "list_id": 4,
  "name": "Last Name Miller"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetAllSegments : Retrieve all segments
// GET /contactdb/segments
func GetAllSegments() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateSegment : Update a segment
// PATCH /contactdb/segments/{segment_id}
func UpdateSegment() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments/{segment_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    }
  ],
  "list_id": 5,
  "name": "The Millers"
}`)
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "test_string"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetSegment : Retrieve a segment
// GET /contactdb/segments/{segment_id}
func GetSegment() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments/{segment_id}", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteSegment : Delete a segment
// DELETE /contactdb/segments/{segment_id}
func DeleteSegment() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments/{segment_id}", host)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetRecipientsBySegment : Retrieve recipients on a segment
// GET /contactdb/segments/{segment_id}/recipients
func GetRecipientsBySegment() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/segments/{segment_id}/recipients", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// add your function calls here
}
