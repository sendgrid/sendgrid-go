package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// Createadomainwhitelabel : Create a domain whitelabel.
// POST /whitelabel/domains
func Createadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Listalldomainwhitelabels : List all domain whitelabels.
// GET /whitelabel/domains
func Listalldomainwhitelabels() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Getthedefaultdomainwhitelabel : Get the default domain whitelabel.
// GET /whitelabel/domains/default
func Getthedefaultdomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Listthedomainwhitelabelassociatedwiththegivenuser : List the domain whitelabel associated with the given user.
// GET /whitelabel/domains/subuser
func Listthedomainwhitelabelassociatedwiththegivenuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Disassociateadomainwhitelabelfromagivenuser : Disassociate a domain whitelabel from a given user.
// DELETE /whitelabel/domains/subuser
func Disassociateadomainwhitelabelfromagivenuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Updateadomainwhitelabel : Update a domain whitelabel.
// PATCH /whitelabel/domains/{domain_id}
func Updateadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Retrieveadomainwhitelabel : Retrieve a domain whitelabel.
// GET /whitelabel/domains/{domain_id}
func Retrieveadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Deleteadomainwhitelabel : Delete a domain whitelabel.
// DELETE /whitelabel/domains/{domain_id}
func Deleteadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Associateadomainwhitelabelwithagivenuser : Associate a domain whitelabel with a given user.
// POST /whitelabel/domains/{domain_id}/subuser
func Associateadomainwhitelabelwithagivenuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// AddanIPtoadomainwhitelabel : Add an IP to a domain whitelabel.
// POST /whitelabel/domains/{id}/ips
func AddanIPtoadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RemoveanIPfromadomainwhitelabel : Remove an IP from a domain whitelabel.
// DELETE /whitelabel/domains/{id}/ips/{ip}
func RemoveanIPfromadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Validateadomainwhitelabel : Validate a domain whitelabel.
// POST /whitelabel/domains/{id}/validate
func Validateadomainwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// CreateanIPwhitelabel : Create an IP whitelabel
// POST /whitelabel/ips
func CreateanIPwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RetrieveallIPwhitelabels : Retrieve all IP whitelabels
// GET /whitelabel/ips
func RetrieveallIPwhitelabels() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RetrieveanIPwhitelabel : Retrieve an IP whitelabel
// GET /whitelabel/ips/{id}
func RetrieveanIPwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// DeleteanIPwhitelabel : Delete an IP whitelabel
// DELETE /whitelabel/ips/{id}
func DeleteanIPwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// ValidateanIPwhitelabel : Validate an IP whitelabel
// POST /whitelabel/ips/{id}/validate
func ValidateanIPwhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// CreateaLinkWhitelabel : Create a Link Whitelabel
// POST /whitelabel/links
func CreateaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// Retrievealllinkwhitelabels : Retrieve all link whitelabels
// GET /whitelabel/links
func Retrievealllinkwhitelabels() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RetrieveaDefaultLinkWhitelabel : Retrieve a Default Link Whitelabel
// GET /whitelabel/links/default
func RetrieveaDefaultLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RetrieveAssociatedLinkWhitelabel : Retrieve Associated Link Whitelabel
// GET /whitelabel/links/subuser
func RetrieveAssociatedLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// DisassociateaLinkWhitelabel : Disassociate a Link Whitelabel
// DELETE /whitelabel/links/subuser
func DisassociateaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// UpdateaLinkWhitelabel : Update a Link Whitelabel
// PATCH /whitelabel/links/{id}
func UpdateaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// RetrieveaLinkWhitelabel : Retrieve a Link Whitelabel
// GET /whitelabel/links/{id}
func RetrieveaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// DeleteaLinkWhitelabel : Delete a Link Whitelabel
// DELETE /whitelabel/links/{id}
func DeleteaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// ValidateaLinkWhitelabel : Validate a Link Whitelabel
// POST /whitelabel/links/{id}/validate
func ValidateaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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

// AssociateaLinkWhitelabel : Associate a Link Whitelabel
// POST /whitelabel/links/{link_id}/subuser
func AssociateaLinkWhitelabel() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
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
