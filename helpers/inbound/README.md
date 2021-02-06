**This helper is a stand-alone module to help get you started consuming and processing Inbound Parse data.**

## Table of Contents

* [Example Usage](#example-usage)
* [Testing the Source Code](#testing)
* [Contributing](#contributing)

# Example Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go/helpers/inbound"
)

func inboundHandler(response http.ResponseWriter, request *http.Request) {
	parsedEmail, err := Parse(request)
	if err != nil {
		log.Fatal(err)
	}
    
	fmt.Print(parsedEmail.Headers["From"])
	
	for filename, contents := range parsedEmail.Attachments {
		// Do something with an attachment
		handleAttachment(filename, contents)
	}
    
	for section, body := range parsedEmail.Body {
		// Do something with the email body
		handleEmail(body)
	}
    
	// Twilio SendGrid needs a 200 OK response to stop POSTing
	response.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/inbound", inboundHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("Error")
	}
}
```

<a name="testing"></a>
# Testing the Source Code

Tests are located in the `helpers/inbound` folder:

- [inbound_test.go](inbound_test.go)

Learn about testing this code [here](../../CONTRIBUTING.md#testing).

<a name="contributing"></a>
# Contributing

If you would like to contribute to this project, please see our [contributing guide](../../CONTRIBUTING.md). Thanks!
