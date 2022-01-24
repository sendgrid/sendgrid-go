**This helper is a stand-alone module to help get you started consuming and processing Inbound Parse data.**

## Table of Contents

* [Fields](#fields)
* [Example Usage](#example-usage)
* [Testing the Source Code](#testing)
* [Contributing](#contributing)

# Fields

### ParsedEmail.Envelope
  ParsedEmail.Envelope.To and ParsedEmail.Envelope.From represent the exact email addresses that the email was sent to and the exact email address of the sender. There are no special characters and these fields are safe to use without further parsing as email addresses

### ParsedEmail.ParsedValues
  Please see [SendGrid Docs](https://docs.sendgrid.com/for-developers/parsing-email/setting-up-the-inbound-parse-webhook) to see what fields are available and preparsed by SendGrid. Use these fields over the Headers as they are parsed by SendGrid and gauranteed to be consistent

### ParsedEmail.TextBody
  this field will satisfy  most cases. SendGrid pre-parses the body into a plain text string separated with \n

### ParsedEmail.ParsedAttachments
  populated **only** when processing the email with ParseWithAttachments(). Provides the following ease of use values
  - File: full attachment for uploading or processing (see example to upload to s3)
  - Size: file size, useful for filtering or setting upper limits to attachments
  - Filename: copies the original filename of the attachment, if there is not one, it defaults to 'Untitled'
  - ContentType: the type of file

### ParsedEmail.Body
  populated *only* when the raw option is checked in the SendGrid Dashboard. Provides the raw HTML body of the email, unless you need to record the exact unparsed HTML payload from the email client, you should use the parsed fields instead. The field is named Body for backward compatability

### ParsedEmail.Attachments
  populated *only* when the raw option is checked in the SendGrid Dashboard. This field is deprecated. Use ParsedAttachments instead which does not require the Raw setting, and provides parsed values to use and process the attachments

### ParsedEmail.Headers
  this field is deprecated. Use the SendGrid processed fields in ParsedValues instead. While it maintains its presence to avoid breaking changes, it provides raw, unprocessed headers and not all email clients are compatible. For example. these fields will be empty if the email cient is Outlook.com


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


	for section, body := range strings.Split(parsedEmail.TextBody, "\n") {
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
