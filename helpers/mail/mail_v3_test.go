package mail

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestV3NewMail will test New mail method
func TestV3NewMail(t *testing.T) {
	m := NewV3Mail()

	assert.NotNil(t, m, "NewV3Mail() shouldn't return nil")
	assert.NotNil(t, m.Personalizations, "Personalizations shouldn't be nil")
	assert.NotNil(t, m.Attachments, "Attachments shouldn't be nil")
	assert.NotNil(t, m.Content, "Content shouldn't be nil")
}

func TestV3NewMailInit(t *testing.T) {
	from := NewEmail("Example User", "test@example.com")
	subject := "Hello World from the Twilio SendGrid Go Library"
	to := NewEmail("Example User", "test@example.com")
	content := NewContent("text/plain", "some text here")
	m := NewV3MailInit(from, subject, to, content)

	assert.NotNil(t, m, "NewV3MailInit() shouldn't return nil")
	assert.NotNil(t, m.From, "From shouldn't be nil")
	assert.NotNil(t, m.Personalizations, "Personalizations shouldn't be nil")
	assert.NotNil(t, m.Content, "Content shouldn't be nil")
}

func TestV3AddPersonalizations(t *testing.T) {
	numOfPersonalizations := rand.New(rand.NewSource(99)).Intn(10)
	personalizations := make([]*Personalization, 0)
	for i := 0; i < numOfPersonalizations; i++ {
		personalizations = append(personalizations, NewPersonalization())
	}

	m := NewV3Mail()
	m.AddPersonalizations(personalizations...)

	assert.Equal(t, numOfPersonalizations, len(m.Personalizations), fmt.Sprintf("Mail should have %d personalizations, got %d personalizations", len(personalizations), len(m.Personalizations)))
}

func TestV3AddContent(t *testing.T) {
	numOfContent := 2
	content := make([]*Content, 0)
	for i := 0; i < numOfContent; i++ {
		content = append(content, NewContent("type", "value"))
	}

	m := NewV3Mail()
	m.AddContent(content...)

	assert.Equal(t, numOfContent, len(m.Content), fmt.Sprintf("Mail should have %d contents, got %d contents", numOfContent, len(m.Content)))
}

func TestV3AddAttachment(t *testing.T) {
	numOfAttachments := 2
	attachment := make([]*Attachment, 0)
	for i := 0; i < numOfAttachments; i++ {
		attachment = append(attachment, NewAttachment())
	}

	m := NewV3Mail()
	m.AddAttachment(attachment...)

	assert.Equal(t, numOfAttachments, len(m.Attachments), fmt.Sprintf("Mail should have %d attachments, got %d attachments", numOfAttachments, len(m.Attachments)))
}

func TestV3SetFrom(t *testing.T) {
	m := NewV3Mail()

	address := "test@example.com"
	name := "Test User"
	e := NewEmail(name, address)
	m.SetFrom(e)

	assert.Equal(t, name, m.From.Name, fmt.Sprintf("name should be %s, got %s", name, e.Name))
	assert.Equal(t, address, m.From.Address, fmt.Sprintf("address should be %s, got %s", address, e.Address))
}

func TestV3SetReplyTo(t *testing.T) {
	m := NewV3Mail()

	address := "test@example.com"
	name := "Test User"
	e := NewEmail(name, address)
	m.SetReplyTo(e)

	assert.Equal(t, name, m.ReplyTo.Name, fmt.Sprintf("name should be %s, got %s", name, e.Name))
	assert.Equal(t, address, m.ReplyTo.Address, fmt.Sprintf("address should be %s, got %s", address, e.Address))
}

func TestV3SetTemplateID(t *testing.T) {
	m := NewV3Mail()

	templateID := "templateabcd12345"

	m.SetTemplateID(templateID)

	assert.Equal(t, templateID, m.TemplateID, fmt.Sprintf("templateID should be %s, got %s", templateID, m.TemplateID))
}

