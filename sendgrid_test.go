package sendgrid

import (
	"net/mail"
	"os"
	"testing"
)

func Test_Send(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddTo("Yamil Asusta <yamil.asusta@sendgrid.com>")
	message.AddBCC("yamil.asusta@sendgrid.com")
	address, _ := mail.ParseAddress("Yamil Asusta <yamil.asusta@upr.edu>")
	message.AddReceipient(address)
	message.AddSubject("SendGrid Testing")
	message.AddHTML("WIN")
	message.AddFrom("yamil@sendgrid.com")
	message.AddCategory("Test")
	message.AddCategory("Testing")
	filepath, _ := os.Getwd()
	message.AddAttachment(filepath + "/sendgrid.go")
	if r := sg.Send(message); r == nil {
		t.Log("Test_Send Passed")
	} else {
		t.Error("Test_Send Failed", r)
	}
}

func Test_SendSMTP(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddTo("Yamil Asusta <yamil.asusta@sendgrid.com>")
	address, _ := mail.ParseAddress("Yamil Asusta <yamil.asusta@upr.edu>")
	message.AddReceipient(address)
	message.AddSubject("SendGrid Testing")
	message.AddHTML("<html>SMTP :(</html>")
	message.AddText("Work!")
	message.AddFrom("yamil@sendgrid.com")
	message.AddCategory("Testing")
	filepath, _ := os.Getwd()
	message.AddAttachment(filepath + "/sendgrid.go")
	if r := sg.SendSMTP(message); r == nil {
		t.Log("Test_SendAPI Passed")
	} else {
		t.Error("Test_SendAPI Failed", r)
	}
}

func Test_SendAPI(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddTo("Yamil Asusta <yamil.asusta@sendgrid.com>")
	address, _ := mail.ParseAddress("Yamil Asusta <yamil.asusta@upr.edu>")
	message.AddReceipient(address)
	message.AddSubject("SendGrid Testing")
	message.AddHTML("WIN")
	message.AddFrom("yamil@sendgrid.com")
	filepath, _ := os.Getwd()
	message.AddAttachment(filepath + "/sendgrid.go")
	if r := sg.SendAPI(message); r == nil {
		t.Log("Test_SendAPI Passed")
	} else {
		t.Error("Test_SendAPI Failed", r)
	}
}
