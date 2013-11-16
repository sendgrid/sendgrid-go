package sendgrid

import (
	"os"
	"testing"
)

func Test_Send(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("SendGrid is Baller")
	message.AddHTML("Simple Text")
	message.AddFrom("yamil@sendgrid.com")
	message.AddHeader("X-Mailer", "Test")
	if r := sg.Send(message); r == nil {
		t.Log("Test_Send Passed")
	} else {
		t.Error("Test_Send Failed")
	}
}

func Test_SendAPI(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("SendGrid is Baller")
	message.AddHTML("Simple Text")
	message.AddFrom("yamil@sendgrid.com")
	message.AddHeader("X-Mailer", "Test")
	if r := sg.SendAPI(message); r == nil {
		t.Log("Test_Send Passed")
	} else {
		t.Error("Test_Send Failed")
	}
}
