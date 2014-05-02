package smtpmail

import (
	"os"
	"testing"
)

func Test_Send(t *testing.T) {
	c := NewSMTPClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"), "smtp.sendgrid.net", "587")
	m := NewMail()
	m.AddTo("hello@yamilasusta.com")
	m.AddFrom("yamil@sendgrid.com")
	m.AddSubject("Testing")
	m.AddText("Some text :)")
	filepath, _ := os.Getwd()
	m.AddAttachment(filepath + "/mail.go")
	if e := c.Send(m); e != nil {
		t.Error(e)
	} else {
		t.Log("All good brah")
	}
}
