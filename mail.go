package sendgrid

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/mail"
	"time"

	"github.com/sendgrid/smtpapi-go"
)

// SGMail is representation of a valid SendGrid Mail
type SGMail struct {
	To       []string
	ToName   []string
	Cc       []string
	Subject  string
	Text     string
	HTML     string
	From     string
	Bcc      []string
	FromName string
	ReplyTo  string
	Date     string
	Files    map[string]string
	Content  map[string]string
	Headers  map[string]string
	smtpapi.SMTPAPIHeader
}

// NewMail returns a new *SGMail
func NewMail() *SGMail {
	return &SGMail{}
}

// AddTo adds a valid email address
func (m *SGMail) AddTo(emails ...string) error {
	for i := 0; i < len(emails); i++ {
		address, err := mail.ParseAddress(emails[i])
		if err != nil {
			return err
		}
		m.AddRecipient(address)
	}
	return nil
}

// AddTos adds multiple email addresses
// Deprecated
func (m *SGMail) AddTos(emails []string) error {
	return m.AddTo(emails...)
}

// AddRecipient will add mail.Address emails to recipients.
func (m *SGMail) AddRecipient(recipients ...*mail.Address) {
	for i := 0; i < len(recipients); i++ {
		m.To = append(m.To, recipients[i].Address)
		if recipients[i].Name != "" {
			m.ToName = append(m.ToName, recipients[i].Name)
		}
	}
}

// AddRecipients calls AddRecipient per email
// Deprecated
func (m *SGMail) AddRecipients(recipients []*mail.Address) {
	m.AddRecipient(recipients...)
}

// AddToName sets the "pretty" name for a recipient
func (m *SGMail) AddToName(name string) {
	m.ToName = append(m.ToName, name)
}

// AddToNames sets the "pretty" name for multiple recipients
func (m *SGMail) AddToNames(names []string) {
	m.ToName = append(m.ToName, names...)
}

// AddCc ...
func (m *SGMail) AddCc(ccs ...string) error {
	for i := 0; i < len(ccs); i++ {
		address, err := mail.ParseAddress(ccs[i])
		if err != nil {
			return err
		}
		m.AddCcRecipient(address)
	}
	return nil
}

// AddCcs ...
// Deprecated
func (m *SGMail) AddCcs(ccs []string) error {
	return m.AddCc(ccs...)
}

// AddCcRecipient ...
func (m *SGMail) AddCcRecipient(recipients ...*mail.Address) {
	for i := 0; i < len(recipients); i++ {
		m.Cc = append(m.Cc, recipients[i].Address)
	}
}

// AddCcRecipients ...
// Deprecated
func (m *SGMail) AddCcRecipients(recipients []*mail.Address) {
	m.AddCcRecipient(recipients...)
}

// SetSubject sets the email's subject
func (m *SGMail) SetSubject(subject string) {
	m.Subject = subject
}

// SetText sets the email's text
func (m *SGMail) SetText(text string) {
	m.Text = text
}

// SetHTML sets the email's HTML
func (m *SGMail) SetHTML(html string) {
	m.HTML = html
}

// SetFrom will set the senders email property
func (m *SGMail) SetFrom(from string) error {
	address, err := mail.ParseAddress(from)
	if err != nil {
		return err
	}
	m.SetFromEmail(address)
	return nil
}

// SetFromEmail sets the senders email property
func (m *SGMail) SetFromEmail(address *mail.Address) {
	m.From = address.Address
	if address.Name != "" {
		m.SetFromName(address.Name)
	}
}

// AddBcc ...
func (m *SGMail) AddBcc(bccs ...string) error {
	for i := 0; i < len(bccs); i++ {
		address, err := mail.ParseAddress(bccs[i])
		if err != nil {
			return err
		}
		m.AddBccRecipient(address)
	}
	return nil
}

// AddBccs ...
// Deprecated
func (m *SGMail) AddBccs(bccs []string) error {
	return m.AddBcc(bccs...)
}

// AddBccRecipient ...
func (m *SGMail) AddBccRecipient(recipients ...*mail.Address) {
	for i := 0; i < len(recipients); i++ {
		m.Bcc = append(m.Bcc, recipients[i].Address)
	}

}

// AddBccRecipients ...
// Deprecated
func (m *SGMail) AddBccRecipients(recipients []*mail.Address) {
	m.AddBccRecipient(recipients...)
}

// SetFromName ...
func (m *SGMail) SetFromName(fromname string) {
	m.FromName = fromname
}

// SetReplyTo ...
func (m *SGMail) SetReplyTo(replyto string) error {
	address, err := mail.ParseAddress(replyto)
	if err != nil {
		return err
	}
	m.SetReplyToEmail(address)
	return nil
}

// SetReplyToEmail ...
func (m *SGMail) SetReplyToEmail(address *mail.Address) {
	m.ReplyTo = address.Address
}

// SetDate ...
func (m *SGMail) SetDate(date string) {
	m.Date = date
}

// SetRFCDate ...
func (m *SGMail) SetRFCDate(date time.Time) {
	m.Date = date.Format(time.RFC822)
}

// AddAttachment allows file attachments to be sent. For security reasons,
// this method doesn't take filepaths only the io.Reader interface.
func (m *SGMail) AddAttachment(filename string, file io.Reader) error {
	stream, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	m.AddAttachmentFromStream(filename, string(stream))
	return nil
}

// AddAttachmentFromStream ...
func (m *SGMail) AddAttachmentFromStream(filename, file string) {
	if m.Files == nil {
		m.Files = make(map[string]string)
	}
	m.Files[filename] = file
}

// AddContentID ...
func (m *SGMail) AddContentID(id, value string) {
	if m.Content == nil {
		m.Content = make(map[string]string)
	}
	m.Content[id] = value
}

// AddHeader ...
func (m *SGMail) AddHeader(header, value string) {
	if m.Headers == nil {
		m.Headers = make(map[string]string)
	}
	m.Headers[header] = value
}

func (m *SGMail) HeadersString() (string, error) {
	headers, e := json.Marshal(m.Headers)
	return string(headers), e
}
