**This helper is a stand-alone module to help get you started consuming and processing Inbound Parse data.**

## Table of Contents

* [Fields](#fields)
* [Example Usage](#example-usage)
* [Testing the Source Code](#testing)
* [Contributing](#contributing)

# Fields

### parsedEmail.Envelope
  parsedEmail.Envelope.To and parsedEmail.Envelope.From represent the exact email addresses that the email was sent to and the exact email address of the sender. There are no special characters and these fields are safe to use without further parsing as email addresses 

### parsedEmail.ParsedValues 
  Please see [Send Grid Docs](https://docs.sendgrid.com/for-developers/parsing-email/setting-up-the-inbound-parse-webhook) to see what fields are available and preparsed by SendGrid. Use these fields over the Headers as they are parsed by SendGrid and gauranteed to be consistent

### parsedEmail.TextBody
  this field will satisfy  most cases. SendGrid pre-parses the body into a plain text string separated with \n 

### parsedEmail.Body and parsedEmail.Attachments
  are populated *only* when the raw option is checked in the SendGrid Dashboard. However unless you need the raw HTML body, it is not necessary. The fields are named as they are for backward compatability 

### parsedEmail.Headers
  these may change depending on the email client and are not pre-parsed by SendGrid, use carefully


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
	parsedEmail, err := ParseWithAttachments(request)
	if err != nil {
		log.Fatal(err)
	}
    
	fmt.Print(parsedEmail.Envelope.From)
	
	for filename, contents := range parsedEmail.ParsedAttachments {
		// Do something with an attachment
		handleAttachment(filename, contents)
	}
    
  emailLines := strings.Split(parsedEmail.TextBody, "\n")
	for section, body := range emailLines {
		// Do something with the email lines
	}

    
	// Twilio SendGrid needs a 200 OK response to stop POSTing
	response.WriteHeader(http.StatusOK)
}

// example of uploading an attachment to s3 using the Go sdk-2
func handleAttachment(parsedEmail *ParsedEmail) {
	for _, contents := range parsedEmail.ParsedAttachments {
			if _, err := sgh.Client.Upload(ctx, &s3.PutObjectInput{
				Bucket:               &bucket,
				Key:                  &uploadPath,
				Body:                 contents.File,
				ContentType:          aws.String(contents.ContentType),
      }
  }
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
