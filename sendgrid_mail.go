package sendgrid

import (
	"encoding/json"
	"github.com/elbuo8/smtpmail"
	"net/mail"
)

type SGMail struct {
	mail    smtpmail.Mail
	headers SGHeaders
}

// SGHeaders will be used to set up X-SMTPAPI params
type SGHeaders struct {
	Sub         map[string][]string                     `json:"sub,omitempty"`
	Section     map[string]string                       `json:"section,omitempty"`
	Category    []string                                `json:"category,omitempty"`
	Unique_args map[string]string                       `json:"unique_args,omitempty"`
	Filters     map[string]map[string]map[string]string `json:"filters,omitempty"`
}

func NewMail() SGMail {
	return SGMail{
		mail: smtpmail.NewMail(),
	}
}

func (m *SGMail) AddTo(email string) error {
	return m.mail.AddTo(email)
}

func (m *SGMail) AddToName(name string) {
	m.mail.AddToName(name)
}

func (m *SGMail) AddReceipient(receipient *mail.Address) {
	m.mail.AddReceipient(receipient)
}

func (m *SGMail) AddSubject(subject string) {
	m.mail.AddSubject(subject)
}

func (m *SGMail) AddHTML(html string) {
	m.mail.AddHTML(html)
}

func (m *SGMail) AddText(text string) {
	m.mail.AddText(text)
}

func (m *SGMail) AddFrom(from string) {
	m.mail.AddFrom(from)
}

func (m *SGMail) AddBCC(bcc string) error {
	return m.mail.AddBCC(bcc)
}

func (m *SGMail) AddReceipientBCC(bcc *mail.Address) {
	m.mail.AddReceipientBCC(bcc)
}

func (m *SGMail) AddFromName(name string) {
	m.mail.AddFromName(name)
}

func (m *SGMail) AddReplyTo(reply string) {
	m.mail.AddReplyTo(reply)
}

func (m *SGMail) AddDate(date string) {
	m.mail.AddDate(date)
}

func (m *SGMail) AddAttachment(filePath string) {
	m.mail.AddAttachment(filePath)
}

func (m *SGMail) AddSubstitution(key, sub string) {
	if m.headers.Sub == nil {
		m.headers.Sub = make(map[string][]string)
	}
	subKey := "%" + key + "%"
	m.headers.Sub[subKey] = append(m.headers.Sub[subKey], sub)
}

func (m *SGMail) AddSection(section, value string) {
	if m.headers.Section == nil {
		m.headers.Section = make(map[string]string)
	}
	secKey := "%" + section + "%"
	m.headers.Section[secKey] = value
}

func (m *SGMail) AddCategory(value string) {
	m.headers.Category = append(m.headers.Category, value)
}

func (m *SGMail) AddUniqueArg(arg, value string) {
	if m.headers.Unique_args == nil {
		m.headers.Unique_args = make(map[string]string)
	}
	m.headers.Unique_args[arg] = value
}

func (m *SGMail) AddFilter(filter, setting, value string) {
	if m.headers.Filters == nil {
		m.headers.Filters = make(map[string]map[string]map[string]string)
	}
	if m.headers.Filters[filter] == nil {
		m.headers.Filters[filter] = make(map[string]map[string]string)
	}
	if m.headers.Filters[filter]["settings"] == nil {
		m.headers.Filters[filter]["settings"] = make(map[string]string)
	}
	m.headers.Filters[filter]["settings"][setting] = value
}

func (m *SGMail) SetHeaders() error {
	headers, e := json.Marshal(m.headers)
	if e != nil {
		return e
	}
	m.mail.AddHeaders(string(headers))
	return nil
}
