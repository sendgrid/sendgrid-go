package sendgrid

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"testing"
	"time"
)

func TestNewMail(t *testing.T) {
	m := NewMail()
	if m == nil {
		t.Errorf("NewMail() shouldn't return nil")
	}
}

func TestTo(t *testing.T) {
	addr, err := mail.ParseAddress("Email Name<inbox@domain.tld>")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.To(addr)
	if len(m.to) != 1 {
		t.Fatalf("To should append to Mail.to")
	}
}

func TestToMultiple(t *testing.T) {

	addrs, err := mail.ParseAddressList("Email Name <email+1@email.com>, email+2@email.com")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.To(addrs...)
	if len(m.to) != 2 {
		t.Errorf("AddTos should append to Mail.to")
	}
}

func TestCc(t *testing.T) {
	addr, err := mail.ParseAddress("Email Name<inbox@domain.tld>")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.Cc(addr)
	if len(m.cc) != 1 {
		t.Fatalf("To should append to Mail.to")
	}
}

func TestCcMultiple(t *testing.T) {
	addrs, err := mail.ParseAddressList("Email Name <email+1@email.com>, email+2@email.com")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.Cc(addrs...)
	if len(m.cc) != 2 {
		t.Errorf("AddTos should append to Mail.to")
	}
}

func TestBcc(t *testing.T) {
	addr, err := mail.ParseAddress("Email Name<inbox@domain.tld>")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.Bcc(addr)
	if len(m.bcc) != 1 {
		t.Fatalf("To should append to Mail.to")
	}
}

func TestBccMultiple(t *testing.T) {
	addrs, err := mail.ParseAddressList("Email Name <email+1@email.com>, email+2@email.com")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.Bcc(addrs...)
	if len(m.bcc) != 2 {
		t.Errorf("AddTos should append to Mail.to")
	}
}

func TestSubject(t *testing.T) {
	m := NewMail()
	testSubject := "Subject"
	m.Subject(testSubject)
	if m.subject != testSubject {
		t.Errorf("SetSubject should modify Mail.Subject")
	}
}

func TestSetText(t *testing.T) {
	m := NewMail()
	testText := "Text"
	m.Text(testText)
	if m.text != testText {
		t.Errorf("SetText should modify Mail.Text")
	}
}

func TestSetHTML(t *testing.T) {
	m := NewMail()
	testHTML := "<html></html>"
	m.HTML(testHTML)
	if m.html != testHTML {
		t.Errorf("SetHTML should modify Mail.HTML")
	}
}

func TestSetFrom(t *testing.T) {
	m := NewMail()
	testFrom, _ := mail.ParseAddress("Joe <email@email.com>")
	m.From(testFrom)
	if m.from == nil {
		t.Fatalf("From must set m.from")
	}
}

func TestSetReplyTo(t *testing.T) {

	addr, err := mail.ParseAddress("Email Name<inbox@domain.tld>")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMail()
	m.ReplyTo(addr)
	if m.replyTo == nil {
		t.Fatalf("ReplyTo should append to Mail.to")
	}
}

func TestSetDate(t *testing.T) {
	m := NewMail()
	date := "Today"
	m.SetDate(date)
	if m.date != date {
		t.Errorf("SetDate should modify Mail.Date")
	}
}

func TestSetRFCDate(t *testing.T) {
	m := NewMail()
	date := time.Now()
	m.SetRFCDate(date)
	if m.date == "" {
		t.Errorf("SetDate should fail if date is invalid RFC822")
	}
}

func TestAttach(t *testing.T) {
	m := NewMail()
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "THIS IS A TEST!")
	}))
	defer fakeServer.Close()
	res, err := http.Get(fakeServer.URL)
	if err != nil {
		t.Errorf("Fake server could'nt be reached")
	}
	defer res.Body.Close()
	err = m.Attach("Test", res.Body)
	if _, ok := m.files["Test"]; !ok {
		t.Errorf("Attachment not added")
	}
}

func TestAddAttachmentFail(t *testing.T) {
	m := NewMail()
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "THIS IS A TEST!")
	}))
	defer fakeServer.Close()
	res, err := http.Get(fakeServer.URL)
	res.Body.Close()
	if err != nil {
		t.Errorf("Fake server could'nt be reached")
	}

	err = m.Attach("Test", res.Body)
	if _, ok := m.files["Test"]; ok {
		t.Errorf("Attachment should not be added")
	}
}

func TestAddContentIds(t *testing.T) {
	m := NewMail()
	id, value := "id", "im a value"
	m.ContentID(id, value)
	if val, ok := m.content[id]; !ok && val != value {
		t.Errorf("ContentID failed to be added")
	}
}

func TestAddHeaders(t *testing.T) {
	m := NewMail()
	header, value := "id", "im a value"
	m.AddHeader(header, value)
	if val, ok := m.headers[header]; !ok && val != value {
		t.Errorf("Header failed to be added")
	}
}

func TestAddToSMTPAPI(t *testing.T) {
	m := NewMail()
	m.smtpHeaders.AddTo("example@email.com")
	if len(m.smtpHeaders.To) != 1 {
		t.Error("smptHeaders.To should be len of 1")
	}
}
