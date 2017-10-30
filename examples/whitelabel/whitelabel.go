package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// CreateDomainWhitelabel : Create a domain whitelabel.
// POST /whitelabel/domains
func CreateDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "automatic_security": false,
  "custom_spf": true,
  "default": true,
  "domain": "example.com",
  "ips": [
    "192.168.1.1",
    "192.168.1.2"
  ],
  "subdomain": "news",
  "username": "john@example.com"
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

// ListDomainWhitelabels : List all domain whitelabels.
// GET /whitelabel/domains
func ListDomainWhitelabels() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	queryParams["domain"] = "test_string"
	queryParams["exclude_subusers"] = "true"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
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

// GetDefaultDomainWhitelabel : Get the default domain whitelabel.
// GET /whitelabel/domains/default
func GetDefaultDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/default", host)
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

// ListDomainWhitelabelsOfUser : List the domain whitelabels associated with the given user.
// GET /whitelabel/domains/subuser
func ListDomainWhitelabelsOfUser() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/subuser", host)
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

// DeleteDomainWhitelabelFromUser : Disassociate a domain whitelabel from a given user.
// DELETE /whitelabel/domains/subuser
func DeleteDomainWhitelabelFromUser() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/subuser", host)
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

// UpdateDomainWhitelabel : Update a domain whitelabel.
// PATCH /whitelabel/domains/{domain_id}
func UpdateDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{domain_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "custom_spf": true,
  "default": false
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

// GetDomainWhitelabel : Retrieve a domain whitelabel.
// GET /whitelabel/domains/{domain_id}
func GetDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{domain_id}", host)
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

// DeleteDomainWhitelabel : Delete a domain whitelabel.
// DELETE /whitelabel/domains/{domain_id}
func DeleteDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{domain_id}", host)
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

// AddDomainWhitelabelToUser : Associate a domain whitelabel with a given user.
// POST /whitelabel/domains/{domain_id}/subuser
func AddDomainWhitelabelToUser() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{domain_id}/subuser", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
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

// AddIPToDomainWhitelabel : Add an IP to a domain whitelabel.
// POST /whitelabel/domains/{id}/ips
func AddIPToDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{id}/ips", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "192.168.0.1"
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

// DeleteIPFromDomainWhitelabel : Remove an IP from a domain whitelabel.
// DELETE /whitelabel/domains/{id}/ips/{ip}
func DeleteIPFromDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{id}/ips/{ip}", host)
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

// ValidateDomainWhitelabel : Validate a domain whitelabel.
// POST /whitelabel/domains/{id}/validate
func ValidateDomainWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains/{id}/validate", host)
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

// CreateIPWhitelabel : Create an IP whitelabel
// POST /whitelabel/ips
func CreateIPWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/ips", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "domain": "example.com",
  "ip": "192.168.1.1",
  "subdomain": "email"
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

// GetIPWhitelabels : Retrieve all IP whitelabels
// GET /whitelabel/ips
func GetIPWhitelabels() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/ips", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["ip"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
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

// GetIPWhitelabel : Retrieve an IP whitelabel
// GET /whitelabel/ips/{id}
func GetIPWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/ips/{id}", host)
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

// DeleteIPWhitelabel : Delete an IP whitelabel
// DELETE /whitelabel/ips/{id}
func DeleteIPWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/ips/{id}", host)
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

// ValidateIPWhitelabel : Validate an IP whitelabel
// POST /whitelabel/ips/{id}/validate
func ValidateIPWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/ips/{id}/validate", host)
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

// CreateLinkWhitelabel : Create a Link Whitelabel
// POST /whitelabel/links
func CreateLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "default": true,
  "domain": "example.com",
  "subdomain": "mail"
}`)
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
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

// GetLinkWhitelabels : Retrieve all link whitelabels
// GET /whitelabel/links
func GetLinkWhitelabels() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links", host)
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

// GetDefaultLinkWhitelabel : Retrieve a Default Link Whitelabel
// GET /whitelabel/links/default
func GetDefaultLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/default", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["domain"] = "test_string"
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

// GetAssociatedLinkWhitelabel : Retrieve Associated Link Whitelabel
// GET /whitelabel/links/subuser
func GetAssociatedLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/subuser", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
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

// RemoveLinkWhitelabel : Disassociate a Link Whitelabel
// DELETE /whitelabel/links/subuser
func RemoveLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/subuser", host)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
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

// UpdateLinkWhitelabel : Update a Link Whitelabel
// PATCH /whitelabel/links/{id}
func UpdateLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/{id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "default": true
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

// GetLinkWhitelabel : Retrieve a Link Whitelabel
// GET /whitelabel/links/{id}
func GetLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/{id}", host)
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

// DeleteLinkWhitelabel : Delete a Link Whitelabel
// DELETE /whitelabel/links/{id}
func DeleteLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/{id}", host)
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

// ValidateLinkWhitelabel : Validate a Link Whitelabel
// POST /whitelabel/links/{id}/validate
func ValidateLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/{id}/validate", host)
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

// AddLinkWhitelabel : Associate a Link Whitelabel
// POST /whitelabel/links/{link_id}/subuser
func AddLinkWhitelabel() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/links/{link_id}/subuser", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
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

func main() {
	// add your function calls here
}
