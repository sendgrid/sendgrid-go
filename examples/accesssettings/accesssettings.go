package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// RetrieveAllRecentAccessAttempts Retrieve all recent access attempts
// GET /access_settings/activity
func RetrieveAllRecentAccessAttempts() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/activity", host)
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

// AddIPsToWhitelist Add one or more IPs to the whitelist
// POST /access_settings/whitelist
func AddIPsToWhitelist() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/whitelist", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "ips": [
    {
      "ip": "192.168.1.1"
    },
    {
      "ip": "192.*.*.*"
    },
    {
      "ip": "192.168.1.3/32"
    }
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

// RetrieveCurrentWhitelistedIPs Retrieve a list of currently whitelisted IPs
// GET /access_settings/whitelist
func RetrieveCurrentWhitelistedIPs() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/whitelist", host)
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

// RemoveIPsFromWhitelist Remove one or more IPs from the whitelist
// DELETE /access_settings/whitelist
func RemoveIPsFromWhitelist() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/whitelist", host)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "ids": [
    1,
    2,
    3
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

// RetrieveWhitelistedIP Retrieve a specific whitelisted IP
// GET /access_settings/whitelist/{rule_id}
func RetrieveWhitelistedIP() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/whitelist/{rule_id}", host)
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

// RemoveIPfromWhitelist  Remove a specific IP from the whitelist
// DELETE /access_settings/whitelist/{rule_id}
func RemoveIPfromWhitelist() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/access_settings/whitelist/{rule_id}", host)
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

func main() {
	// add your function calls here
}
