package mail

import (
	"encoding/json"
	"log"
)

// SGMailV3
type SGMailV3 struct {
	From             *Email             `json:"from,omitempty"`
	Subject          string             `json:"subject,omitempty"`
	Personalizations []*Personalization `json:"personalizations,omitempty"`
	Content          []*Content         `json:"content,omitempty"`
	Attachments      []*Attachment      `json:"attachments,omitempty"`
	TemplateID       string             `json:"template_id,omitempty"`
	Sections         map[string]string  `json:"sections,omitempty"`
	Headers          map[string]string  `json:"headers,omitempty"`
	Categories       []string           `json:"categories,omitempty"`
	CustomArgs       map[string]string  `json:"custom_args,omitempty"`
	SendAt           int                `json:"send_at,omitempty"`
	BatchID          string             `json:"batch_id,omitempty"`
	Asm              *Asm               `json:"asm,omitempty"`
	IPPoolID         string             `json:"ip_pool_name,omitempty"`
	MailSettings     *MailSettings      `json:"mail_settings,omitempty"`
	TrackingSettings *TrackingSettings  `json:"tracking_settings,omitempty"`
	ReplyTo          *Email             `json:"reply_to,omitempty"`
}

// Personalization
type Personalization struct {
	To            []*Email          `json:"to,omitempty"`
	CC            []*Email          `json:"cc,omitempty"`
	BCC           []*Email          `json:"bcc,omitempty"`
	Subject       string            `json:"subject,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
	Substitutions map[string]string `json:"substitutions,omitempty"`
	CustomArgs    map[string]string `json:"custom_args,omitempty"`
	Categories    []string          `json:"categories,omitempty"`
	SendAt        int               `json:"send_at,omitempty"`
}

// Email
type Email struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"email,omitempty"`
}

