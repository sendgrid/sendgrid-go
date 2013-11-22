package sendgrid

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/mail"
	"path/filepath"
)

// Mail will represent a formatted email
type Mail struct {
	to        []string
	toname    []string
	subject   string
	html      string
	text      string
	from      string
	bcc       []string
	fromname  string
	replyto   string
	date      string
	files     map[string]string
	sgheaders SGHeaders
}

// SGHeaders will be used to set up X-SMTPAPI params
type SGHeaders struct {
	Sub         map[string][]string                     `json:"sub,omitempty"`
	Section     map[string]string                       `json:"section,omitempty"`
	Category    []string                                `json:"category,omitempty"`
	Unique_args map[string]string                       `json:"unique_args,omitempty"`
	Filters     map[string]map[string]map[string]string `json:"filters,omitempty"`
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

// AddAttachment will include file/s in mail
func (m *Mail) AddAttachment(filePath string) error {
	if m.files == nil {
		m.files = make(map[string]string)
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
	m.files[filename] = buf.String()
	return nil
}

func (m *Mail) AddSubstitution(key, sub string) {
	if m.sgheaders.Sub == nil {
		m.sgheaders.Sub = make(map[string][]string)
	}
	subKey := "%" + key + "%"
	m.sgheaders.Sub[subKey] = append(m.sgheaders.Sub[subKey], sub)
}

func (m *Mail) AddSection(section, value string) {
	if m.sgheaders.Section == nil {
		m.sgheaders.Section = make(map[string]string)
	}
	secKey := "%" + section + "%"
	m.sgheaders.Section[secKey] = value
}

func (m *Mail) AddCategory(value string) {
	m.sgheaders.Category = append(m.sgheaders.Category, value)
}

func (m *Mail) AddUniqueArg(arg, value string) {
	if m.sgheaders.Unique_args == nil {
		m.sgheaders.Unique_args = make(map[string]string)
	}
	m.sgheaders.Unique_args[arg] = value
}

func (m *Mail) AddFilter(filter, setting, value string) {
	if m.sgheaders.Filters == nil {
		m.sgheaders.Filters = make(map[string]map[string]map[string]string)
	}
	if m.sgheaders.Filters[filter] == nil {
		m.sgheaders.Filters[filter] = make(map[string]map[string]string)
	}
	if m.sgheaders.Filters[filter]["settings"] == nil {
		m.sgheaders.Filters[filter]["settings"] = make(map[string]string)
	}
	m.sgheaders.Filters[filter]["settings"][setting] = value
}
