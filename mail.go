package sendgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"net/url"
	"time"

	"github.com/sendgrid/smtpapi-go"
)

// Mail is representation of a valid SendGrid Mail
type Mail struct {
	from    *mail.Address
	replyTo *mail.Address
	to      []*mail.Address
	cc      []*mail.Address
	bcc     []*mail.Address
	subject string
	text    string
	html    string
	date    string
	files   map[string][]byte
	content map[string]string
	headers map[string]string

	SMTPHeaders smtpapi.Header
}

// NewMail returns a new *SGMail
func NewMail() *Mail {
	return &Mail{}
}

// To adds a valid email address
func (m *Mail) To(emails ...*mail.Address) {
	m.to = append(m.to, emails...)
}

// Cc ...
func (m *Mail) Cc(emails ...*mail.Address) {
	m.cc = append(m.cc, emails...)
}

// SetSubject sets the email's subject
func (m *Mail) Subject(subject string) {
	m.subject = subject
}

// SetText sets the email's text
func (m *Mail) Text(text string) {
	m.text = text
}

// SetHTML sets the email's HTML
func (m *Mail) HTML(html string) {
	m.html = html
}

// SetFrom will set the senders email property
func (m *Mail) From(from *mail.Address) {
	m.from = from
}

// Bcc ...
func (m *Mail) Bcc(emails ...*mail.Address) {
	m.bcc = append(m.bcc, emails...)
}

// SetReplyToEmail ...
func (m *Mail) ReplyTo(address *mail.Address) {
	m.replyTo = address
}

// SetDate ...
func (m *Mail) SetDate(date string) {
	m.date = date
}

// SetRFCDate ...
func (m *Mail) SetRFCDate(date time.Time) {
	m.date = date.Format(time.RFC822)
}

// Attachment allows file attachments to be sent. For security reasons,
// this method doesn't take filepaths only the io.Reader interface.
func (m *Mail) Attach(filename string, file io.Reader) error {
	content := new(bytes.Buffer)
	_, err := content.ReadFrom(file)
	if err != nil {
		return err
	}
	m.AttachBytes(filename, content.Bytes())
	return nil
}

// AddAttachmentFromStream ...
func (m *Mail) AttachBytes(filename string, file []byte) {
	if m.files == nil {
		m.files = make(map[string][]byte)
	}
	m.files[filename] = file
}

// AddContentID ...
func (m *Mail) ContentID(id, value string) {
	if m.content == nil {
		m.content = make(map[string]string)
	}
	m.content[id] = value
}

// AddHeader ...
func (m *Mail) AddHeader(header, value string) {
	if m.headers == nil {
		m.headers = make(map[string]string)
	}
	m.headers[header] = value
}

func (m *Mail) values() (url.Values, error) {
	values := url.Values{}

	apiHeaders, err := m.SMTPHeaders.JSONString()
	if err != nil {
		return nil, fmt.Errorf("sendgrid.go: error:%v", err)
	}
	values.Set("x-smtpapi", apiHeaders)

	headers, err := json.Marshal(m.headers)
	if err != nil {
		return nil, fmt.Errorf("sendgrid.go: error: %v", err)
	}

	values.Set("headers", string(headers))
	values.Set("subject", m.subject)
	values.Set("html", m.html)
	values.Set("text", m.text)

	if m.from != nil {
		values.Set("from", m.from.Address)
		if m.from.Name != "" {
			values.Set("fromname", m.from.Name)
		}
	}

	for _, to := range m.to {
		values.Add("to[]", to.Address)
		values.Add("toname[]", to.Name)
	}

	if m.replyTo != nil {
		values.Set("replyto", m.replyTo.Address)
	}

	for _, cc := range m.cc {
		values.Add("cc[]", cc.Address)
	}
	for _, bcc := range m.bcc {
		values.Add("bcc[]", bcc.Address)
	}

	for k, v := range m.files {
		//XXX: Converting []byte to string is safe here?
		values.Set("files["+k+"]", string(v))
	}

	for k, v := range m.content {
		values.Set("content["+k+"]", v)
	}
	return values, nil
}
