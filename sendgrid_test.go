package sendgrid

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
	"time"
)

var (
	testAPIKey = "SENDGRID_APIKEY"
	testHost   = ""
	prismPath  = "./prism/" + runtime.GOOS + "/" + runtime.GOARCH
	prismArgs  = []string{"run", "-s", "https://raw.githubusercontent.com/sendgrid/sendgrid-oai/master/oai_stoplight.json"}
	prismCmd   *exec.Cmd
)

func TestMain(m *testing.M) {
	// By default prism runs on localhost:4010
	// Leanrn how to configure prism here: https://designer.stoplight.io/docs/prism
	testHost = "http://localhost:4010"

	prismPath += "/prism"
	if runtime.GOOS == "windows" {
		prismPath += ".exe"
	}

	prismCmd = exec.Command(prismPath, prismArgs...)

	// If you want to see prism's ouput uncomment below.
	prismReader, err := prismCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(prismReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("prism | %s\n", scanner.Text())
		}
	}()

	go func() {
		fmt.Println("Start Prism")
		err := prismCmd.Start()
		if err != nil {
			fmt.Println("Error starting prism", err)
		}
	}()

	// Need to give prism enough time to launch!
	duration := time.Second * 15
	time.Sleep(duration)

	exitCode := m.Run()
	if prismCmd != nil {
		prismCmd.Process.Kill()
		prismCmd = nil
	}

	os.Exit(exitCode)
}

func TestSendGridVersion(t *testing.T) {
	if Version != "3.0.0" {
		t.Error("SendGrid version does not match")
	}
}

func TestGetRequest(t *testing.T) {
	request := GetRequest("", "", "")
	if request.BaseURL != "https://api.sendgrid.com" {
		t.Error("Host default not set")
	}
	if request.Headers["Content-Type"] != "application/json" {
		t.Error("Wrong default content type")
	}
	if request.Headers["Authorization"] != "Bearer " {
		t.Error("Wrong default Authorization")
	}
	if request.Headers["User-Agent"] != "sendgrid/"+Version+";go" {
		t.Error("Wrong default User Agent")
	}

	request = GetRequest("API_KEY", "/v3/endpoint", "https://test.api.com")
	if request.BaseURL != "https://test.api.com/v3/endpoint" {
		t.Error("Host not set correctly")
	}
	if request.Headers["Content-Type"] != "application/json" {
		t.Error("Wrong content type")
	}
	if request.Headers["Authorization"] != "Bearer API_KEY" {
		t.Error("Wrong Authorization")
	}
	if request.Headers["User-Agent"] != "sendgrid/"+Version+";go" {
		t.Error("Wrong User Agent")
	}
}