func TestV3AddSection(t *testing.T) {
	m := NewV3Mail()

	sectionKey := "key"
	sectionValue := "value"

	m.AddSection(sectionKey, sectionValue)

	v, ok := m.Sections[sectionKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in Sections map", sectionKey)) {
		assert.Equal(t, v, sectionValue, fmt.Sprintf("value should be %s, got %s", sectionValue, v))
	}
}

func TestV3SetHeader(t *testing.T) {
	m := NewV3Mail()

	headerKey := "key"
	headerValue := "value"

	m.SetHeader(headerKey, headerValue)

	v, ok := m.Headers[headerKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in Headers map", headerKey)) {
		assert.Equal(t, v, headerValue, fmt.Sprintf("value should be %s, got %s", headerValue, v))
	}
}

func TestV3AddCategory(t *testing.T) {
	m := NewV3Mail()

	categories := []string{"cats", "dogs", "hamburgers", "cheeseburgers"}

	m.AddCategories(categories...)

	assert.Equal(t, len(categories), len(m.Categories), fmt.Sprintf("Length of Categories should be %d, got %d", len(categories), len(m.Categories)))
}

func TestV3SetCustomArg(t *testing.T) {
	m := NewV3Mail()

	customArgKey := "key"
	customArgValue := "value"

	m.SetCustomArg(customArgKey, customArgValue)

	v, ok := m.CustomArgs[customArgKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in CustomArgs map", customArgKey)) {
		assert.Equal(t, v, customArgValue, fmt.Sprintf("value should be %s, got %s", customArgValue, v))
	}
}

func TestV3SetSendAt(t *testing.T) {
	m := NewV3Mail()
	sendAt := time.Now().Second()

	m.SetSendAt(sendAt)
	assert.Equal(t, sendAt, m.SendAt, fmt.Sprintf("SendAt should be %d, got %d", sendAt, m.SendAt))
}

func TestV3SetBatchID(t *testing.T) {
	m := NewV3Mail()
	batchID := "batchID123455"

	m.SetBatchID(batchID)
	assert.Equal(t, batchID, m.BatchID, fmt.Sprintf("BatchID should be %s, got %s", batchID, m.BatchID))
}

func TestV3SetIPPoolID(t *testing.T) {
	m := NewV3Mail()
	ipPoolID := "42"

	m.SetIPPoolID(ipPoolID)
	assert.Equal(t, ipPoolID, m.IPPoolID, fmt.Sprintf("IP Pool ID should be %s, got %s", ipPoolID, m.IPPoolID))
}

func TestV3SetASM(t *testing.T) {
	m := NewV3Mail()
	asm := NewASM()
	groupID := 1
	groupsToDisplay := []int{1, 2, 3, 4}
	asm.SetGroupID(groupID)
	asm.AddGroupsToDisplay(groupsToDisplay...)

	m.SetASM(asm)

	assert.Equal(t, groupID, m.Asm.GroupID, fmt.Sprintf("GroupID should be %d, got %d", groupID, m.Asm.GroupID))
	assert.Equal(t, groupsToDisplay, m.Asm.GroupsToDisplay, fmt.Sprintf("Length of GroupsToDisplay should be %d, got %d", len(groupsToDisplay), len(m.Asm.GroupsToDisplay)))
}

func TestV3SetMailSettings(t *testing.T) {
	m := NewV3Mail()
	ms := NewMailSettings()
	ms.SetBCC(NewBCCSetting().SetEnable(true))
	m.SetMailSettings(ms)

	assert.NotNil(t, m.MailSettings, "Mail Settings should not be nil")
	assert.True(t, *m.MailSettings.BCC.Enable, "BCC should be enabled in Mail Settings")
}

func TestV3SetTrackingSettings(t *testing.T) {
	m := NewV3Mail()
	ts := NewTrackingSettings()
	n := NewClickTrackingSetting()
	n.SetEnable(true)
	n.SetEnableText(true)
	ts.SetClickTracking(n)
	m.SetTrackingSettings(ts)

	assert.NotNil(t, m.TrackingSettings, "Tracking Settings should not be nil")
	assert.True(t, *m.TrackingSettings.ClickTracking.Enable, "Click Tracking should be enabled")
}

