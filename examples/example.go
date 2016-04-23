package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

func main() {

	// GET collection
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys", "https://api.sendgrid.com", "v3")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "100"
	queryParams["offset"] = "0"
	request.QueryParams = queryParams
	request.RequestHeaders["X-Test"] = "true"

	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}

	// POST
	request = sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys", "https://api.sendgrid.com", "v3")
	request.Method = "POST"
	var requestBody = []byte(` {
        "name": "My API Key",
        "scopes": [
            "mail.send",
            "alerts.create",
            "alerts.read"
        ]
    }`)
	request.RequestBody = requestBody
	response, err = sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}

	// Get a particular return value.
	// Note that you can unmarshall into a struct if
	// you know the JSON structure in advance.
	b := []byte(response.ResponseBody)
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	m := f.(map[string]interface{})
	apiKey := m["api_key_id"].(string)

	// GET Single
	request = sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys/"+apiKey, "https://api.sendgrid.com", "v3")
	request.Method = "GET"
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}

	// PATCH
	request = sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys/"+apiKey, "https://api.sendgrid.com", "v3")
	request.Method = "PATCH"
	requestBody = []byte(`{
        "name": "A New Hope"
    }`)
	request.RequestBody = requestBody
	response, err = sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}

	// PUT
	request = sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys/"+apiKey, "https://api.sendgrid.com", "v3")
	request.Method = "PUT"
	requestBody = []byte(`{
        "name": "A New Hope",
        "scopes": [
            "user.profile.read",
            "user.profile.update"
        ]
    }`)
	request.RequestBody = requestBody
	response, err = sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}

	// DELETE
	request = sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys/"+apiKey, "https://api.sendgrid.com", "v3")
	request.Method = "DELETE"
	response, err = sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseHeaders)
	}

}