func Test_test_access_settings_activity_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/activity", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_access_settings_whitelist_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/whitelist", testHost)
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
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_access_settings_whitelist_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/whitelist", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_access_settings_whitelist_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/whitelist", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "ids": [
    1,
    2,
    3
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_access_settings_whitelist__rule_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/whitelist/{rule_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_access_settings_whitelist__rule_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/access_settings/whitelist/{rule_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "My API Key",
  "scopes": [
    "mail.send",
    "alerts.create",
    "alerts.read"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys__api_key_id__put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys/{api_key_id}", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "name": "A New Hope",
  "scopes": [
    "user.profile.read",
    "user.profile.update"
  ]
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys__api_key_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys/{api_key_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "A New Hope"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys__api_key_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys/{api_key_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_api_keys__api_key_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/api_keys/{api_key_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "description": "A group description",
  "is_default": false,
  "name": "A group name"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "description": "Suggestions for items our users might like.",
  "id": 103,
  "name": "Item Suggestions"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__suppressions_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}/suppressions", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__suppressions_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}/suppressions", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_groups__group_id__suppressions__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/groups/{group_id}/suppressions/{email}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_suppressions_global_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/suppressions/global", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_suppressions_global__email__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/suppressions/global/{email}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_asm_suppressions_global__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/asm/suppressions/global/{email}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_browsers_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/browsers/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["browsers"] = "test_string"
	queryParams["limit"] = "test_string"
	queryParams["offset"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns", testHost)
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
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "0"
	queryParams["offset"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}", testHost)
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
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "send_at": 1489451436
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "send_at": 1489771528
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_now_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules/now", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_campaigns__campaign_id__schedules_test_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/campaigns/{campaign_id}/schedules/test", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "to": "your.email@example.com"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_categories_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/categories", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["category"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_categories_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/categories/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["categories"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_categories_stats_sums_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/categories/stats/sums", testHost)
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
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_clients_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/clients/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_clients__client_type__stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/clients/{client_type}/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_custom_fields_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/custom_fields", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "pet",
  "type": "text"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_custom_fields_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/custom_fields", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_custom_fields__custom_field_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/custom_fields/{custom_field_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_custom_fields__custom_field_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/custom_fields/{custom_field_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "202"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 202 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "your list name"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` [
  1,
  2,
  3,
  4
]`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "newlistname"
}`)
	queryParams := make(map[string]string)
	queryParams["list_id"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["list_id"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}", testHost)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "202"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 202 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__recipients_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}/recipients", testHost)
	request.Method = "POST"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__recipients_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}/recipients", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	queryParams["list_id"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__recipients__recipient_id__post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}/recipients/{recipient_id}", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_lists__list_id__recipients__recipient_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/lists/{list_id}/recipients/{recipient_id}", testHost)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["recipient_id"] = "0"
	queryParams["list_id"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` [
  {
    "email": "jones@example.com",
    "first_name": "Guy",
    "last_name": "Jones"
  }
]`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients", testHost)
	request.Method = "POST"
	request.Body = []byte(` [
  {
    "age": 25,
    "email": "example@example.com",
    "first_name": "",
    "last_name": "User"
  },
  {
    "age": 25,
    "email": "example2@example.com",
    "first_name": "Example",
    "last_name": "User"
  }
]`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_billable_count_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/billable_count", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_count_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/count", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients_search_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/search", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["{field_name}"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients__recipient_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/{recipient_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients__recipient_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/{recipient_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_recipients__recipient_id__lists_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/recipients/{recipient_id}/lists", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_reserved_fields_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/reserved_fields", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    },
    {
      "and_or": "and",
      "field": "last_clicked",
      "operator": "gt",
      "value": "01/02/2015"
    },
    {
      "and_or": "or",
      "field": "clicks.campaign_identifier",
      "operator": "eq",
      "value": "513"
    }
  ],
  "list_id": 4,
  "name": "Last Name Miller"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments__segment_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments/{segment_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    }
  ],
  "list_id": 5,
  "name": "The Millers"
}`)
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments__segment_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments/{segment_id}", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments__segment_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments/{segment_id}", testHost)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_contactdb_segments__segment_id__recipients_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/contactdb/segments/{segment_id}/recipients", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_devices_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/devices/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_geo_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/geo/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["country"] = "US"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["subuser"] = "test_string"
	queryParams["ip"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["exclude_whitelabels"] = "true"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_assigned_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/assigned", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "marketing"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools__pool_name__put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools/{pool_name}", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "name": "new_pool_name"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools__pool_name__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools/{pool_name}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools__pool_name__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools/{pool_name}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools__pool_name__ips_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools/{pool_name}/ips", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "0.0.0.0"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_pools__pool_name__ips__ip__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/pools/{pool_name}/ips/{ip}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_warmup_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/warmup", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "0.0.0.0"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_warmup_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/warmup", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_warmup__ip_address__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/warmup/{ip_address}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips_warmup__ip_address__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/warmup/{ip_address}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_ips__ip_address__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/ips/{ip_address}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_batch_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail/batch", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_batch__batch_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail/batch/{batch_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_send_beta_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail/send/beta", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "asm": {
    "group_id": 1,
    "groups_to_display": [
      1,
      2,
      3
    ]
  },
  "attachments": [
    {
      "content": "[BASE64 encoded content block here]",
      "content_id": "ii_139db99fdb5c3704",
      "disposition": "inline",
      "filename": "file1.jpg",
      "name": "file1",
      "type": "jpg"
    }
  ],
  "batch_id": "[YOUR BATCH ID GOES HERE]",
  "categories": [
    "category1",
    "category2"
  ],
  "content": [
    {
      "type": "text/html",
      "value": "<html><p>Hello, world!</p><img src=[CID GOES HERE]></img></html>"
    }
  ],
  "custom_args": {
    "New Argument 1": "New Value 1",
    "activationAttempt": "1",
    "customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
  },
  "from": {
    "email": "sam.smith@example.com",
    "name": "Sam Smith"
  },
  "headers": {},
  "ip_pool_name": "[YOUR POOL NAME GOES HERE]",
  "mail_settings": {
    "bcc": {
      "email": "ben.doe@example.com",
      "enable": true
    },
    "bypass_list_management": {
      "enable": true
    },
    "footer": {
      "enable": true,
      "html": "<p>Thanks</br>The SendGrid Team</p>",
      "text": "Thanks,/n The SendGrid Team"
    },
    "sandbox_mode": {
      "enable": false
    },
    "spam_check": {
      "enable": true,
      "post_to_url": "http://example.com/compliance",
      "threshold": 3
    }
  },
  "personalizations": [
    {
      "bcc": [
        {
          "email": "sam.doe@example.com",
          "name": "Sam Doe"
        }
      ],
      "cc": [
        {
          "email": "jane.doe@example.com",
          "name": "Jane Doe"
        }
      ],
      "custom_args": {
        "New Argument 1": "New Value 1",
        "activationAttempt": "1",
        "customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
      },
      "headers": {
        "X-Accept-Language": "en",
        "X-Mailer": "MyApp"
      },
      "send_at": 1409348513,
      "subject": "Hello, World!",
      "substitutions": {
        "sub": {
          "%name%": [
            "John",
            "Jane",
            "Sam"
          ]
        }
      },
      "to": [
        {
          "email": "john.doe@example.com",
          "name": "John Doe"
        }
      ]
    }
  ],
  "reply_to": {
    "email": "sam.smith@example.com",
    "name": "Sam Smith"
  },
  "sections": {
    "section": {
      ":sectionName1": "section 1 text",
      ":sectionName2": "section 2 text"
    }
  },
  "send_at": 1409348513,
  "subject": "Hello, World!",
  "template_id": "[YOUR TEMPLATE ID GOES HERE]",
  "tracking_settings": {
    "click_tracking": {
      "enable": true,
      "enable_text": true
    },
    "ganalytics": {
      "enable": true,
      "utm_campaign": "[NAME OF YOUR REFERRER SOURCE]",
      "utm_content": "[USE THIS SPACE TO DIFFERENTIATE YOUR EMAIL FROM ADS]",
      "utm_medium": "[NAME OF YOUR MARKETING MEDIUM e.g. email]",
      "utm_name": "[NAME OF YOUR CAMPAIGN]",
      "utm_term": "[IDENTIFY PAID KEYWORDS HERE]"
    },
    "open_tracking": {
      "enable": true,
      "substitution_tag": "%opentrack"
    },
    "subscription_tracking": {
      "enable": true,
      "html": "If you would like to unsubscribe and stop receiving these emails <% clickhere %>.",
      "substitution_tag": "<%click here%>",
      "text": "If you would like to unsubscribe and stop receiveing these emails <% click here %>."
    }
  }
}`)
	request.Headers["X-Mock"] = "202"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 202 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_address_whitelist_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/address_whitelist", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "list": [
    "email1@example.com",
    "example.com"
  ]
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_address_whitelist_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/address_whitelist", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_bcc_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/bcc", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "email@example.com",
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_bcc_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/bcc", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_bounce_purge_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/bounce_purge", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "hard_bounces": 5,
  "soft_bounces": 5
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_bounce_purge_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/bounce_purge", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_footer_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/footer", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "...",
  "plain_content": "..."
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_footer_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/footer", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_forward_bounce_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/forward_bounce", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "example@example.com",
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_forward_bounce_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/forward_bounce", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_forward_spam_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/forward_spam", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "",
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_forward_spam_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/forward_spam", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_plain_content_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/plain_content", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_plain_content_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/plain_content", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_spam_check_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/spam_check", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "max_score": 5,
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_spam_check_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/spam_check", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_template_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/template", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "<% body %>"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mail_settings_template_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mail_settings/template", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_mailbox_providers_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/mailbox_providers/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["mailbox_providers"] = "test_string"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_partner_settings_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/partner_settings", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_partner_settings_new_relic_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/partner_settings/new_relic", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enable_subuser_statistics": true,
  "enabled": true,
  "license_key": ""
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_partner_settings_new_relic_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/partner_settings/new_relic", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_scopes_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/scopes", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers", testHost)
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
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	queryParams["limit"] = "0"
	queryParams["offset"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_reputations_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/reputations", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["usernames"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["subusers"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_stats_monthly_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/stats/monthly", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["subuser"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers_stats_sums_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/stats/sums", testHost)
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
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "disabled": false
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__ips_put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/ips", testHost)
	request.Method = "PUT"
	request.Body = []byte(` [
  "127.0.0.1"
]`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__monitor_put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/monitor", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 500
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__monitor_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/monitor", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 50000
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__monitor_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/monitor", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__monitor_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/monitor", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_subusers__subuser_name__stats_monthly_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/subusers/{subuser_name}/stats/monthly", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
	queryParams["limit"] = "0"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_blocks_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/blocks", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_blocks_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/blocks", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_blocks__email__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/blocks/{email}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_blocks__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/blocks/{email}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_bounces_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/bounces", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "0"
	queryParams["end_time"] = "0"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_bounces_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/bounces", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": true,
  "emails": [
    "example@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_bounces__email__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/bounces/{email}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_bounces__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/bounces/{email}", testHost)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["email_address"] = "example@example.com"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_invalid_emails_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/invalid_emails", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_invalid_emails_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/invalid_emails", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_invalid_emails__email__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/invalid_emails/{email}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_invalid_emails__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/invalid_emails/{email}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_spam_report__email__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/spam_report/{email}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_spam_report__email__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/spam_report/{email}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_spam_reports_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/spam_reports", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_spam_reports_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/spam_reports", testHost)
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_suppression_unsubscribes_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/suppression/unsubscribes", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "example_name"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "new_example_name"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__versions_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}/versions", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "active": 1,
  "html_content": "<%body%>",
  "name": "example_version_name",
  "plain_content": "<%body%>",
  "subject": "<%subject%>",
  "template_id": "ddb96bbc-9b92-425e-8979-99464621b543"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__versions__version_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}/versions/{version_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "active": 1,
  "html_content": "<%body%>",
  "name": "updated_example_name",
  "plain_content": "<%body%>",
  "subject": "<%subject%>"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__versions__version_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}/versions/{version_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__versions__version_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}/versions/{version_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_templates__template_id__versions__version_id__activate_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/templates/{template_id}/versions/{version_id}/activate", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_click_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/click", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_click_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/click", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_google_analytics_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/google_analytics", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "utm_campaign": "website",
  "utm_content": "",
  "utm_medium": "email",
  "utm_source": "sendgrid.com",
  "utm_term": ""
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_google_analytics_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/google_analytics", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_open_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/open", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_open_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/open", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_subscription_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/subscription", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "html content",
  "landing": "landing page html",
  "plain_content": "text content",
  "replace": "replacement tag",
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_tracking_settings_subscription_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/tracking_settings/subscription", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_account_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/account", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_credits_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/credits", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_email_put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/email", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_email_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/email", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_password_put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/password", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "new_password": "new_password",
  "old_password": "old_password"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_profile_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/profile", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "city": "Orange",
  "first_name": "Example",
  "last_name": "User"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_profile_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/profile", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_scheduled_sends_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/scheduled_sends", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "batch_id": "YOUR_BATCH_ID",
  "status": "pause"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_scheduled_sends_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/scheduled_sends", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_scheduled_sends__batch_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/scheduled_sends/{batch_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "status": "pause"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_scheduled_sends__batch_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/scheduled_sends/{batch_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_scheduled_sends__batch_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/scheduled_sends/{batch_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_settings_enforced_tls_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/settings/enforced_tls", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "require_tls": true,
  "require_valid_cert": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_settings_enforced_tls_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/settings/enforced_tls", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_username_put(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/username", testHost)
	request.Method = "PUT"
	request.Body = []byte(` {
  "username": "test_username"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_username_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/username", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_webhooks_event_settings_patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/webhooks/event/settings", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "bounce": true,
  "click": true,
  "deferred": true,
  "delivered": true,
  "dropped": true,
  "enabled": true,
  "group_resubscribe": true,
  "group_unsubscribe": true,
  "open": true,
  "processed": true,
  "spam_report": true,
  "unsubscribe": true,
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_webhooks_event_settings_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/webhooks/event/settings", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_webhooks_event_test_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/webhooks/event/test", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "url": "url"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_webhooks_parse_settings_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/webhooks/parse/settings", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_user_webhooks_parse_stats_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/user/webhooks/parse/stats", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains", testHost)
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
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	queryParams["domain"] = "test_string"
	queryParams["exclude_subusers"] = "true"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains_default_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/default", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains_subuser_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/subuser", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains_subuser_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/subuser", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__domain_id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{domain_id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "custom_spf": true,
  "default": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__domain_id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{domain_id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__domain_id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{domain_id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__domain_id__subuser_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{domain_id}/subuser", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__id__ips_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{id}/ips", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "192.168.0.1"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__id__ips__ip__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{id}/ips/{ip}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_domains__id__validate_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/domains/{id}/validate", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_ips_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/ips", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "domain": "example.com",
  "ip": "192.168.1.1",
  "subdomain": "email"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_ips_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/ips", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["ip"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_ips__id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/ips/{id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_ips__id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/ips/{id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_ips__id__validate_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/ips/{id}/validate", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links", testHost)
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
	request.Headers["X-Mock"] = "201"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 201 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links_default_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/default", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["domain"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links_subuser_get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/subuser", testHost)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links_subuser_delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/subuser", testHost)
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links__id__patch(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/{id}", testHost)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "default": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links__id__get(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/{id}", testHost)
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links__id__delete(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/{id}", testHost)
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 204 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links__id__validate_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/{id}/validate", testHost)
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}

func Test_test_whitelabel_links__link_id__subuser_post(t *testing.T) {
	request := GetRequest(testAPIKey, "/v3/whitelabel/links/{link_id}/subuser", testHost)
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := API(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != 200 {
		t.Error("Wrong status code returned")
	}
}
