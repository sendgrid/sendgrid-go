package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// CreateSubuser : Creates a Subuser
// POST /subusers
func CreateSubuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "email": "John@example.com",
  "ips": [
    "1.1.1.1",
    "2.2.2.2"
  ],
  "password": "johns_password",
  "username": "John@example.com"
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

// ListSubusers : List all Subusers
// GET /subusers
func ListSubusers() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
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

// GetSubuserReputations : Retrieve Subuser Reputations
// GET /subusers/reputations
func GetSubuserReputations() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/reputations", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["usernames"] = "test_string"
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

// GetEmailStatsBySubusers : Retrieve email statistics for your subusers.
// GET /subusers/stats
func GetEmailStatsBySubusers() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["subusers"] = "test_string"
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

// GetMonthlyStatsOfAllSubusers : Retrieve monthly stats for all subusers
// GET /subusers/stats/monthly
func GetMonthlyStatsOfAllSubusers() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/stats/monthly", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["subuser"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
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

// GetSumsOfStatsOfSubusers : Retrieve the totals for each email statistic metric for all subusers.
// GET /subusers/stats/sums
func GetSumsOfStatsOfSubusers() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/stats/sums", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["sort_by_direction"] = "asc"
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

// ToggleSubuser : Enable/disable a subuser
// PATCH /subusers/{subuser_name}
func ToggleSubuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "disabled": false
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

// DeleteSubuser : Delete a subuser
// DELETE /subusers/{subuser_name}
func DeleteSubuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}", host)
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

// UpdateSubuserAssignedIPs : Update IPs assigned to a subuser
// PUT /subusers/{subuser_name}/ips
func UpdateSubuserAssignedIPs() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/ips", host)
	request.Method = "PUT"
	request.Body = []byte(` [
  "127.0.0.1"
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

// UpdateSubuserMonitorSettings : Update Monitor Settings for a subuser
// PUT /subusers/{subuser_name}/monitor
func UpdateSubuserMonitorSettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/monitor", host)
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 500
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

// CreateMonitorSettings : Create monitor settings
// POST /subusers/{subuser_name}/monitor
func CreateMonitorSettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/monitor", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 50000
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

// GetMonitorSettingsForSubuser : Retrieve monitor settings for a subuser
// GET /subusers/{subuser_name}/monitor
func GetMonitorSettingsForSubuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/monitor", host)
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

// DeleteMonitorSettings : Delete monitor settings
// DELETE /subusers/{subuser_name}/monitor
func DeleteMonitorSettings() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/monitor", host)
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

// RetrieveMonthlyEmailStatsForSubuser : Retrieve the monthly email statistics for a single subuser
// GET /subusers/{subuser_name}/stats/monthly
func RetrieveMonthlyEmailStatsForSubuser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/subusers/{subuser_name}/stats/monthly", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
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

func main() {
	// add your function calls here
}
