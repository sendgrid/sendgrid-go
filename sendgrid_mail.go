package sendgrid

import (
	"github.com/sendgrid/smtpapi-go"
	//"io/ioutil"
	//"net/http"
	"net/mail"
	//"path/filepath"
)

type SGMail struct {
	To       []string
	ToName   []string
	Subject  string
	Text     string
	Html     string
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

func NewMail() *SGMail {
	return &SGMail{}
}

func (m *SGMail) AddTo(email string) error {
	m.SMTPAPIHeader.AddTo(email)
	if address, err := mail.ParseAddress(email); err != nil {
		return err
	} else {
		m.AddRecipient(address)
		return nil
	}
}

func (m *SGMail) AddTos(emails []string) error {
	for i := 0; i < len(emails); i++ {
		m.SMTPAPIHeader.AddTo(emails[i])
	}
	if addresses, err := mail.ParseAddressList(emails); err != nil {
		return err
		m.AddRecipients(addresses)
	}
}

func (m *SGMail) AddRecipient(recipient *mail.Address) {
	m.To = append(m.To, recipient.Address)
	if recipient.Name != "" {
		m.ToName = append(m.ToName, recipient)
	}
}

func (m *SGMail) AddRecipients(recipients *mail.Address) {
	for i := 0; i < len(recipients); i++ {
		m.AddRecipient(recipients[i])
	}
}
