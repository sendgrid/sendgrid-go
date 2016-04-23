package main

```go
import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

func main() {

	// GET collection
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/api_keys", "https://api.sendgrid.com", "v3")
	request.Method = "GET" // or "POST", "PATCH", "PUT", "DELETE"
	// optional query parameters
    queryParams := make(map[string]string) 
	queryParams["limit"] = "100"
	queryParams["offset"] = "0"
	request.QueryParams = queryParams
	// optional request headers
    request.RequestHeaders["X-Test"] = "true"

	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}
}
```