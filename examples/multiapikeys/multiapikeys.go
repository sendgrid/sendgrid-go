package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
)

// CreateAPIMultikeys : Create Multiple API keys
// POST /api_keys
func CreateAPIMultikeys() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys", host)

	for _, request := range requests {
		request.Method = "POST"
		request.Body = []byte(` {
		"name": "My API Key", 
		"sample": "data", 
		"scopes": [
			"mail.send", 
			"alerts.create", 
			"alerts.read"
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
}

// RetrieveallMultiAPIKeysbelongingtotheauthenticateduser : Retrieve all Multi API Keys belonging to the authenticated user
// GET /api_keys
func RetrieveallMultiAPIKeysbelongingtotheauthenticateduser() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys", host)

	for _, request := range requests {
		request.Method = "GET"
		queryParams := make(map[string]string)
		queryParams["limit"] = "1"
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
}

// UpdatethenamescopesofanMultiAPIKey : Update the name & scopes of Multi  API Key
// PUT /api_keys/{api_key_id}
func UpdatethenamescopesofanMultiAPIKey() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys/{api_key_id}", host)

	for _, request := range requests {
		request.Method = "PUT"
		request.Body = []byte(` {
								"name": "A New Hope", 
								"scopes": [
									"user.profile.read", 
									"user.profile.update"
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
}

// UpdateMultiAPIkeys : Update Multi API keys
// PATCH /api_keys/{api_key_id}
func UpdateMultiAPIkeys() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys/{api_key_id}", host)

	for _, request := range requests {
		request.Method = "PATCH"
		request.Body = []byte(` {
									"name": "A New Hope"
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
}

// RetrieveanexistingMultiAPIKey : Retrieve an existing Multi API Key
// GET /api_keys/{api_key_id}
func RetrieveanexistingMultiAPIKey() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys/{api_key_id}", host)

	for _, request := range requests {
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
}

// DeleteMultiAPIkeys : Delete Multi API keys
// DELETE /api_keys/{api_key_id}
func DeleteMultiAPIkeys() {
	apiKey := strings.Split(os.Getenv("YOUR_SENDGRID_MULTI_APIKEY"), "|")
	host := "https://api.sendgrid.com"
	requests := sendgrid.GetRequests(apiKey, "/v3/api_keys/{api_key_id}", host)

	for _, request := range requests {
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
}

func main() {
	arrString := []string{"a", "b", "c"}
	for _, str := range arrString {
		fmt.Println(str)
	}
}