func TestV3NewPersonalization(t *testing.T) {
	p := NewPersonalization()

	assert.NotNil(t, p, "NewPersonalization() shouldn't return nil")

	assert.NotNil(t, p.To, "To should't be nil")
	assert.Equal(t, 0, len(p.To), "Length of p.To should be 0")

	assert.NotNil(t, p.CC, "CC should't be nil")
	assert.Equal(t, 0, len(p.CC), "Length of p.CCs should be 0")

	assert.NotNil(t, p.BCC, "BCC should't be nil")
	assert.Equal(t, 0, len(p.BCC), "Length of p.BCC should be 0")

	assert.NotNil(t, p.Headers, "Headers should't be nil")
	assert.Equal(t, 0, len(p.Headers), "Length of p.Headers should be 0")

	assert.NotNil(t, p.Substitutions, "Substitutions should't be nil")
	assert.Equal(t, 0, len(p.Substitutions), "Length of p.Substitutions should be 0")

	assert.NotNil(t, p.CustomArgs, "CustomArgs should't be nil")
	assert.Equal(t, 0, len(p.CustomArgs), "Length of p.CustomArgs should be 0")

	assert.NotNil(t, p.Categories, "Categories should't be nil")
	assert.Equal(t, 0, len(p.Categories), "Length of p.Categories should be 0")
}

func TestV3PersonalizationAddTos(t *testing.T) {
	tos := []*Email{
		NewEmail("Example User", "test@example.com"),
		NewEmail("Example User", "test@example.com"),
	}

	p := NewPersonalization()
	p.AddTos(tos...)

	assert.Equal(t, len(tos), len(p.To), fmt.Sprintf("length of To should be %d, got %d", len(tos), len(p.To)))
}

func TestV3PersonalizationAddCCs(t *testing.T) {
	ccs := []*Email{
		NewEmail("Example User", "test@example.com"),
		NewEmail("Example User", "test@example.com"),
	}

	p := NewPersonalization()
	p.AddCCs(ccs...)

	assert.Equal(t, len(ccs), len(p.CC), fmt.Sprintf("length of CC should be %d, got %d", len(ccs), len(p.CC)))
}

func TestV3PersonalizationAddBCCs(t *testing.T) {
	bccs := []*Email{
		NewEmail("Example User", "test@example.com"),
		NewEmail("Example User", "test@example.com"),
	}

	p := NewPersonalization()
	p.AddBCCs(bccs...)

	assert.Equal(t, len(bccs), len(p.BCC), fmt.Sprintf("length of BCC should be %d, got %d", len(bccs), len(p.BCC)))
}

func TestV3PersonalizationSetHeader(t *testing.T) {
	p := NewPersonalization()

	headerKey := "key"
	headerValue := "value"

	p.SetHeader(headerKey, headerValue)

	v, ok := p.Headers[headerKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in Headers map", headerKey)) {
		assert.Equal(t, headerValue, v, fmt.Sprintf("value should be %s, got %s", headerValue, v))
	}
}

func TestV3PersonalizationSetSubstitution(t *testing.T) {
	p := NewPersonalization()

	substitutionKey := "key"
	substitutionValue := "value"

	p.SetSubstitution(substitutionKey, substitutionValue)

	v, ok := p.Substitutions[substitutionKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in Substitutions map", substitutionKey)) {
		assert.Equal(t, substitutionValue, v, fmt.Sprintf("value should be %s, got %s", substitutionValue, v))
	}
}

func TestV3PersonalizationSetCustomArg(t *testing.T) {
	p := NewPersonalization()

	customArgKey := "key"
	customArgValue := "value"

	p.SetCustomArg(customArgKey, customArgValue)

	v, ok := p.CustomArgs[customArgKey]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in CustomArgs map", customArgKey)) {
		assert.Equal(t, customArgValue, v, fmt.Sprintf("value should be %s, got %s", customArgValue, v))
	}
}

