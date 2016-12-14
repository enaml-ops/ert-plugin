package deploy_notifications_ui 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type NotificationsUi struct {

	/*AppDomain - Descr: Domain used to host application Default: <nil>
*/
	AppDomain interface{} `yaml:"app_domain,omitempty"`

	/*Network - Descr: Network used to host application Default: <nil>
*/
	Network interface{} `yaml:"network,omitempty"`

	/*Space - Descr: Space to deploy to Default: <nil>
*/
	Space interface{} `yaml:"space,omitempty"`

	/*EncryptionKey - Descr: Key used to encrypt secret values Default: <nil>
*/
	EncryptionKey interface{} `yaml:"encryption_key,omitempty"`

	/*SyslogUrl - Descr: Syslog URL to bind the app logs to Default: <nil>
*/
	SyslogUrl interface{} `yaml:"syslog_url,omitempty"`

	/*FooterText - Descr: Footer text to display in the UI Default: <nil>
*/
	FooterText interface{} `yaml:"footer_text,omitempty"`

	/*FooterImage - Descr: Footer image source Default: <nil>
*/
	FooterImage interface{} `yaml:"footer_image,omitempty"`

	/*Uaa - Descr: Secret of the notifications-ui UAA client Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Cf - Descr: Username of the CF admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*EnableDiego - Descr: Enable diego deployment Default: <nil>
*/
	EnableDiego interface{} `yaml:"enable_diego,omitempty"`

	/*InstanceCount - Descr: Number of instances to run Default: <nil>
*/
	InstanceCount interface{} `yaml:"instance_count,omitempty"`

	/*Organization - Descr: Organization to deploy to Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

	/*AppSubdomain - Descr: Subdomain of route serving the application, usually the app name, e.g. notifications-ui Default: <nil>
*/
	AppSubdomain interface{} `yaml:"app_subdomain,omitempty"`

}