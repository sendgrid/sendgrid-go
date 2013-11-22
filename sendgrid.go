// Package sendgrid provides a simple interface to interact with the SendGrid API
// Special thanks to this gist -> https://gist.github.com/rmulley/6603544
package sendgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
)

// SGClient will contain the credentials and default values
type SGClient struct {
	apiUser  string
	apiPwd   string
	apiUrl   string
	smtpUrl  string
	smtpPort string
	smtpAuth smtp.Auth
	Client   *http.Client
}

// NewSendGridClient will return a new SGClient.
func NewSendGridClient(apiUser, apiPwd string) SGClient {
	smtpUrl := "smtp.sendgrid.net"
	smtpPort := "587"
	apiUrl := "https://api.sendgrid.com/api/mail.send.json?"
	smtpAuth := smtp.PlainAuth("", apiUser, apiPwd, smtpUrl)
	return SGClient{
		apiUser:  apiUser,
		apiPwd:   apiPwd,
		apiUrl:   apiUrl,
		smtpUrl:  smtpUrl,
		smtpPort: smtpPort,
		smtpAuth: smtpAuth,
	}
}

// Send will try to use the WebAPI first. If it's a success then a nill will be returned.
// Else, the SMTP API will be used as a fail over. If this happens regardless of what is the result
// of the SMTP attempt, you will receive an array of errors.
// If the length of the array is 1, then SMTP succeeded, else it also failed.
func (sg *SGClient) Send(m Mail) []error {
	if apiError := sg.SendAPI(m); apiError != nil {
		var errors []error
		errors = append(errors, apiError)
		if smtpError := sg.SendSMTP(m); smtpError != nil {
			return append(errors, smtpError)
		} else {
			return errors
		}
	} else {
		return nil
	}
}

// SendSMTP - It can be used for generic SMTP stuff
func (sg *SGClient) SendSMTP(m Mail) error {
	boundary := "SENDGRIDGOLIB"
	var message bytes.Buffer
	headers, e := json.Marshal(m.sgheaders)
	if e != nil {
		return fmt.Errorf("sendgrid.go: Error parsing JSON headers")
	}
	message.WriteString(fmt.Sprintf("X-SMTPAPI: %s\r\n", headers))
	message.WriteString(fmt.Sprintf("From: %s <%s>\r\n", m.fromname, m.from))
	message.WriteString(fmt.Sprintf("To: <%s>", m.to[0]))
	for i := 1; i < len(m.to); i++ {
		message.WriteString(fmt.Sprintf(", <%s>", m.to[i]))
	}
	if len(m.bcc) > 0 {
		message.WriteString(fmt.Sprintf("Bcc: <%s>", m.bcc[0]))
		for i := 1; i < len(m.to); i++ {
			message.WriteString(fmt.Sprintf(", <%s>", m.bcc[i]))
		}
	}
	message.WriteString("\r\n")
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", m.subject))
	message.WriteString("MIME-Version: 1.0\r\n")
	if m.files != nil {
		message.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\n--%s\r\n", boundary, boundary))
	}
	if m.html != "" {
		part := fmt.Sprintf("Content-Type: text/html\r\n\n%s\r\n\n", m.html)
		message.WriteString(part)
	} else {
		part := fmt.Sprintf("Content-Type: text/plain\r\n\n%s\r\n\n", m.text)
		message.WriteString(part)
	}
	if m.files != nil {
		for key, value := range m.files {
			message.WriteString(fmt.Sprintf("--%s\r\n", boundary))
			message.WriteString("Content-Type: application/octect-stream\r\n")
			message.WriteString("Content-Transfer-Encoding:base64\r\n")
			message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n\r\n%s\r\n\n", key, value))
		}
		message.WriteString(fmt.Sprintf("--%s--", boundary))
	}
	return smtp.SendMail(sg.smtpUrl+":"+sg.smtpPort, sg.smtpAuth, m.from, m.to, message.Bytes())
}

// SendAPI will send mail using SG web API
func (sg *SGClient) SendAPI(m Mail) error {
	values := url.Values{}
	values.Set("api_user", sg.apiUser)
	values.Set("api_key", sg.apiPwd)
	values.Set("subject", m.subject)
	values.Set("html", m.html)
	values.Set("text", m.text)
	values.Set("from", m.from)
	headers, e := json.Marshal(m.sgheaders)
	if e != nil {
		return fmt.Errorf("sendgrid.go: Error parsing JSON headers")
	}
	values.Set("headers", string(headers[:]))
	for i := 0; i < len(m.to); i++ {
		values.Set("to[]", m.to[i])
	}
	for i := 0; i < len(m.bcc); i++ {
		values.Set("bcc[]", m.bcc[i])
	}
	for i := 0; i < len(m.toname); i++ {
		values.Set("toname[]", m.toname[i])
	}
	for k, v := range m.files {
		values.Set("files["+k+"]", v)
	}
	if sg.Client == nil {
		sg.Client = http.DefaultClient
	}
	r, e := sg.Client.PostForm(sg.apiUrl, values)
	defer r.Body.Close()
	if r.StatusCode == 200 && e == nil {
		return nil
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("sendgrid.go: code:%d error:%v body:%s", r.StatusCode, e, body)
	}
}