func TestV3PersonalizationSetDynamicTemplateData(t *testing.T) {
	p := NewPersonalization()

	dynamicTemplateDataKey0 := "simpleString"
	dynamicTemplateDataValue0 := "value"
	p.SetDynamicTemplateData(dynamicTemplateDataKey0, dynamicTemplateDataValue0)

	v, ok := p.DynamicTemplateData[dynamicTemplateDataKey0]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in DynamictemplateData map", dynamicTemplateDataKey0)) {
		assert.Equal(t, dynamicTemplateDataValue0, v, fmt.Sprintf("value should be %s, got %s", dynamicTemplateDataValue0, v))
	}

	dynamicTemplateDataKey1 := "arr"
	dynamicTemplateDataValue1 := "[true, false, true]"
	p.SetDynamicTemplateData(dynamicTemplateDataKey1, dynamicTemplateDataValue1)

	v, ok = p.DynamicTemplateData[dynamicTemplateDataKey1]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in DynamictemplateData map", dynamicTemplateDataKey1)) {
		assert.Equal(t, dynamicTemplateDataValue1, v, fmt.Sprintf("value should be %s, got %s", dynamicTemplateDataValue1, v))
	}

	dynamicTemplateDataKey2 := "obj"
	dynamicTemplateDataValue2 := map[string]string{
		"dynamic":    "templates",
		"dynamicArr": "[]int{0, 1, 2}",
		"bool":       "false",
		"int":        "10",
	}
	p.SetDynamicTemplateData(dynamicTemplateDataKey2, dynamicTemplateDataValue2)

	v, ok = p.DynamicTemplateData[dynamicTemplateDataKey2]
	if !assert.True(t, ok, fmt.Sprintf("key %s not found in DynamictemplateData map", dynamicTemplateDataKey2)) {
		assert.Equal(t, dynamicTemplateDataValue2, v, fmt.Sprintf("value should be %s, got %s", dynamicTemplateDataValue2, v))
	}
}

func TestV3PersonalizationSetSendAt(t *testing.T) {
	p := NewPersonalization()
	sendAt := time.Now().Second()

	p.SetSendAt(sendAt)
	assert.Equal(t, sendAt, p.SendAt, fmt.Sprintf("sendat should be %d, got %d", sendAt, p.SendAt))
}

func TestV3NewAttachment(t *testing.T) {
	assert.NotNil(t, NewAttachment(), "NewAttachment() shouldn't return nil")
}

func TestV3AttachmentSetContent(t *testing.T) {
	content := "somebase64encodedcontent"
	a := NewAttachment().SetContent(content)

	assert.Equal(t, content, a.Content, fmt.Sprintf("Content should be %s, got %s", content, a.Content))
}

func TestV3AttachmentSetType(t *testing.T) {
	contentType := "pdf"
	a := NewAttachment().SetType(contentType)

	assert.Equal(t, contentType, a.Type, fmt.Sprintf("ContentType should be %s, got %s", contentType, a.Type))
}

func TestV3AttachmentSetContentID(t *testing.T) {
	contentID := "contentID"
	a := NewAttachment().SetContentID(contentID)

	assert.Equal(t, contentID, a.ContentID, fmt.Sprintf("ContentID should be %s, got %s", contentID, a.ContentID))
}

func TestV3AttachmentSetDisposition(t *testing.T) {
	disposition := "inline"
	a := NewAttachment().SetDisposition(disposition)

	assert.Equal(t, disposition, a.Disposition, fmt.Sprintf("Disposition should be %s, got %s", disposition, a.Disposition))
}

func TestV3AttachmentSetFilename(t *testing.T) {
	filename := "mydoc.pdf"
	a := NewAttachment().SetFilename(filename)

	assert.Equal(t, filename, a.Filename, fmt.Sprintf("Filename should be %s, got %s", filename, a.Filename))
}

