package deploy_notifications_ui 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DeployNotificationsUiJob struct {

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Domain - Descr: Cloud Foundry System Domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*NotificationsUi - Descr: Organization to deploy to Default: <nil>
*/
	NotificationsUi *NotificationsUi `yaml:"notifications_ui,omitempty"`

}