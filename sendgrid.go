// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"path/filepath"
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

// SendSMTP - SMTP interface. Still being developed. Use API instead.
func (sg *SGClient) SendSMTP(m Mail) error {
	return smtp.SendMail(sg.smtpUrl+":"+sg.smtpPort, sg.smtpAuth, m.from, m.to, []byte(m.html))
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
	headers, e := json.Marshal(m.headers)
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

// Mail will represent a formatted email
type Mail struct {
	to       []string
	toname   []string
	subject  string
	html     string
	text     string
	from     string
	bcc      []string
	fromname string
	replyto  string
	date     string
	files    map[string]string
	headers  map[string]string
}

// NewMail returns a new Mail
func NewMail() Mail {
	return Mail{}
}

// AddTo will take a valid email address and store it in the mail.
// It will return an error if the email is invalid.
func (m *Mail) AddTo(email string) error {
	if parsedAddess, e := mail.ParseAddress(email); e != nil {
		return e
	} else {
		m.to = append(m.to, parsedAddess.Address)
		if parsedAddess.Name != "" {
			m.toname = append(m.toname, parsedAddess.Name)
		}
		return nil
	}
}

// AddToName will add a new receipient name to mail
func (m *Mail) AddToName(name string) {
	m.toname = append(m.toname, name)
}

// AddReceipient will take an already parsed mail.Address
func (m *Mail) AddReceipient(receipient *mail.Address) {
	m.to = append(m.to, receipient.Address)
	if receipient.Name != "" {
		m.toname = append(m.toname, receipient.Name)
	}
}

// AddSubject will set the subject of the mail
func (m *Mail) AddSubject(s string) {
	m.subject = s
}

// AddHTML will set the body of the mail
func (m *Mail) AddHTML(html string) {
	m.html = html
}

// AddText will set the body of the email
func (m *Mail) AddText(text string) {
	m.text = text
}

// AddFrom will set the senders email
func (m *Mail) AddFrom(from string) {
	m.from = from
}

// AddBCC works like AddTo but for BCC
func (m *Mail) AddBCC(email string) error {
	if parsedAddess, e := mail.ParseAddress(email); e != nil {
		return e
	} else {
		m.bcc = append(m.bcc, parsedAddess.Address)
		return nil
	}
}

// AddReceipientBCC works like AddReceipient but for BCC
func (m *Mail) AddReceipientBCC(email mail.Address) {
	m.bcc = append(m.bcc, email.Address)
}

// AddFromName will set the senders name
func (m *Mail) AddFromName(name string) {
	m.fromname = name
}

// AddReplyTo will set the return address
func (m *Mail) AddReplyTo(reply string) {
	m.replyto = reply
}

// AddDate specifies the date
func (m *Mail) AddDate(date string) {
	m.date = date
}

// AddHeader allows custom headers to be added to mail
func (m *Mail) AddHeader(header, value string) {
	if m.headers == nil {
		m.headers = make(map[string]string)
	}
	m.headers[header] = value
}

// AddAttachment will include file/s in mail
func (m *Mail) AddAttachment(filePath string) error {
	if m.files == nil {
		m.files = make(map[string]string)
	}
	buf, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	m.files[filename] = base64.StdEncoding.EncodeToString(buf)
	return nil
}
