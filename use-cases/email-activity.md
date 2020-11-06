## Email activity filter
### To find the email sent to `testing@sendgrid.net`
```go
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func main() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/messages", host)
	request.Method = "GET"

	filterKey := "to_email"
	filterOperator := url.QueryEscape("=")
	filterValue := "testing@sendgrid.net"
	filterValue = url.QueryEscape(fmt.Sprintf("\"%s\"", filterValue))

	queryParams := make(map[string]string)
	queryParams["query"] = fmt.Sprintf("%s%s%s", filterKey, filterOperator, filterValue)
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
```
