package permissions

import (
	"strings"
)

const (
	Alerts             = "alerts"
	ApiKeys            = "api_keys"
	AsmGroups          = "asm.groups"
	Billing            = "billing"
	Categories         = "categories"
	Credentials        = "credentials"
	Ips                = "ips"
	Mail               = "mail"
	MarketingCampaigns = "marketing_campaigns"
	MailSettings       = "mail_settings"
	Stats              = "stats"
	Suppression        = "suppression"
	Teammates          = "teammates"
	Templates          = "templates"
	TrackingSettings   = "tracking_settings"
	PartnerSettings    = "partner_settings"
	ScheduledSends     = "user.scheduled_sends"
	Subusers           = "subusers"
	UserSettings       = "userSettings"
	Webhooks           = "webhooks"
	Whitelabel         = "whitelabel"
	Whitelist          = "whitelist"
)

const (
	categoriesStats              = Categories + ".stats"
	categoriesStatsSums          = Categories + ".stats.sums"
	emailActivity                = "email_activity"
	statsGlobal                  = Stats + ".global"
	browsersStats                = "browsers.stats"
	devicesStats                 = "devices.stats"
	geoStats                     = "geo.stats"
	mailboxProviderStats         = "mailbox_providers.stats"
	clients                      = "clients"
	clientsDesktopStats          = clients + ".desktop.stats"
	clientsPhoneStats            = clients + ".phone.stats"
	clientsStats                 = clients + ".stats"
	clientsTabletStats           = clients + ".tablet.stats"
	clientsWebmailStats          = clients + ".webmail.stats"
	ipsAssigned                  = Ips + ".assigned"
	ipsPools                     = Ips + ".pools"
	ipsPoolsIps                  = Ips + ".pools.ips"
	ipsWarmup                    = Ips + ".warmup"
	mailBatch                    = Mail + ".batch"
	mailSettingsAddressWhitelist = MailSettings + ".address_whitelist"
	mailSettingsBcc              = MailSettings + ".bcc"
	mailSettingsForwardBounce    = MailSettings + ".forward_bounce"
	mailSettingsBouncePurge      = MailSettings + ".bounce_purge"
	mailSettingsFooter           = MailSettings + ".footer"
	mailSettingsForwardSpam      = MailSettings + ".forward_spam"
	mailSettingsPlainContent     = MailSettings + ".plain_content"
	mailSettingsSpamCheck        = MailSettings + ".spam_check"
	mailSettingsTemplate         = MailSettings + ".template"
	partnerSettingsNewRelic      = PartnerSettings + ".new_relic"
	partnerSettingsSendWithUs    = PartnerSettings + ".sendwithus"
	subusersCredits              = Subusers + ".credits"
	subusersCreditsRemaining     = subusersCredits + ".remaining"
	subusersMonitors             = Subusers + ".monitor"
	subusersReputations          = Subusers + ".reputations"
	subusersStats                = Subusers + ".stats"
	subusersStatsMonthly         = subusersStats + ".monthly"
	subusersStatsSums            = subusersStats + ".sums"
	subusersSummary              = Subusers + ".summary"
	suppressionBounces           = Suppression + ".bounces"
	suppressionBlocks            = Suppression + ".blocks"
	suppressionInvalidEmails     = Suppression + ".invalid_emails"
	suppressionSpamReports       = Suppression + ".spam_reports"
	suppressionUnsubscribes      = Suppression + ".unsubscribes"
	templatesVersions            = Templates + ".versions"
	templatesVersionsActivate    = templatesVersions + ".activate"
	trackingSettingsClick        = TrackingSettings + ".click"
	trackingSettingsGoogle       = TrackingSettings + ".google_analytics"
	trackingSettingsOpen         = TrackingSettings + ".open"
	trackingSettingsSubscription = TrackingSettings + ".subscription"
	user                         = "user"
	userAccount                  = user + ".account"
	userCredits                  = user + ".credits"
	userEmail                    = user + ".email"
	userMultifactorAuth          = user + ".multifactor_authentication"
	userPassword                 = user + ".password"
	userProfile                  = user + ".profile"
	userSettingsEnforcedTLS      = user + ".settings.enforced_tls"
	userTimezone                 = user + ".timezone"
	userUsername                 = user + ".username"
	userWebhooks                 = user + ".webhooks"
	userWebhooksEvent            = userWebhooks + ".event"
	userWebhooksEventSettings    = userWebhooksEvent + ".settings"
	userWebhooksEventTest        = userWebhooksEvent + ".test"
	userWebhooksParse            = userWebhooks + ".parse"
	userWebhooksParseSettings    = userWebhooksParse + ".settings"
	userWebhooksParseStats       = userWebhooksParse + ".stats"
	accessSettings               = "access_settings"
	accessSettingsActivity       = accessSettings + ".activity"
	accessSettingsWhitelist      = accessSettings + ".whitelist"
	read                         = "read"
	update                       = "update"
	create                       = "create"
	delete                       = "delete"
	send                         = "send"
	delimiter                    = "."
)