func TestV3NewASM(t *testing.T) {
	assert.NotNil(t, NewASM(), "NewASM() should not return nil")
}

func TestV3ASMSetGroupID(t *testing.T) {
	groupID := 1
	a := NewASM().SetGroupID(groupID)

	assert.Equal(t, groupID, a.GroupID, fmt.Sprintf("GroupID should be %v, got %v", groupID, a.GroupID))
}

func TestV3ASMSetGroupstoDisplay(t *testing.T) {
	groupsToDisplay := []int{1, 2, 3, 4}
	a := NewASM().AddGroupsToDisplay(groupsToDisplay...)

	assert.Equal(t, len(groupsToDisplay), len(a.GroupsToDisplay), fmt.Sprintf("Length of GroupsToDisplay should be %d, got %d", groupsToDisplay, a.GroupsToDisplay))
}

func TestV3NewMailSettings(t *testing.T) {
	assert.NotNil(t, NewMailSettings(), "NewMailSettings() shouldn't return nil")
}

func TestV3MailSettingsSetBCC(t *testing.T) {
	m := NewMailSettings().SetBCC(NewBCCSetting().SetEnable(true))

	assert.NotNil(t, m.BCC, "BCC should not be nil")
	assert.True(t, *m.BCC.Enable, "BCC should be enabled")
}

func TestV3MailSettingsSetBypassListManagement(t *testing.T) {
	m := NewMailSettings().SetBypassListManagement(NewSetting(true))

	assert.NotNil(t, m.BypassListManagement, "BypassListManagement should not be nil")
	assert.True(t, *m.BypassListManagement.Enable, "BypassListManagement should be enabled")
}

func TestV3MailSettingsSetSandboxMode(t *testing.T) {
	m := NewMailSettings().SetSandboxMode(NewSetting(true))

	assert.NotNil(t, m.SandboxMode, "SandboxMode should not be nil")
	assert.True(t, *m.SandboxMode.Enable, "SandboxMode should be enabled")
}

func TestV3MailSettingsSpamCheckSettings(t *testing.T) {
	m := NewMailSettings()
	s := NewSpamCheckSetting()
	s.SetEnable(true)
	s.SetPostToURL("http://test.com")
	s.SetSpamThreshold(1)
	m.SetSpamCheckSettings(s)

	assert.True(t, *m.SpamCheckSetting.Enable, "SpamCheckSetting should be enabled")
	assert.NotNil(t, m.SpamCheckSetting.PostToURL, "PostToURL should not be nil")
	assert.Equal(t, 1, m.SpamCheckSetting.SpamThreshold, "Spam threshold should be 1")
}

func TestV3MailSettingsSetFooter(t *testing.T) {
	m := NewMailSettings().SetFooter(NewFooterSetting().SetEnable(true))

	assert.NotNil(t, m.Footer, "Footer should not be nil")
	assert.True(t, *m.Footer.Enable, "Footer should be enabled")
}

func TestV3NewTrackingSettings(t *testing.T) {
	assert.NotNil(t, NewTrackingSettings(), "NewTrackingSettings() shouldn't return nil")
}

func TestV3TrackingSettingsSetClickTracking(t *testing.T) {
	n := NewClickTrackingSetting()
	n.SetEnable(true)
	n.SetEnableText(true)
	ts := NewTrackingSettings().SetClickTracking(n)

	assert.NotNil(t, ts.ClickTracking, "Click Tracking shouldn't return nil")
	assert.True(t, *ts.ClickTracking.Enable, "ClickTracking should be enabled")
}

