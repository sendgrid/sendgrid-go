package sendgrid

import (
	"net/mail"
	"net/url"
	"os"
	"testing"
)

func Test_Send(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	address, _ := mail.ParseAddress("John Doe <john@email.com>")
	message.AddRecipient(address)
	message.AddSubject("test")
	message.AddHTML("WIN")
	message.AddText("WIN")
	message.AddFrom("doe@email.com")
	message.AddSubstitution("subKey", "subValue")
	message.AddSection("testSection", "sectionValue")
	message.AddCategory("testCategory")
	message.AddUniqueArg("testUnique", "uniqueValue")
	message.AddFilter("testFilter", "filter", "filterValue")
	message.AddAttachmentStream("testFile", []byte("fileValue"))
	if reqUrl, e := sg.buildUrl(message); e != nil {
		t.Error(e)
	} else {
		reqUrl.Del("api_user")
		reqUrl.Del("api_key")
		testUrl := url.Values{}
		testUrl.Set("x-smtpapi", `{"sub":{"subKey":["subValue"]},"section":{"testSection":"sectionValue"},"category":["testCategory"],"unique_args":{"testUnique":"uniqueValue"},"filters":{"testFilter":{"settings":{"filter":"filterValue"}}}}`)
		testUrl.Set("to[]", "john@email.com")
		testUrl.Set("toname[]", "John Doe")
		testUrl.Set("html", "WIN")
		testUrl.Set("text", "WIN")
		testUrl.Set("subject", "test")
		testUrl.Set("files[testFile]", "fileValue")
		testUrl.Set("from", "doe@email.com")
		testUrl.Set("headers", "")
		testUrl.Set("replyto", "")
		if testUrl.Encode() == reqUrl.Encode() {
			t.Log("win")
		} else {
			t.Errorf("string builder failed:\n%s\n%s", testUrl.Encode(), reqUrl.Encode())
		}
	}
}
