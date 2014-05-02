package smtpmail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"path/filepath"
)

// Mail will represent a formatted email
type Mail struct {
	To       []string
	ToName   []string
	Subject  string
	HTML     string
	Text     string
	From     string
	Bcc      []string
	FromName string
	ReplyTo  string
	Date     string
	Files    map[string]string
	Headers  string
}

// NewMail returns a new Mail
func NewMail() Mail {
	return Mail{}
}

type SMTPClient struct {
	smtpAuth smtp.Auth
	url      string
	port     string
}

func NewSMTPClient(user, pwd, url, port string) SMTPClient {
	return SMTPClient{
		smtpAuth: smtp.PlainAuth("", user, pwd, url),
		url:      url,
		port:     port,
	}
}

// SendSMTP - It can be used for generic SMTP stuff
func (c *SMTPClient) Send(m Mail) error {
	boundary := "ELBUO8BOUNDARYFORSMTPGOLIB"
	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("X-SMTPAPI: %s\r\n", m.Headers))
	message.WriteString(fmt.Sprintf("From: %s <%s>\r\n", m.FromName, m.From))
	if len(m.To) > 0 {
		message.WriteString(fmt.Sprintf("To: <%s>", m.To[0]))
		for i := 1; i < len(m.To); i++ {
			message.WriteString(fmt.Sprintf(", <%s>", m.To[i]))
		}
	}
	if len(m.Bcc) > 0 {
		message.WriteString(fmt.Sprintf("Bcc: <%s>", m.Bcc[0]))
		for i := 1; i < len(m.Bcc); i++ {
			message.WriteString(fmt.Sprintf(", <%s>", m.Bcc[i]))
		}
	}
	message.WriteString("\r\n")
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))
	message.WriteString("MIME-Version: 1.0\r\n")
	if m.Files != nil {
		message.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\n--%s\r\n", boundary, boundary))
	}
	if m.HTML != "" {
		part := fmt.Sprintf("Content-Type: text/html\r\n\n%s\r\n\n", m.HTML)
		message.WriteString(part)
	} else {
		part := fmt.Sprintf("Content-Type: text/plain\r\n\n%s\r\n\n", m.Text)
		message.WriteString(part)
	}
	if m.Files != nil {
		for key, value := range m.Files {
			message.WriteString(fmt.Sprintf("--%s\r\n", boundary))
			message.WriteString("Content-Type: application/octect-stream\r\n")
			message.WriteString("Content-Transfer-Encoding:base64\r\n")
			message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n\r\n%s\r\n\n", key, value))
		}
		message.WriteString(fmt.Sprintf("--%s--", boundary))
	}
	return smtp.SendMail(c.url+":"+c.port, c.smtpAuth, m.From, m.To, message.Bytes())
}

// AddTo will take a valid email address and store it in the mail.
// It will return an error if the email is invalid.
func (m *Mail) AddTo(email string) error {
	if parsedAddess, e := mail.ParseAddress(email); e != nil {
		return e
	} else {
		m.AddRecipient(parsedAddess)
		return nil
	}
}

func (m *Mail) SetTos(emails []string) {
	m.To = emails
}

// AddToName will add a new receipient name to mail
func (m *Mail) AddToName(name string) {
	m.ToName = append(m.ToName, name)
}

// AddRecipient will take an already parsed mail.Address
func (m *Mail) AddRecipient(receipient *mail.Address) {
	m.To = append(m.To, receipient.Address)
	if receipient.Name != "" {
		m.ToName = append(m.ToName, receipient.Name)
	}
}

// AddSubject will set the subject of the mail
func (m *Mail) AddSubject(s string) {
	m.Subject = s
}

// AddHTML will set the body of the mail
func (m *Mail) AddHTML(html string) {
	m.HTML = html
}

// AddText will set the body of the email
func (m *Mail) AddText(text string) {
	m.Text = text
}

// AddFrom will set the senders email
func (m *Mail) AddFrom(from string) {
	m.From = from
}

// AddBCC works like AddTo but for BCC
func (m *Mail) AddBCC(email string) error {
	if parsedAddess, e := mail.ParseAddress(email); e != nil {
		return e
	} else {
		m.Bcc = append(m.Bcc, parsedAddess.Address)
		return nil
	}
}

// AddRecipientBCC works like AddRecipient but for BCC
func (m *Mail) AddRecipientBCC(email *mail.Address) {
	m.Bcc = append(m.Bcc, email.Address)
}

// AddFromName will set the senders name
func (m *Mail) AddFromName(name string) {
	m.FromName = name
}

// AddReplyTo will set the return address
func (m *Mail) AddReplyTo(reply string) {
	m.ReplyTo = reply
}

// AddDate specifies the date
func (m *Mail) AddDate(date string) {
	m.Date = date
}

// AddAttachment will include file/s in mail
func (m *Mail) AddAttachment(filePath string) error {
	if m.Files == nil {
		m.Files = make(map[string]string)
	}
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	encoded := base64.StdEncoding.EncodeToString(file)
	totalChars := len(encoded)
	maxLength := 500
	totalLines := totalChars / maxLength
	var buf bytes.Buffer
	for i := 0; i < totalLines; i++ {
		buf.WriteString(encoded[i*maxLength:(i+1)*maxLength] + "\n")
	}
	buf.WriteString(encoded[totalLines*maxLength:])
	m.Files[filename] = buf.String()
	return nil
}

func (m *Mail) AddHeaders(headers string) {
	m.Headers = headers
}