func TestV3TrackingSettingsSetOpenTracking(t *testing.T) {
	substitutionTag := "subTag"
	ts := NewTrackingSettings().SetOpenTracking(NewOpenTrackingSetting().SetEnable(true).SetSubstitutionTag(substitutionTag))

	assert.NotNil(t, ts.OpenTracking, "Open Tracking should not be nil")
	assert.True(t, *ts.OpenTracking.Enable, "OpenTracking should be enabled")
	assert.Equal(t, substitutionTag, ts.OpenTracking.SubstitutionTag, fmt.Sprintf("Substitution Tag should be %s, got %s", substitutionTag, ts.OpenTracking.SubstitutionTag))
}

func TestV3TrackingSettingsSetSubscriptionTracking(t *testing.T) {
	ts := NewTrackingSettings().SetSubscriptionTracking(NewSubscriptionTrackingSetting())
	assert.NotNil(t, ts.SubscriptionTracking, "SubscriptionTracking should not be nil")
}

func TestV3TrackingSettingsSetGoogleAnalytics(t *testing.T) {
	campaignName := "campaign1"
	campaignTerm := "campaign1_term"
	campaignSource := "campaign1_source"
	campaignContent := "campaign1_content"
	campaignMedium := "campaign1_medium"

	ts := NewTrackingSettings().SetGoogleAnalytics(NewGaSetting().SetCampaignName(campaignName).SetCampaignTerm(campaignTerm).SetCampaignSource(campaignSource).SetCampaignContent(campaignContent).SetCampaignMedium(campaignMedium).SetEnable(true))

	assert.NotNil(t, ts.GoogleAnalytics, "GoogleAnalytics should not be nil")

	assert.Equal(t, campaignName, ts.GoogleAnalytics.CampaignName, fmt.Sprintf("CampaignName should be %s, got %s", campaignName, ts.GoogleAnalytics.CampaignName))

	assert.Equal(t, campaignTerm, ts.GoogleAnalytics.CampaignTerm, fmt.Sprintf("CampaignTerm should be %s, got %s", campaignTerm, ts.GoogleAnalytics.CampaignTerm))

	assert.Equal(t, campaignSource, ts.GoogleAnalytics.CampaignSource, fmt.Sprintf("CampaignSource should be %s, got %s", campaignSource, ts.GoogleAnalytics.CampaignSource))

	assert.Equal(t, campaignContent, ts.GoogleAnalytics.CampaignContent, fmt.Sprintf("CampaignContent should be %s, got %s", campaignContent, ts.GoogleAnalytics.CampaignContent))

	assert.Equal(t, campaignMedium, ts.GoogleAnalytics.CampaignMedium, fmt.Sprintf("CampaignMedium should be %s, got %s", campaignMedium, ts.GoogleAnalytics.CampaignMedium))
}

func TestV3NewBCCSetting(t *testing.T) {
	assert.NotNil(t, NewBCCSetting(), "NewBCCSetting() shouldn't return nil")
}

func TestV3BCCSettingSetEnable(t *testing.T) {
	b := NewBCCSetting().SetEnable(true)

	assert.True(t, *b.Enable, "BCCSetting should be enabled")
}

func TestV3BCCSettingSetEmail(t *testing.T) {
	address := "joe@schmoe.net"
	b := NewBCCSetting().SetEmail(address)

	assert.NotNil(t, b.Email, "Email should not be empty")
}

func TestV3NewFooterSetting(t *testing.T) {
	assert.NotNil(t, NewFooterSetting(), "NewFooterSetting() shouldn't return nil")
}

func TestV3FooterSettingSetEnable(t *testing.T) {
	f := NewFooterSetting().SetEnable(true)

	assert.True(t, *f.Enable, "FooterSetting should be enabled")
}

func TestV3FooterSettingSetText(t *testing.T) {
	text := "some test here"
	f := NewFooterSetting().SetText(text)

	assert.Equal(t, text, f.Text, fmt.Sprintf("Text should be %s, got %s", text, f.Text))
}

func TestV3FooterSettingSetHtml(t *testing.T) {
	html := "<h1>some html</h1>"
	f := NewFooterSetting().SetHTML(html)

	assert.Equal(t, html, f.Html, fmt.Sprintf("Html should be %s, got %s", html, f.Html))
}

