package sendgrid

type SGMailV3 struct {
	From             string                 `json:"from"`
	Subject          string                 `json:"subject"`
	Personalization  []*Personalization     `json:"personalization,omitempty"`
	Content          []*Content             `json:"content"`
	Attachments      []*Attachment          `json:"attachments,omitempty"`
	TemplateID       string                 `json:"template_id,omitempty"`
	Sections         map[string]string      `json:"sections,omitempty"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
	Categories       []string               `json:"categories,omitempty"`
	CustomArgs       map[string]interface{} `json:"custom_args,omitempty"`
	SendAt           int                    `json:"send_at,omitempty"`
	BatchID          string                 `json:"batch_id,omitempty"`
	Asm              *Asm                   `json:"asm,omitempty"`
	IPPoolID         int                    `json:"ip_pool_id,omitempty"`
	MailSettings     *MailSettings          `json:"mail_settings,omitempty"`
	TrackingSettings *TrackingSettings      `json:"tracking_settings,omitempty"`
}

type Personalization struct {
	To            []string               `json:"to"`
	CC            []string               `json:"cc"`
	BCC           []string               `json:"bcc"`
	Subject       string                 `json:"subject"`
	Headers       map[string]interface{} `json:"headers"`
	Substitutions map[string]interface{} `json:"substitutions"`
	CustomArgs    map[string]interface{} `json:"custom_args"`
	Categories    []string               `json:"categories"`
	SendAt        int                    `json:"send_at"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Attachment struct {
	Content     string `json:"content"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	FileName    string `json:"filename"`
	Disposition string `json:"disposition"`
	ContentID   string `json:"content_id,omitempty"`
}

type Asm struct {
	GroupID         int   `json:"group_id"`
	GroupsToDisplay []int `json:"groups_to_display,omitempty"`
}

type MailSettings struct {
	BCC                  BccSetting    `json:"bcc"`
	BypassListManagement Setting       `json:"bypass_list_management"`
	Footer               FooterSetting `json:"footer"`
	SandboxMode          Setting       `json:"sandbox_mode"`
}

type TrackingSettings struct {
	ClickTracking        Setting                     `json:"click_tracking"`
	OpenTracking         OpenTrackingSetting         `json:"open_tracking"`
	SubscriptionTracking SubscriptionTrackingSetting `json:"subscription_tracking"`
	GoogleAnalytics      GaSetting                   `json:"ganalytics"`
}

type BccSetting struct {
	Enable bool   `json:"enable"`
	Email  string `json:"email"`
}

type FooterSetting struct {
	Enable bool   `json:"enable"`
	Text   string `json:"text"`
	Html   string `json:"html"`
}

type OpenTrackingSetting struct {
	Enable          bool   `json:"enable"`
	SubstitutionTag string `json:"substitution_tag"`
}

type SubscriptionTrackingSetting struct {
	Enable          bool   `json:"enable"`
	Text            string `json:"text"`
	Html            string `json:"html"`
	SubstitutionTag string `json:"substitution_tag"`
}

type GaSetting struct {
	Enable          bool   `json:"enable"`
	CampaignSource  string `json:"Campaign Source"`
	CampaignTerm    string `json:"Campaign Term"`
	CampaignContent string `json:"Campaign Content"`
	CampaignName    string `json:"Campaign Name"`
}
type Setting struct {
	Enable bool `json:"enable"`
}
