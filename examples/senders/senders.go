package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// CreateaSenderIdentity Creates a Sender Identity
// POST /senders
func CreateaSenderIdentity() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "address": "123 Elm St.",
  "address_2": "Apt. 456",
  "city": "Denver",
  "country": "United States",
  "from": {
    "email": "from@example.com",
    "name": "Example INC"
  },
  "nickname": "My Sender ID",
  "reply_to": {
    "email": "replyto@example.com",
    "name": "Example INC"
  },
  "state": "Colorado",
  "zip": "80202"
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

// GetallSenderIdentities Gets all Sender Identities
// GET /senders
func GetallSenderIdentities() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders", host)
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

// UpdateaSenderIdentity Updates a Sender Identity
// PATCH /senders/{sender_id}
func UpdateaSenderIdentity() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders/{sender_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "address": "123 Elm St.",
  "address_2": "Apt. 456",
  "city": "Denver",
  "country": "United States",
  "from": {
    "email": "from@example.com",
    "name": "Example INC"
  },
  "nickname": "My Sender ID",
  "reply_to": {
    "email": "replyto@example.com",
    "name": "Example INC"
  },
  "state": "Colorado",
  "zip": "80202"
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

// ViewaSenderIdentity Views a Sender Identity
// GET /senders/{sender_id}
func ViewaSenderIdentity() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders/{sender_id}", host)
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

// DeleteaSenderIdentity Deletes a Sender Identity
// DELETE /senders/{sender_id}
func DeleteaSenderIdentity() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders/{sender_id}", host)
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

// ResendSenderIdentityVerification Resend Sender Identity Verification
// POST /senders/{sender_id}/resend_verification
func ResendSenderIdentityVerification() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/senders/{sender_id}/resend_verification", host)
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

func main() {
	// add your function calls here
}