func TestV3NewOpenTrackingSetting(t *testing.T) {
	assert.NotNil(t, NewOpenTrackingSetting(), "NewOpenTrackingSetting() shouldn't return nil")
}

func TestV3OpenTrackingSettingSetEnable(t *testing.T) {
	f := NewOpenTrackingSetting().SetEnable(true)

	assert.True(t, *f.Enable, "OpenTrackingSetting should be enabled")
}

func TestV3OpenTrackingSettingSetSubstitutionTag(t *testing.T) {
	substitutionTag := "tag"
	f := NewOpenTrackingSetting().SetSubstitutionTag(substitutionTag)

	assert.Equal(t, substitutionTag, f.SubstitutionTag, fmt.Sprintf("SubstitutionTag should be %s, got %s", substitutionTag, f.SubstitutionTag))
}

func TestV3NewSubscriptionTrackingSetting(t *testing.T) {
	assert.NotNil(t, NewSubscriptionTrackingSetting(), "NewSubscriptionTrackingSetting() shouldn't return nil")
}

func TestV3NewSubscriptionTrackingSetEnable(t *testing.T) {
	s := NewSubscriptionTrackingSetting().SetEnable(true)

	assert.True(t, *s.Enable, "SubscriptionTracking should be enabled")
}

func TestV3NewSubscriptionTrackingSetSubstitutionTag(t *testing.T) {
	substitutionTag := "subTag"
	s := NewSubscriptionTrackingSetting().SetSubstitutionTag(substitutionTag)

	assert.Equal(t, substitutionTag, s.SubstitutionTag, fmt.Sprintf("SubstitutionTag should be %s, got %s", substitutionTag, s.SubstitutionTag))
}

func TestV3NewSubscriptionTrackingSetText(t *testing.T) {
	text := "text"
	s := NewSubscriptionTrackingSetting().SetText(text)

	assert.Equal(t, text, s.Text, fmt.Sprintf("Text should be %s, got %s", text, s.Text))
}

func TestV3NewSubscriptionTrackingSetHtml(t *testing.T) {
	html := "<h1>hello</h1>"
	s := NewSubscriptionTrackingSetting().SetHTML(html)

	assert.Equal(t, html, s.Html, fmt.Sprintf("Html should be %s, got %s", html, s.Html))
}

func TestV3NewGaSetting(t *testing.T) {
	assert.NotNil(t, NewGaSetting(), "NewGaSetting() shouldn't return nil")
}

func TestV3GaSettingSetCampaignName(t *testing.T) {
	campaignName := "campaign1"
	g := NewGaSetting().SetCampaignName(campaignName)

	assert.Equal(t, campaignName, g.CampaignName, fmt.Sprintf("CampaignName should be %s, got %s", campaignName, g.CampaignName))
}

func TestV3GaSettingSetCampaignTerm(t *testing.T) {
	campaignTerm := "campaign1_term"
	g := NewGaSetting().SetCampaignTerm(campaignTerm)

	assert.Equal(t, campaignTerm, g.CampaignTerm, fmt.Sprintf("CampaignTerm should be %s, got %s", campaignTerm, g.CampaignTerm))
}

func TestV3GaSettingSetCampaignSource(t *testing.T) {
	campaignSource := "campaign1_source"
	g := NewGaSetting().SetCampaignSource(campaignSource)

	assert.Equal(t, campaignSource, g.CampaignSource, fmt.Sprintf("CampaignSource should be %s, got %s", campaignSource, g.CampaignSource))
}

func TestV3GaSettingSetCampaignContent(t *testing.T) {
	campaignContent := "campaign1_content"
	g := NewGaSetting().SetCampaignContent(campaignContent)

	assert.Equal(t, campaignContent, g.CampaignContent, fmt.Sprintf("CampaignContent should be %s, got %s", campaignContent, g.CampaignContent))
}

