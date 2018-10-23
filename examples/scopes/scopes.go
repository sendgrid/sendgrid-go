package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// GetAllowedScopesForUser : Retrieve a list of scopes for which this user has access.
// GET /scopes
func GetAllowedScopesForUser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/scopes", host)
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
