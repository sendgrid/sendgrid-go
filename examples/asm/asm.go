package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// CreateSuppressionGroup : Create a new suppression group
// POST /asm/groups
func CreateSuppressionGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "description": "Suggestions for products our users might like.",
  "is_default": true,
  "name": "Product Suggestions"
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

// GetSuppressionGroupsInfo : Retrieve information about multiple suppression groups
// GET /asm/groups
func GetSuppressionGroupsInfo() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["id"] = "1"
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

// UpdateSuppressionGroup : Update a suppression group.
// PATCH /asm/groups/{group_id}
func UpdateSuppressionGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "description": "Suggestions for items our users might like.",
  "id": 103,
  "name": "Item Suggestions"
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

// GetSuppressionGroupInfo : Get information on a single suppression group.
// GET /asm/groups/{group_id}
func GetSuppressionGroupInfo() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
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

// DeleteSuppressionGroup : Delete a suppression group.
// DELETE /asm/groups/{group_id}
func DeleteSuppressionGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
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

// AddSuppressionsToGroup : Add suppressions to a suppression group
// POST /asm/groups/{group_id}/suppressions
func AddSuppressionsToGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
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

// GetSuppressionsFromGroup : Retrieve all suppressions for a suppression group
// GET /asm/groups/{group_id}/suppressions
func GetSuppressionsFromGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions", host)
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

// SearchSuppressionsInGroup : Search for suppressions within a group
// POST /asm/groups/{group_id}/suppressions/search
func SearchSuppressionsInGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions/search", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "exists1@example.com",
    "exists2@example.com",
    "doesnotexists@example.com"
  ]
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

// DeleteSuppressionFromGroup Delete a suppression from a suppression group
// DELETE /asm/groups/{group_id}/suppressions/{email}
func DeleteSuppressionFromGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions/{email}", host)
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

// GetAllSuppressions : Retrieve all suppressions
// GET /asm/suppressions
func GetAllSuppressions() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions", host)
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

// AddRecipientAddressesToGlobalSuppressionGroup : Add recipient addresses to the global suppression group.
// POST /asm/suppressions/global
func AddRecipientAddressesToGlobalSuppressionGroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
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

// GetGlobalSuppression : Retrieve a Global Suppression
// GET /asm/suppressions/global/{email}
func GetGlobalSuppression() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global/{email}", host)
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

// DeleteGlobalSuppression : Delete a Global Suppression
// DELETE /asm/suppressions/global/{email}
func DeleteGlobalSuppression() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global/{email}", host)
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

// GetSuppressionGroupsByEmail : Retrieve all suppression groups for an email address
// GET /asm/suppressions/{email}
func GetSuppressionGroupsByEmail() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/{email}", host)
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

func main() {
	// add your function calls here
}
