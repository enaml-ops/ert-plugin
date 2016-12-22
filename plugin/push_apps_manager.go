package cloudfoundry

import (
	"fmt"

	"github.com/enaml-ops/enaml"
	pam "github.com/enaml-ops/ert-plugin/enaml-gen/push-apps-manager"
	"github.com/enaml-ops/ert-plugin/plugin/config"
)

type pushAppsManager struct {
	Config *config.Config
}

func NewPushAppsManager(c *config.Config) InstanceGroupCreator {
	return &pushAppsManager{Config: c}
}

func (p *pushAppsManager) ToInstanceGroup() *enaml.InstanceGroup {

	var appsDomain = ""
	if len(p.Config.AppDomains) > 0 {
		appsDomain = p.Config.AppDomains[0]
	}

	return &enaml.InstanceGroup{
		Name:      "push-apps-manager",
		Instances: 1,
		VMType:    p.Config.ErrandVMType,
		Lifecycle: "errand",
		AZs:       p.Config.AZs,
		Stemcell:  p.Config.StemcellName,
		Networks: []enaml.Network{
			{Name: p.Config.NetworkName},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
		Jobs: []enaml.InstanceJob{
			{
				Name:    "push-apps-manager",
				Release: PushAppsReleaseName,
				Properties: &pam.PushAppsManagerJob{
					Cf: &pam.Cf{
						AppsDomain:    appsDomain,
						ApiUrl:        fmt.Sprintf("https://api.%s", p.Config.SystemDomain),
						AdminUsername: "push_apps_manager",
						AdminPassword: p.Config.PushAppsManagerPassword,
						SystemDomain:  p.Config.SystemDomain,
					},
					Services: &pam.Services{
						Authentication: &pam.Authentication{
							CFCLIENTID:       "portal",
							CFCLIENTSECRET:   p.Config.PortalClientSecret,
							CFUAASERVERURL:   fmt.Sprintf("https://uaa.%s", p.Config.SystemDomain),
							CFLOGINSERVERURL: fmt.Sprintf("https://login.%s", p.Config.SystemDomain),
						},
					},
					Env: &pam.Env{
						SecretToken:               p.Config.AppsManagerSecretToken,
						CfCcApiUrl:                fmt.Sprintf("https://api.%s", p.Config.SystemDomain),
						CfLoggregatorHttpUrl:      fmt.Sprintf("http://loggregator.%s", p.Config.SystemDomain),
						CfLoggregatorWebsocketUrl: fmt.Sprintf("wss://loggregator.%s:%v", p.Config.SystemDomain, p.Config.LoggregatorPort),
						CfNotificationsServiceUrl: fmt.Sprintf("https://notifications.%s", p.Config.SystemDomain),
						UsageServiceHost:          fmt.Sprintf("https://app-usage.%s", p.Config.SystemDomain),
						GenericWhiteLabelConfigJson: &pam.GenericWhiteLabelConfigJson{
							CompanyName: "Pivotal",
							FooterText:  "©2016 Pivotal Software, Inc. All Rights Reserved.",
							FooterLinks: []string{},
							Logo:        `data:image/png;base64,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`,
							SquareLogo:  `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGwAAABsCAYAAACPZlfNAAAAAXNSR0IArs4c6QAABYtJREFUeAHtnVtsFFUYx7/d3ruWotUKVIkNaCw02YgJGBRTMd4CokUejD4QH4gxQcIDeHnBmPjkhSghUYLGe3ywPtAHNCo0QgkWwi2tXG2V1kIpLXTbLt1tS9dzlmzSJssZhv32zDk7/2km2znn7Pd9+/vt2Z2dmW0D9Obat4gCiwiLBQQSLflSViAQeN6Can1fYiJBFPQ9BcsAQBiEWUbAsnIxwyDMMgKWlYsZBmGWEbCsXMwwCLOMgGXlYoZBmGUELCsXMwzCLCNgWbmYYRBmGQHLysUMgzDLCFhWLmYYhFlGwLJyMcMgzDIClpWLGQZhlhGwrFzMMAizjIBl5WKGQZhlBCwrV1xbb96y59V1VFJQmLawQNrWa43x8XEaHo1fW+Oj1H8lSqf6eulEbw+dvNhLvcNDinvb0WWksAdm3UWhwiJ2gt2RAWo80UY7jrdSU8cZGrt6lT1HtgMaKSxbD7qqfDq99tAjyTUSG6FP9v1BH+3dTUPxeLZSssf17U5HeXEJbXr8aerY+A6tf7iOxFeu2OFmI6BvhaVgVoRCtHl5PTW8/AoV5xekmo299b2wlJn6+WFqWrOWKkpDqSYjbyFskpZFs++hL1e9NKnFvF+t3OmQOwzdkcgUmnnBABXm5Ys1j8qKisVadFPvS8tramn1goX09eEDU+KbsmGlsMbjbbT6x++UDOVORGXoFppXOYMerLqbVsyrpcWzqykYdH5R+fjZlcnd/8sjV5Q5vOh0rt6LqhhyJsQ3uC+ID8ry89aHYtf90W1bKLzlffr19EnH6HIP8oXasOM4LwbkrLB0MP+6cJ6e+eoz+vTP5nTdU9peDC+Ysm3Khq+ESehy5r3e2ECHu7uUDuqq59Id4iXVtMV3wqSACSHt3V2/KF3I97qayjuVY7zo9KUwCfq3M6coNjamZD6zrFzZ70Wnb4XFxseoK3JZyXzWtGnKfi86fStMwu6LRpXMZ5RBmBKQ7k75XqZa8gLmPZ/Nq0hFkLnvttJSZUT5Oc60xbfC5CGs6lsrlD56hgaV/V50+lbYkuo5VFygPp3SMwxhXjwp0+bcsGRp2vZU48TEBB09153aNObWlzNMHo1/6r4apYTmsx10MTqsHONFp5VH6zMBtWbhYtq6YpVjiJ/ajjmO8WKAL4QFxamWZffPT1678dicex05D4jTKj8cO+Q4zosBOSXs7bonktci5ovjgPIUye3ieo3wzKrk+TC5faPLGz83On6ovtFY3ONySth7Ty67qbPMk6Hu+edv+vzg/slNRv3uy52O6xk40HWW6r/94nrdRrTn1AzLhOju9tP03DfbKTo6mkmYrN/X98L6xQHgTb/vpG0t+5LnybJOPMMEvhXWOXCJvj9yiD7Yu4sGRkYyxKjv7r4RJi+Na+05Rwf/66SG1qO0v/NffZQZM+WUsI07d1BC/MTE144GYzHxJYcYDYq1vb/f8WQlI9OshsopYZubm7IKy4Tg2K03wYKLGiDMBSwThkKYCRZc1ABhLmCZMBTCTLDgogYIcwHLhKEQZoIFFzVAmAtYJgyFMBMsuKgBwlzAMmEohJlgwUUNEOYClglDIcwECy5qgDAXsEwYCmEmWHBRA4S5gGXCUAgzwYKLGow84yyvuyhR/GW19kt9Lh5ibg01UtjS7VtzizLjo8FLIiNMHaEgTAdlxhwQxghTRygI00GZMQeEMcLUEQrCdFBmzAFhjDB1hIIwHZQZc0AYI0wdoSBMB2XGHBDGCFNHKAjTQZkxB4QxwtQRCsJ0UGbMAWGMMHWEgjAdlBlzQBgjTB2hIEwHZcYcEMYIU0coCNNBmTEHhDHC1BEKwnRQZswBYYwwdYSCMB2UGXNAGCNMHaEgTAdlxhziUu1Ei8M/+WFMh1CZEUi0/A+j7hNSB5Wo2wAAAABJRU5ErkJggg==`,
						},
					},
					Databases: &pam.Databases{
						AppUsageService: &pam.DatabasesAppUsageService{
							Name:     "app_usage_service",
							Ip:       p.Config.MySQLProxyHost(),
							Port:     3306,
							Username: "root",
							Password: p.Config.MySQLAdminPassword,
						},
					},
					Ssl: &pam.Ssl{
						SkipCertVerify: p.Config.SkipSSLCertVerify,
					},
				},
			},
		},
	}
}
