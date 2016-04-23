package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

// Minimum required to send an email
func buildHelloEmail() string {
	m := mail.NewV3Mail()
	address := "dx@sendgrid.com"
	name := "DX"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	p := mail.NewPersonalization()
	address = "elmer.thomas@sendgrid.com"
	name = "Elmer Thomas"
	e = mail.NewEmail(name, address)
	p.AddTos(e)
	m.AddPersonalizations(p)

	m.Subject = "Hello World from the SendGrid Go Library"

	c := mail.NewContent("text/plain", "some text here")
	m.AddContent(c)

	c = mail.NewContent("text/html", "some html here")
	m.AddContent(c)

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

// Fully populated Mail object
func buildKitchenSink() string {
	m := mail.NewV3Mail()
	address := "dx@sendgrid.com"
	name := "DX"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.Subject = "Hello World from the SendGrid Go Library"

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Elmer Thomas", "elmer.thomas@sendgrid.com"),
		mail.NewEmail("Elmer Thomas Aliass", "elmer.thomas@gmail.com"),
	}
	p.AddTos(tos...)
	ccs := []*mail.Email{
		mail.NewEmail("Matt Bernier", "matt.bernier@sendgrid.com"),
		mail.NewEmail("Eric Shallock", "eric.shallock@sendgrid.com"),
	}
	p.AddCCs(ccs...)
	p.AddBCCs("dx+matt@sendgrid.com")
	p.AddBCCs("dx+eric@sendgrid.com")
	p.Subject = "Hello World from the Personalized SendGrid Go Library"
	p.SetHeader("X-Test", "test")
	p.SetHeader("X-Mock", "true")
	p.SetSubstitution("%name%", "Tim")
	p.SetSubstitution("%city%", "Riverside")
	p.SetCustomArg("user_id", "343")
	p.SetCustomArg("type", "marketing")
	p.SetSendAt(1461356286)
	m.AddPersonalizations(p)

	p2 := mail.NewPersonalization()
	tos2 := []*mail.Email{
		mail.NewEmail("Elmer Thomas", "elmer.thomas@sendgrid.com"),
		mail.NewEmail("Elmer Thomas Aliass", "elmer.thomas@gmail.com"),
	}
	p2.AddTos(tos2...)
	ccs2 := []*mail.Email{
		mail.NewEmail("Matt Bernier", "matt.bernier@sendgrid.com"),
		mail.NewEmail("Eric Shallock", "eric.shallock@sendgrid.com"),
	}
	p2.AddCCs(ccs2...)
	p2.AddBCCs("dx+matt@sendgrid.com")
	p2.AddBCCs("dx+eric@sendgrid.com")
	p2.Subject = "Hello World from the Personalized SendGrid Go Library"
	p2.SetHeader("X-Test", "test")
	p2.SetHeader("X-Mock", "true")
	p2.SetSubstitution("%name%", "Tim")
	p2.SetSubstitution("%city%", "Riverside")
	p2.SetCustomArg("user_id", "343")
	p2.SetCustomArg("type", "marketing")
	p2.SetSendAt(1461356286)
	m.AddPersonalizations(p2)

	c := mail.NewContent("text/plain", "some text here")
	m.AddContent(c)

	c = mail.NewContent("text/html", "some html here")
	m.AddContent(c)

	a := mail.NewAttachment()
	a.SetContent("TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4gQ3JhcyBwdW12")
	a.SetType("application/pdf")
	a.SetFilename("balance_001.pdf")
	a.SetDisposition("attachment")
	a.SetContentID("Balance Sheet")
	m.AddAttachment(a)

	a2 := mail.NewAttachment()
	a2.SetContent("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")
	a2.SetType("image/png")
	a2.SetFilename("banner.png")
	a2.SetDisposition("inline")
	a2.SetContentID("Banner")
	m.AddAttachment(a2)

	m.SetTemplateID("13b8f94f-bcae-4ec6-b752-70d6cb59f932")

	m.AddSection("%section1%", "Substitution Text for Section 1")
	m.AddSection("%section2%", "Substitution Text for Section 2")

	m.SetHeader("X-Test1", "1")
	m.SetHeader("X-Test2", "2")

	m.AddCategories("May")
	m.AddCategories("2016")

	m.SetCustomArg("campaign", "welcome")
	m.SetCustomArg("weekday", "morning")

	m.SetSendAt(1461356286)

	asm := mail.NewASM()
	asm.SetGroupID(99)
	asm.AddGroupsToDisplay(99)
	m.SetASM(asm)

	m.SetBatchID("sendgrid_batch_id")

	m.SetIPPoolID("23")

	mailSettings := mail.NewMailSettings()
	bccSettings := mail.NewBCCSetting()
	bccSettings.SetEnable(true)
	bccSettings.SetEmail("dx@sendgrid.com")
	mailSettings.SetBCC(bccSettings)
	sandBoxMode := mail.NewSetting(true)
	mailSettings.SetSandboxMode(sandBoxMode)
	bypassListManagement := mail.NewSetting(true)
	mailSettings.SetBypassListManagement(bypassListManagement)
	footerSetting := mail.NewFooterSetting()
	footerSetting.SetText("Footer Text")
	footerSetting.SetEnable(true)
	footerSetting.SetHTML("<html><body>Footer Text</body></html>")
	mailSettings.SetFooter(footerSetting)
	spamCheckSetting := mail.NewSpamCheckSetting()
	spamCheckSetting.SetEnable(true)
	spamCheckSetting.SetSpamThreshold(1)
	spamCheckSetting.SetPostToURL("https://spamcatcher.sendgrid.com")
	mailSettings.SetSpamCheckSettings(spamCheckSetting)
	m.SetMailSettings(mailSettings)

	trackingSettings := mail.NewTrackingSettings()
	clickTrackingSettings := mail.NewClickTrackingSetting()
	clickTrackingSettings.SetEnable(true)
	clickTrackingSettings.SetEnableText(true)
	trackingSettings.SetClickTracking(clickTrackingSettings)
	openTrackingSetting := mail.NewOpenTrackingSetting()
	openTrackingSetting.SetEnable(true)
	openTrackingSetting.SetSubstitutionTag("Optional tag to replace with the open image in the body of the message")
	trackingSettings.SetOpenTracking(openTrackingSetting)
	subscriptionTrackingSetting := mail.NewSubscriptionTrackingSetting()
	subscriptionTrackingSetting.SetEnable(true)
	subscriptionTrackingSetting.SetText("text to insert into the text/plain portion of the message")
	subscriptionTrackingSetting.SetHTML("<html><body>html to insert into the text/html portion of the message</body></html>")
	subscriptionTrackingSetting.SetSubstitutionTag("Optional tag to replace with the open image in the body of the message")
	trackingSettings.SetSubscriptionTracking(subscriptionTrackingSetting)
	googleAnalyticsSetting := mail.NewGaSetting()
	googleAnalyticsSetting.SetEnable(true)
	googleAnalyticsSetting.SetCampaignSource("some source")
	googleAnalyticsSetting.SetCampaignTerm("some term")
	googleAnalyticsSetting.SetCampaignContent("some content")
	googleAnalyticsSetting.SetCampaignName("some name")
	googleAnalyticsSetting.SetCampaignMedium("some medium")
	trackingSettings.SetGoogleAnalytics(googleAnalyticsSetting)
	m.SetTrackingSettings(trackingSettings)

	replyToEmail := mail.NewEmail("Mr. Elmer Thomas", "dx+reply@sendgrid.com")
	m.SetReplyTo(replyToEmail)

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func baselineExample() {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/mail/send", "https://api.sendgrid.com", "v3")
	request.Method = "POST"
	var requestBody = []byte(buildHelloEmail())
	request.RequestBody = requestBody
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}
}

func kitchenSinkExample() {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/mail/send", "https://api.sendgrid.com", "v3")
	request.Method = "POST"
	var requestBody = []byte(buildKitchenSink())
	request.RequestBody = requestBody
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.ResponseBody)
		fmt.Println(response.ResponseHeaders)
	}
}

func main() {
	baselineExample()
	kitchenSinkExample()
}