// Content
type Content struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// Attachment
type Attachment struct {
	Content     string `json:"content,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Disposition string `json:"disposition,omitempty"`
	ContentID   string `json:"content_id,omitempty"`
}

// Asm
type Asm struct {
	GroupID         int   `json:"group_id,omitempty"`
	GroupsToDisplay []int `json:"groups_to_display,omitempty"`
}

// MailSettings
type MailSettings struct {
	BCC                  *BccSetting       `json:"bcc,omitempty"`
	BypassListManagement *Setting          `json:"bypass_list_management,omitempty"`
	Footer               *FooterSetting    `json:"footer,omitempty"`
	SandboxMode          *Setting          `json:"sandbox_mode,omitempty"`
	SpamCheckSetting     *SpamCheckSetting `json:"spam_check,omitempty"`
}

// TrackingSettings
type TrackingSettings struct {
	ClickTracking        *ClickTrackingSetting        `json:"click_tracking,omitempty"`
	OpenTracking         *OpenTrackingSetting         `json:"open_tracking,omitempty"`
	SubscriptionTracking *SubscriptionTrackingSetting `json:"subscription_tracking,omitempty"`
	GoogleAnalytics      *GaSetting                   `json:"ganalytics,omitempty"`
	BCC                  *BccSetting                  `json:"bcc,omitempty"`
	BypassListManagement *Setting                     `json:"bypass_list_management,omitempty"`
	Footer               *FooterSetting               `json:"footer,omitempty"`
	SandboxMode          *SandboxModeSetting          `json:"sandbox_mode,omitempty"`
}

// BccSetting
type BccSetting struct {
	Enable *bool  `json:"enable,omitempty"`
	Email  string `json:"email,omitempty"`
}

// FooterSetting
type FooterSetting struct {
	Enable *bool  `json:"enable,omitempty"`
	Text   string `json:"text,omitempty"`
	Html   string `json:"html,omitempty"`
}

// ClickTrackingSetting
type ClickTrackingSetting struct {
	Enable     *bool `json:"enable,omitempty"`
	EnableText *bool `json:"enable_text,omitempty"`
}

// OpenTrackingSetting
type OpenTrackingSetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	SubstitutionTag string `json:"substitution_tag,omitempty"`
}

// SandboxModeSetting
type SandboxModeSetting struct {
	Enable      *bool             `json:"enable,omitempty"`
	ForwardSpam *bool             `json:"forward_spam,omitempty"`
	SpamCheck   *SpamCheckSetting `json:"spam_check,omitempty"`
}

// SpamCheckSetting
type SpamCheckSetting struct {
	Enable        *bool  `json:"enable,omitempty"`
	SpamThreshold int    `json:"threshold,omitempty"`
	PostToURL     string `json:"post_to_url,omitempty"`
}

// SubscriptionTrackingSetting
type SubscriptionTrackingSetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	Text            string `json:"text,omitempty"`
	Html            string `json:"html,omitempty"`
	SubstitutionTag string `json:"substitution_tag,omitempty"`
}

// GaSetting
type GaSetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	CampaignSource  string `json:"utm_source,omitempty"`
	CampaignTerm    string `json:"utm_term,omitempty"`
	CampaignContent string `json:"utm_content,omitempty"`
	CampaignName    string `json:"utm_campaign,omitempty"`
	CampaignMedium  string `json:"utm_medium,omitempty"`
}

// Setting
type Setting struct {
	Enable *bool `json:"enable,omitempty"`
}

// NewV3Mail
func NewV3Mail() *SGMailV3 {
	return &SGMailV3{
		Personalizations: make([]*Personalization, 0),
		Content:          make([]*Content, 0),
		Attachments:      make([]*Attachment, 0),
	}
}

// NewV3MailInit
func NewV3MailInit(from *Email, subject string, to *Email, content ...*Content) *SGMailV3 {
	m := new(SGMailV3)
	m.SetFrom(from)
	m.Subject = subject
	p := NewPersonalization()
	p.AddTos(to)
	m.AddPersonalizations(p)
	m.AddContent(content...)
	return m
}

// GetRequestBody
func GetRequestBody(m *SGMailV3) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	return b
}

// SGMailV3.AddPersonalizations
func (s *SGMailV3) AddPersonalizations(p ...*Personalization) *SGMailV3 {
	s.Personalizations = append(s.Personalizations, p...)
	return s
}

// SGMailV3.AddContent
func (s *SGMailV3) AddContent(c ...*Content) *SGMailV3 {
	s.Content = append(s.Content, c...)
	return s
}

// SGMailV3.AddAttachment
func (s *SGMailV3) AddAttachment(a ...*Attachment) *SGMailV3 {
	s.Attachments = append(s.Attachments, a...)
	return s
}

// SGMailV3.SetFrom
func (s *SGMailV3) SetFrom(e *Email) *SGMailV3 {
	s.From = e
	return s
}

// SGMailV3.SetReplyTo
func (s *SGMailV3) SetReplyTo(e *Email) *SGMailV3 {
	s.ReplyTo = e
	return s
}

// SGMailV3.SetTemplateID
func (s *SGMailV3) SetTemplateID(templateID string) *SGMailV3 {
	s.TemplateID = templateID
	return s
}

// SGMailV3.AddSection
func (s *SGMailV3) AddSection(key string, value string) *SGMailV3 {
	if s.Sections == nil {
		s.Sections = make(map[string]string)
	}

	s.Sections[key] = value
	return s
}

// SGMailV3.SetHeader
func (s *SGMailV3) SetHeader(key string, value string) *SGMailV3 {
	if s.Headers == nil {
		s.Headers = make(map[string]string)
	}

	s.Headers[key] = value
	return s
}

// SGMailV3.AddCategories
func (s *SGMailV3) AddCategories(category ...string) *SGMailV3 {
	s.Categories = append(s.Categories, category...)
	return s
}

// SGMailV3.SetCustomArg
func (s *SGMailV3) SetCustomArg(key string, value string) *SGMailV3 {
	if s.CustomArgs == nil {
		s.CustomArgs = make(map[string]string)
	}

	s.CustomArgs[key] = value
	return s
}

// SGMailV3.SetSendAt
func (s *SGMailV3) SetSendAt(sendAt int) *SGMailV3 {
	s.SendAt = sendAt
	return s
}

// SGMailV3.SetBatchID
func (s *SGMailV3) SetBatchID(batchID string) *SGMailV3 {
	s.BatchID = batchID
	return s
}

// SGMailV3.SetASM
func (s *SGMailV3) SetASM(asm *Asm) *SGMailV3 {
	s.Asm = asm
	return s
}

// SGMailV3.SetIPPoolID
func (s *SGMailV3) SetIPPoolID(ipPoolID string) *SGMailV3 {
	s.IPPoolID = ipPoolID
	return s
}

// SGMailV3.SetMailSettings
func (s *SGMailV3) SetMailSettings(mailSettings *MailSettings) *SGMailV3 {
	s.MailSettings = mailSettings
	return s
}

// SGMailV3.SetTrackingSettings
func (s *SGMailV3) SetTrackingSettings(trackingSettings *TrackingSettings) *SGMailV3 {
	s.TrackingSettings = trackingSettings
	return s
}

// NewPersonalization
func NewPersonalization() *Personalization {
	return &Personalization{
		To:            make([]*Email, 0),
		CC:            make([]*Email, 0),
		BCC:           make([]*Email, 0),
		Headers:       make(map[string]string),
		Substitutions: make(map[string]string),
		CustomArgs:    make(map[string]string),
		Categories:    make([]string, 0),
	}
}

// Personalization.AddTos
func (p *Personalization) AddTos(to ...*Email) {
	p.To = append(p.To, to...)
}

// Personalization.AddCCs
func (p *Personalization) AddCCs(cc ...*Email) {
	p.CC = append(p.CC, cc...)
}

// Personalization.AddBCCs
func (p *Personalization) AddBCCs(bcc ...*Email) {
	p.BCC = append(p.BCC, bcc...)
}

// Personalization.SetHeader
func (p *Personalization) SetHeader(key string, value string) {
	p.Headers[key] = value
}

// Personalization.SetSubstitution
func (p *Personalization) SetSubstitution(key string, value string) {
	p.Substitutions[key] = value
}

// Personalization.SetCustomArg
func (p *Personalization) SetCustomArg(key string, value string) {
	p.CustomArgs[key] = value
}

// Personalization.SetSendAt
func (p *Personalization) SetSendAt(sendAt int) {
	p.SendAt = sendAt
}

// NewAttachment
func NewAttachment() *Attachment {
	return &Attachment{}
}

// Attachment.SetContent
func (a *Attachment) SetContent(content string) *Attachment {
	a.Content = content
	return a
}

// Attachment.SetType
func (a *Attachment) SetType(contentType string) *Attachment {
	a.Type = contentType
	return a
}

// Attachment.SetFilename
func (a *Attachment) SetFilename(filename string) *Attachment {
	a.Filename = filename
	return a
}

// Attachment.SetDisposition
func (a *Attachment) SetDisposition(disposition string) *Attachment {
	a.Disposition = disposition
	return a
}

// Attachment.SetContentID
func (a *Attachment) SetContentID(contentID string) *Attachment {
	a.ContentID = contentID
	return a
}

// NewASM
func NewASM() *Asm {
	return &Asm{}
}

// Asm.SetGroupID
func (a *Asm) SetGroupID(groupID int) *Asm {
	a.GroupID = groupID
	return a
}

// Asm.AddGroupsToDisplay
func (a *Asm) AddGroupsToDisplay(groupsToDisplay ...int) *Asm {
	a.GroupsToDisplay = append(a.GroupsToDisplay, groupsToDisplay...)
	return a
}

// NewMailSettings
func NewMailSettings() *MailSettings {
	return &MailSettings{}
}

// MailSettings.SetBCC
func (m *MailSettings) SetBCC(bcc *BccSetting) *MailSettings {
	m.BCC = bcc
	return m
}

// MailSettings.SetBypassListManagement
func (m *MailSettings) SetBypassListManagement(bypassListManagement *Setting) *MailSettings {
	m.BypassListManagement = bypassListManagement
	return m
}

// MailSettings.SetFooter
func (m *MailSettings) SetFooter(footerSetting *FooterSetting) *MailSettings {
	m.Footer = footerSetting
	return m
}

// MailSettings.SetSandboxMode
func (m *MailSettings) SetSandboxMode(sandboxMode *Setting) *MailSettings {
	m.SandboxMode = sandboxMode
	return m
}

// MailSettings.SetSpamCheckSettings
func (m *MailSettings) SetSpamCheckSettings(spamCheckSetting *SpamCheckSetting) *MailSettings {
	m.SpamCheckSetting = spamCheckSetting
	return m
}

// NewTrackingSettings
func NewTrackingSettings() *TrackingSettings {
	return &TrackingSettings{}
}

// TrackingSettings.SetClickTracking
func (t *TrackingSettings) SetClickTracking(clickTracking *ClickTrackingSetting) *TrackingSettings {
	t.ClickTracking = clickTracking
	return t

}

// TrackingSettings.SetOpenTracking
func (t *TrackingSettings) SetOpenTracking(openTracking *OpenTrackingSetting) *TrackingSettings {
	t.OpenTracking = openTracking
	return t
}

// TrackingSettings.SetSubscriptionTracking
func (t *TrackingSettings) SetSubscriptionTracking(subscriptionTracking *SubscriptionTrackingSetting) *TrackingSettings {
	t.SubscriptionTracking = subscriptionTracking
	return t
}

// TrackingSettings.SetGoogleAnalytics
func (t *TrackingSettings) SetGoogleAnalytics(googleAnalytics *GaSetting) *TrackingSettings {
	t.GoogleAnalytics = googleAnalytics
	return t
}

// NewBCCSetting
func NewBCCSetting() *BccSetting {
	return &BccSetting{}
}

// BccSetting.SetEnable
func (b *BccSetting) SetEnable(enable bool) *BccSetting {
	setEnable := enable
	b.Enable = &setEnable
	return b
}

// BccSetting.SetEmail
func (b *BccSetting) SetEmail(email string) *BccSetting {
	b.Email = email
	return b
}

// NewFooterSetting
func NewFooterSetting() *FooterSetting {
	return &FooterSetting{}
}

// FooterSetting.SetEnable
func (f *FooterSetting) SetEnable(enable bool) *FooterSetting {
	setEnable := enable
	f.Enable = &setEnable
	return f
}

// FooterSetting.SetText
func (f *FooterSetting) SetText(text string) *FooterSetting {
	f.Text = text
	return f
}

// FooterSetting.SetHTML
func (f *FooterSetting) SetHTML(html string) *FooterSetting {
	f.Html = html
	return f
}

// NewOpenTrackingSetting
func NewOpenTrackingSetting() *OpenTrackingSetting {
	return &OpenTrackingSetting{}
}

// OpenTrackingSetting.SetEnable
func (o *OpenTrackingSetting) SetEnable(enable bool) *OpenTrackingSetting {
	setEnable := enable
	o.Enable = &setEnable
	return o
}

// OpenTrackingSetting.SetSubstitutionTag
func (o *OpenTrackingSetting) SetSubstitutionTag(subTag string) *OpenTrackingSetting {
	o.SubstitutionTag = subTag
	return o
}

// NewSubscriptionTrackingSetting
func NewSubscriptionTrackingSetting() *SubscriptionTrackingSetting {
	return &SubscriptionTrackingSetting{}
}

// SubscriptionTrackingSetting.SetEnable
func (s *SubscriptionTrackingSetting) SetEnable(enable bool) *SubscriptionTrackingSetting {
	setEnable := enable
	s.Enable = &setEnable
	return s
}

// SubscriptionTrackingSetting.SetText
func (s *SubscriptionTrackingSetting) SetText(text string) *SubscriptionTrackingSetting {
	s.Text = text
	return s
}

// SubscriptionTrackingSetting.SetHTML
func (s *SubscriptionTrackingSetting) SetHTML(html string) *SubscriptionTrackingSetting {
	s.Html = html
	return s
}

// SubscriptionTrackingSetting.SetSubstitutionTag
func (s *SubscriptionTrackingSetting) SetSubstitutionTag(subTag string) *SubscriptionTrackingSetting {
	s.SubstitutionTag = subTag
	return s
}

// NewGaSetting
func NewGaSetting() *GaSetting {
	return &GaSetting{}
}

// GaSetting.SetEnable
func (g *GaSetting) SetEnable(enable bool) *GaSetting {
	setEnable := enable
	g.Enable = &setEnable
	return g
}

// GaSetting.SetCampaignSource
func (g *GaSetting) SetCampaignSource(campaignSource string) *GaSetting {
	g.CampaignSource = campaignSource
	return g
}

// GaSetting.SetCampaignContent
func (g *GaSetting) SetCampaignContent(campaignContent string) *GaSetting {
	g.CampaignContent = campaignContent
	return g
}

// GaSetting.SetCampaignTerm
func (g *GaSetting) SetCampaignTerm(campaignTerm string) *GaSetting {
	g.CampaignTerm = campaignTerm
	return g
}

// GaSetting.SetCampaignName
func (g *GaSetting) SetCampaignName(campaignName string) *GaSetting {
	g.CampaignName = campaignName
	return g
}

// GaSetting.SetCampaignMedium
func (g *GaSetting) SetCampaignMedium(campaignMedium string) *GaSetting {
	g.CampaignMedium = campaignMedium
	return g
}

// NewSetting
func NewSetting(enable bool) *Setting {
	setEnable := enable
	return &Setting{Enable: &setEnable}
}

// NewEmail
func NewEmail(name string, address string) *Email {
	return &Email{
		Name:    name,
		Address: address,
	}
}

// NewSingleEmail
func NewSingleEmail(from *Email, subject string, to *Email, plainTextContent string, htmlContent string) *SGMailV3 {
	plainText := NewContent("text/plain", plainTextContent)
	html := NewContent("text/html", htmlContent)
	return NewV3MailInit(from, subject, to, plainText, html)
}

// NewContent
func NewContent(contentType string, value string) *Content {
	return &Content{
		Type:  contentType,
		Value: value,
	}
}

// NewClickTrackingSetting
func NewClickTrackingSetting() *ClickTrackingSetting {
	return &ClickTrackingSetting{}
}

// ClickTrackingSetting.SetEnable
func (c *ClickTrackingSetting) SetEnable(enable bool) *ClickTrackingSetting {
	setEnable := enable
	c.Enable = &setEnable
	return c
}

// ClickTrackingSetting.SetEnableText
func (c *ClickTrackingSetting) SetEnableText(enableText bool) *ClickTrackingSetting {
	setEnable := enableText
	c.EnableText = &setEnable
	return c
}

// NewSpamCheckSetting
func NewSpamCheckSetting() *SpamCheckSetting {
	return &SpamCheckSetting{}
}

// SpamCheckSetting.SetEnable
func (s *SpamCheckSetting) SetEnable(enable bool) *SpamCheckSetting {
	setEnable := enable
	s.Enable = &setEnable
	return s
}

// SpamCheckSetting.SetSpamThreshold
func (s *SpamCheckSetting) SetSpamThreshold(spamThreshold int) *SpamCheckSetting {
	s.SpamThreshold = spamThreshold
	return s
}

// SpamCheckSetting.SetPostToURL
func (s *SpamCheckSetting) SetPostToURL(postToURL string) *SpamCheckSetting {
	s.PostToURL = postToURL
	return s
}

// NewSandboxModeSetting
func NewSandboxModeSetting(enable bool, forwardSpam bool, spamCheck *SpamCheckSetting) *SandboxModeSetting {
	setEnable := enable
	setForwardSpam := forwardSpam
	return &SandboxModeSetting{
		Enable:      &setEnable,
		ForwardSpam: &setForwardSpam,
		SpamCheck:   spamCheck,
	}
}
