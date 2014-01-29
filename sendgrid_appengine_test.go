// +build appengine

package sendgrid

import (
	"net/mail"
	"testing"

	aemail "appengine/mail"
)

func TestMigrate(t *testing.T) {
	message := NewMail()
	address, _ := mail.ParseAddress("John Doe <john@email.com>")
	message.AddRecipient(address)
	message.AddSubject("test")
	message.AddHTML("html")
	message.AddText("text")
	message.AddFrom("doe@email.com")
	message.AddFromName("Doe Email")
	message.AddBCC("bcc@host.com")

	aemessage := aemail.Message{
		To:       []string{"John Doe <john@email.com>"},
		Sender:   "Doe Email <doe@email.com>",
		Subject:  "test",
		HTMLBody: "html",
		Body:     "text",
		Bcc:      []string{"bcc@host.com"},
	}

	message2, err := migrateMail(&aemessage)
	if err != nil {
		t.Errorf("Error migrating mail - %v", err)
	}

	sg := NewSendGridClient("", "")
	val1, err := sg.buildUrl(message)
	if err != nil {
		t.Errorf("%s", err)
	}

	val2, err := sg.buildUrl(*message2)
	if err != nil {
		t.Errorf("%s", err)
	}

	for key := range val1 { // make sure everything in val1 is in val2
		if val1.Get(key) != val2.Get(key) {
			t.Errorf("Expected %#v, got %#v", val1.Get(key), val2.Get(key))
		}
		val2.Del(key)
	}
	for key := range val2 { // make sure nothing in val2 isn't in val1
		if val1.Get(key) != val2.Get(key) {
			t.Errorf("Expected %#v, got %#v", val1.Get(key), val2.Get(key))
		}
	}
}