func TestV3NewSetting(t *testing.T) {
	s := NewSetting(true)

	assert.NotNil(t, s, "NewSetting() shouldn't return nil")
	assert.True(t, *s.Enable, "NewSetting(true) should return a setting with Enabled = true")
}

func TestV3NewEmail(t *testing.T) {
	name := "Johnny"
	address := "Johnny@rocket.io"

	e := NewEmail(name, address)

	assert.Equal(t, name, e.Name, fmt.Sprintf("Name should be %s, got %s", name, e.Name))
	assert.Equal(t, address, e.Address, fmt.Sprintf("Address should be %s, got %s", address, e.Address))
}

func TestV3NewSingleEmail(t *testing.T) {
	from := NewEmail("Example User", "test@example.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	message := NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	assert.NotNil(t, message, "NewV3MailInit() shouldn't return nil")
	assert.NotNil(t, message.From, "From shouldn't return nil")
	assert.Equal(t, subject, message.Subject, fmt.Sprintf("Subject should be %s, got %s", subject, message.Subject))
	assert.NotNil(t, message.Content, "Content shouldn't be nil")
}

func TestV3NewSingleEmailWithEmptyHTMLContent(t *testing.T) {
	from := NewEmail("Example User", "test@example.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"

	message := NewSingleEmail(from, subject, to, plainTextContent, "")

	assert.NotNil(t, message, "NewV3MailInit() shouldn't return nil")
	assert.NotNil(t, message.From, "From shouldn't be nil")
	assert.Equal(t, message.Subject, subject, fmt.Sprintf("Subject should be %s, got %s", subject, message.Subject))
	assert.NotNil(t, message.Content, "Content shouldn't be nil")
}

func TestV3NewClickTrackingSetting(t *testing.T) {
	c := NewClickTrackingSetting()
	c.SetEnable(true)
	c.SetEnableText(false)

	assert.True(t, *c.Enable, "Click Tracking should be enabled")
	assert.False(t, *c.EnableText, "Enable Text should not be enabled")
}

func TestV3NewSpamCheckSetting(t *testing.T) {
	spamThreshold := 8
	postToURL := "http://myurl.com"
	s := NewSpamCheckSetting()
	s.SetEnable(true)
	s.SetSpamThreshold(spamThreshold)
	s.SetPostToURL(postToURL)

	assert.True(t, *s.Enable, "SpamCheck should be enabled")
	assert.Equal(t, spamThreshold, s.SpamThreshold, fmt.Sprintf("SpamThreshold should be %d, got %d", spamThreshold, s.SpamThreshold))
	assert.Equal(t, postToURL, s.PostToURL, fmt.Sprintf("PostToURL should be %s, got %s", postToURL, s.PostToURL))
}

func TestV3NewSandboxModeSetting(t *testing.T) {
	spamCheck := NewSpamCheckSetting()
	spamCheck.SetEnable(true)
	spamCheck.SetSpamThreshold(1)
	spamCheck.SetPostToURL("http://wwww.google.com")
	s := NewSandboxModeSetting(true, true, spamCheck)

	assert.True(t, *s.Enable, "Sandbox Mode should be enabled")
	assert.True(t, *s.ForwardSpam, "ForwardSpam should be enabled")
	assert.NotNil(t, s.SpamCheck, "SpamCheck should not be nil")
}

func TestParseEmail(t *testing.T) {
	e, err := ParseEmail("example example <example@example.com>")
	if err != nil {
		t.Error("Email should have been parsed successfully")
	}
	expectedName := "example example"
	if e.Name != expectedName {
		t.Errorf("Expect email with name %s but got %s", expectedName, e.Name)
	}
	expectedAddress := "example@example.com"
	if e.Address != expectedAddress {
		t.Errorf("Expect email with address %s but got %s", expectedAddress, e.Address)
	}
}

func TestParseInvalidEmail(t *testing.T) {
	_, err := ParseEmail("example example <example/example.com>")
	if err == nil {
		t.Error("Expected an error to be thrown from ParseEmail")
	}
}
