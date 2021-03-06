package destroy_notifications_ui 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type NotificationsUi struct {

	/*Organization - Descr: Organization where the app is deployed Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

	/*Space - Descr: Space where the app is deployed Default: <nil>
*/
	Space interface{} `yaml:"space,omitempty"`

	/*Cf - Descr: CF CLI connection dial timeout Default: 5
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*AppDomain - Descr: Domain used to host application Default: <nil>
*/
	AppDomain interface{} `yaml:"app_domain,omitempty"`

}