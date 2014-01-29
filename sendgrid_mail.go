package sendgrid

import (
	"io/ioutil"
	"net/http"
	"net/mail"
	"path/filepath"
)

type SGMail struct {
	smtpmail.Mail
	smtpapi.SMTPAPIHeader
}

func NewMail() SGMail {
	return SGMail{}
}

func (m *SGMail) AddAttachment(filePath string) error {
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	m.AddAttachmentStream(filename, bytes)
	return nil
}

func (m *SGMail) AddAttachmentLink(filename, link string) error {
	if r, e := http.Get(link); e != nil {
		return e
	} else {
		defer r.Body.Close()
		bytes, _ := ioutil.ReadAll(r.Body)
		m.AddAttachmentStream(filename, bytes)
		return nil
	}
}

func (m *SGMail) AddAttachmentStream(filename string, stream []byte) {
	if m.Files == nil {
		m.Files = make(map[string]string)
	}
	m.Files[filename] = string(stream)
}

func (m *SGMail) AddTo(email string) {
	m.Mail.AddTo(email)
}

func (m *SGMail) AddRecipient(email *mail.Address) {
	m.Mail.AddRecipient(email)
}
