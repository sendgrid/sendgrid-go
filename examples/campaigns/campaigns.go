package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// CreateCampaign Create a Campaign
// POST /campaigns
func CreateCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "categories": [
    "spring line"
  ],
  "custom_unsubscribe_url": "",
  "html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
  "ip_pool": "marketing",
  "list_ids": [
    110,
    124
  ],
  "plain_content": "Check out our spring line!",
  "segment_ids": [
    110
  ],
  "sender_id": 124451,
  "subject": "New Products for Spring!",
  "suppression_group_id": 42,
  "title": "March Newsletter"
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

// GetAllCampaigns Retrieve all Campaigns
// GET /campaigns
func GetAllCampaigns() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns", host)
	request.Method = "GET"
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

// UpdateCampaign Update a Campaign
// PATCH /campaigns/{campaign_id}
func UpdateCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "categories": [
    "summer line"
  ],
  "html_content": "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>",
  "plain_content": "Check out our summer line!",
  "subject": "New Products for Summer!",
  "title": "May Newsletter"
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

// GetCampaign Retrieve a single campaign by id
// GET /campaigns/{campaign_id}
func GetCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}", host)
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

// DeleteCampaign Delete a Campaign by id
// DELETE /campaigns/{campaign_id}
func DeleteCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}", host)
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

// UpdateScheduledCampaign Update a Scheduled Campaign
// PATCH /campaigns/{campaign_id}/schedules
func UpdateScheduledCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "send_at": 1489451436
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

// ScheduleCampaign Schedule a Campaign
// POST /campaigns/{campaign_id}/schedules
func ScheduleCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "send_at": 1489771528
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

// GetCampaignSchedules View Scheduled Time of a Campaign
// GET /campaigns/{campaign_id}/schedules
func GetCampaignSchedules() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules", host)
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

// UnscheduleScheduledCampaign Unschedule a Scheduled Campaign
// DELETE /campaigns/{campaign_id}/schedules
func UnscheduleScheduledCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules", host)
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

// SendCampaign Send a Campaign
// POST /campaigns/{campaign_id}/schedules/now
func SendCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules/now", host)
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

// SendTestCampaign Send a Test Campaign
// POST /campaigns/{campaign_id}/schedules/test
func SendTestCampaign() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/campaigns/{campaign_id}/schedules/test", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "to": "your.email@example.com"
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
