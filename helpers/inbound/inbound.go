package inbound

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

// ParsedEmail defines a multipart parsed email
// Body and Attachments are only populated if the Raw option is checked on the SendGrid inbound configuration and are named for backwards compatability
type ParsedEmail struct {
	// Header values are raw and not pre-processed by SendGrid. They may change depending on the email client. Use carefully
	Headers map[string]string
	// Please see https://docs.sendgrid.com/for-developers/parsing-email/setting-up-the-inbound-parse-webhook to see the available fields in the email headers
	// all fields listed there are available within the headers map except for text which lives in the TextBody field
	ParsedValues map[string]string
	// Primary email body parsed with \n. A common approach is to Split by the \n to bring every line of the email into a string array
	TextBody string

	// Envelope expresses the exact email address that the email was addressed to and the exact email address it was from, without extra characters
	Envelope struct {
		From string   `json:"from"`
		To   []string `json:"to"`
	}

	// Attachments have been fully parsed to include the filename, size, content type and actual file for uploading or processing
	ParsedAttachments map[string]*EmailAttachment

	// Raw only
	Attachments map[string][]byte
	// accessed with text/html and text/plain. text/plain is always parsed to the TextBody field
	Body map[string]string

	rawRequest      *http.Request
	rawValues       map[string][]string
	withAttachments bool
}

// EmailAttachment defines information related to an email attachment
type EmailAttachment struct {
	File        multipart.File `json:"-"`
	Filename    string         `json:"filename"`
	Size        int64          `json:"-"`
	ContentType string         `json:"type"`
}

func newParsedEmail(request *http.Request) ParsedEmail {
	return ParsedEmail{
		Headers:           make(map[string]string),
		ParsedValues:      make(map[string]string),
		ParsedAttachments: make(map[string]*EmailAttachment),

		Body:        make(map[string]string),
		Attachments: make(map[string][]byte),

		rawRequest:      request,
		withAttachments: false,
	}
}

// Parse parses an email using Go's multipart parser and populates the headers, and body
// This method skips processing the attachment file and is therefore more performant
func Parse(request *http.Request) (*ParsedEmail, error) {
	result := newParsedEmail(request)

	err := result.parse()
	return &result, err
}

// ParseWithAttachments parses an email using Go's multipart parser and populates the headers, body and processes attachments
func ParseWithAttachments(request *http.Request) (*ParsedEmail, error) {
	result := newParsedEmail(request)
	result.withAttachments = true

	err := result.parse()
	return &result, err
}

func (email *ParsedEmail) parse() error {
	err := email.rawRequest.ParseMultipartForm(0)
	if err != nil {
		return err
	}

	email.rawValues = email.rawRequest.MultipartForm.Value

	// unmarshal the envelope
	if len(email.rawValues["envelope"]) > 0 {
		if err := json.Unmarshal([]byte(email.rawValues["envelope"][0]), &email.Envelope); err != nil {
			return err
		}
	}

	// parse included headers
	if len(email.rawValues["headers"]) > 0 {
		email.parseHeaders(email.rawValues["headers"][0])
	}

	// apply the rest of the SendGrid fields to the headers map
	for k, v := range email.rawValues {
		if k == "text" || k == "email" || k == "headers" || k == "envelope" {
			continue
		}

		if len(v) > 0 {
			email.ParsedValues[k] = v[0]
		}
	}

	// apply the plain text body
	if len(email.rawValues["text"]) > 0 {
		email.TextBody = email.rawValues["text"][0]
	}

	// only included if the raw box is checked
	if len(email.rawValues["email"]) > 0 {
		email.parseRawEmail(email.rawValues["email"][0])
	}

	// if the client chose not to parse attachments, return as is
	if !email.withAttachments {
		return nil
	}

	return email.parseAttachments(email.rawValues)
}

func (email *ParsedEmail) parseAttachments(values map[string][]string) error {
	if len(values["attachment-info"]) != 1 {
		return nil
	}
	// unmarshal the sendgrid parsed aspects of the email attachment into the attachment struct
	if err := json.Unmarshal([]byte(values["attachment-info"][0]), &email.ParsedAttachments); err != nil {
		return err
	}

	// range through the multipart files
	for key, val := range email.rawRequest.MultipartForm.File {
		// open the attachment file for processing
		file, err := val[0].Open()
		if err != nil {
			return err
		}

		// add the actual file and the size to the parsed files
		email.ParsedAttachments[key].File = file
		email.ParsedAttachments[key].Size = val[0].Size

		// if the file does not have a name. give it Untitled
		if email.ParsedAttachments[key].Filename == "" {
			email.ParsedAttachments[key].Filename = "Untitled"
		}
	}

	return nil
}

func (email *ParsedEmail) parseRawEmail(rawEmail string) error {
	sections := strings.SplitN(rawEmail, "\n\n", 2)
	email.parseHeaders(sections[0])
	raw, err := parseMultipart(strings.NewReader(sections[1]), email.Headers["Content-Type"])
	if err != nil {
		return err
	}

	for {
		emailPart, err := raw.NextPart()
		if err == io.EOF {
			return nil
		}
		rawEmailBody, err := parseMultipart(emailPart, emailPart.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
		if rawEmailBody != nil {
			for {
				emailBodyPart, err := rawEmailBody.NextPart()
				if err == io.EOF {
					break
				}
				header := emailBodyPart.Header.Get("Content-Type")
				b, err := io.ReadAll(emailPart)
				if err != nil {
					return err
				}
				email.Body[header] = string(b)
			}

		} else if emailPart.FileName() != "" {
			b, err := io.ReadAll(emailPart)
			if err != nil {
				return err
			}
			email.Attachments[emailPart.FileName()] = b
		} else {
			header := emailPart.Header.Get("Content-Type")
			b, err := io.ReadAll(emailPart)
			if err != nil {
				return err
			}

			email.Body[header] = string(b)
		}
	}
}

func parseMultipart(body io.Reader, contentType string) (*multipart.Reader, error) {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		return multipart.NewReader(body, params["boundary"]), nil
	}
	return nil, nil
}

func (email *ParsedEmail) parseHeaders(headers string) {
	splitHeaders := strings.Split(strings.TrimSpace(headers), "\n")
	for _, header := range splitHeaders {
		splitHeader := strings.SplitN(header, ": ", 2)
		// keeps outlook emails from causing a panic
		if len(splitHeader) != 2 {
			continue
		}

		email.Headers[splitHeader[0]] = splitHeader[1]
	}
}

// Validate validates the DKIM and SPF scores to ensure that the email client and address was not spoofed
func (email *ParsedEmail) Validate() error {
	if len(email.rawValues["dkim"]) == 0 || len(email.rawValues["SPF"]) == 0 {
		return fmt.Errorf("missing DKIM and SPF score")
	}

	for _, val := range email.rawValues["dkim"] {
		if !strings.Contains(val, "pass") {
			return fmt.Errorf("DKIM validation failed")
		}
	}

	for _, val := range email.rawValues["SPF"] {
		if !strings.Contains(val, "pass") {
			return fmt.Errorf("SPF validation failed")
		}
	}

	return nil
}