var accessorMaps = map[string]permissionsAccessor{
	Alerts:             alertsPermissions(),
	ApiKeys:            apiKeysPermissions(),
	AsmGroups:          asmGroupsPermissions(),
	Billing:            billingPermissions(),
	Categories:         categoriesPermissions(),
	Credentials:        credentialsPermissions(),
	Ips:                iPsPermissions(),
	Mail:               mailPermissions(),
	MailSettings:       mailSettingsPermissions(),
	MarketingCampaigns: marketingCampaignsPermissions(),
	PartnerSettings:    partnerSettingsPermissions(),
	ScheduledSends:     scheduledSendsPermissions(),
	Stats:              statsPermissions(),
	Subusers:           subusersPermissions(),
	Suppression:        suppressionsPermissions(),
	Teammates:          teammatesPermissions(),
	Templates:          templatesPermissions(),
	TrackingSettings:   trackingPermissions(),
	UserSettings:       userSettingsPermissions(),
	Webhooks:           webhookPermissions(),
	Whitelabel:         whitelabelPermissions(),
	Whitelist:          whitelistPermissions(),
}

func permissionScopes(permission string, scopes ...string) []string {
	var permissionScopes []string
	for _, scope := range scopes {
		permissionScopes = append(permissionScopes, strings.Join([]string{permission, scope}, delimiter))
	}
	return permissionScopes
}

func merge(slices ...[]string) []string {
	var m []string
	for _, slice := range slices {
		m = append(m, slice...)
	}
	return m
}

type permissionsAccessor func() []string

func alertsPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(Alerts, create, read, update, delete)
	}
}
func apiKeysPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(ApiKeys, create, read, update, delete)
	}
}

func asmGroupsPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(AsmGroups, create, read, update, delete)
	}
}

func billingPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(Billing, create, read, update, delete)
	}
}

func categoriesPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(Categories, create, read, update, delete),
			permissionScopes(categoriesStats, read),
			permissionScopes(categoriesStatsSums, read),
		)
	}
}

func credentialsPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(Credentials, create, read, update, delete)
	}
}

func statsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(emailActivity, read),
			permissionScopes(Stats, read),
			permissionScopes(statsGlobal, read),
			permissionScopes(browsersStats, read),
			permissionScopes(devicesStats, read),
			permissionScopes(geoStats, read),
			permissionScopes(mailboxProviderStats, read),
			permissionScopes(clientsDesktopStats, read),
			permissionScopes(clientsPhoneStats, read),
			permissionScopes(clientsStats, read),
			permissionScopes(clientsTabletStats, read),
			permissionScopes(clientsWebmailStats, read),
		)
	}
}

func iPsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(ipsAssigned, read),
			permissionScopes(Ips, read),
			permissionScopes(ipsPools, create, read, update, delete),
			permissionScopes(ipsPoolsIps, create, read, update, delete),
			permissionScopes(ipsWarmup, create, read, update, delete),
		)
	}
}

func mailPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(mailBatch, create, read, update, delete),
			permissionScopes(Mail, send),
		)
	}
}

func mailSettingsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(mailSettingsAddressWhitelist, read, update),
			permissionScopes(mailSettingsBcc, read, update),
			permissionScopes(mailSettingsBouncePurge, read, update),
			permissionScopes(mailSettingsFooter, read, update),
			permissionScopes(mailSettingsForwardBounce, read, update),
			permissionScopes(mailSettingsForwardSpam, read, update),
			permissionScopes(mailSettingsPlainContent, read, update),
			permissionScopes(MailSettings, read),
			permissionScopes(mailSettingsSpamCheck, read, update),
			permissionScopes(mailSettingsTemplate, read, update),
		)
	}
}

func marketingCampaignsPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(MarketingCampaigns, create, read, update, delete)
	}
}

func partnerSettingsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(partnerSettingsNewRelic, read, update),
			permissionScopes(partnerSettingsSendWithUs, read, update),
			permissionScopes(PartnerSettings, read),
		)
	}
}

func scheduledSendsPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(ScheduledSends, create, read, update, delete)
	}
}

func subusersPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(Subusers, create, read, update, delete),
			permissionScopes(subusersCredits, create, read, update, delete),
			permissionScopes(subusersCreditsRemaining, create, read, update, delete),
			permissionScopes(subusersMonitors, create, read, update, delete),
			permissionScopes(subusersReputations, read),
			permissionScopes(subusersStats, read),
			permissionScopes(subusersStatsMonthly, read),
			permissionScopes(subusersStatsSums, read),
			permissionScopes(subusersSummary, read),
		)
	}
}

func suppressionsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(Suppression, create, read, update, delete),
			permissionScopes(suppressionBounces, create, read, update, delete),
			permissionScopes(suppressionBlocks, create, read, update, delete),
			permissionScopes(suppressionInvalidEmails, create, read, update, delete),
			permissionScopes(suppressionSpamReports, create, read, update, delete),
			permissionScopes(suppressionUnsubscribes, create, read, update, delete),
		)
	}
}

func teammatesPermissions() permissionsAccessor {
	return func() []string {
		return permissionScopes(Teammates, create, read, update, delete)
	}
}
func templatesPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(Templates, create, read, update, delete),
			permissionScopes(templatesVersions, create, read, update, delete),
			permissionScopes(templatesVersionsActivate, create, read, update, delete),
		)
	}
}

func trackingPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(TrackingSettings, read),
			permissionScopes(trackingSettingsClick, read, update),
			permissionScopes(trackingSettingsGoogle, read, update),
			permissionScopes(trackingSettingsOpen, read, update),
			permissionScopes(trackingSettingsSubscription, read, update),
		)
	}
}

func userSettingsPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(userAccount, read),
			permissionScopes(userCredits, read),
			permissionScopes(userEmail, create, read, update, delete),
			permissionScopes(userMultifactorAuth, create, read, update, delete),
			permissionScopes(userPassword, read, update),
			permissionScopes(userProfile, read, update),
			permissionScopes(userSettingsEnforcedTLS, read, update),
			permissionScopes(userTimezone, read, update),
			permissionScopes(userUsername, read, update),
		)
	}
}

func webhookPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(userWebhooksEventSettings, read, update),
			permissionScopes(userWebhooksEventTest, create, read, update),
			permissionScopes(userWebhooksParseSettings, create, read, update, delete),
			permissionScopes(userWebhooksParseStats, read),
		)
	}
}

func whitelabelPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(Whitelabel, create, read, update, delete),
		)
	}
}

func whitelistPermissions() permissionsAccessor {
	return func() []string {
		return merge(
			permissionScopes(accessSettingsWhitelist, create, read, update, delete),
			permissionScopes(accessSettingsActivity, read),
		)
	}
}

func GetPermissions(permissionClass string, fullAccess bool) []string {
	permissionScopes := accessorMaps[permissionClass]()
	if fullAccess {
		return permissionScopes
	}
	var readOnlyPerms []string
	for _, permission := range permissionScopes {
		a := strings.Split(permission, delimiter)
		if a[len(a)-1] == read {
			readOnlyPerms = append(readOnlyPerms, permission)
		}
	}
	return readOnlyPerms
}

func GetAdminPermissions() []string {
	var permissions []string
	for _, accessor := range accessorMaps {
		permissions = append(permissions, accessor()...)
	}
	return permissions
}
