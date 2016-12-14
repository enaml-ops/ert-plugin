package push_apps_manager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Authentication struct {

	/*CFCLIENTID - Descr: UAA client id for Console application Default: <nil>
*/
	CFCLIENTID interface{} `yaml:"CF_CLIENT_ID,omitempty"`

	/*CFUAASERVERURL - Descr: Cloud Foundry UAA server URL Default: <nil>
*/
	CFUAASERVERURL interface{} `yaml:"CF_UAA_SERVER_URL,omitempty"`

	/*CFLOGINSERVERURL - Descr: Cloud Foundry login server URL Default: <nil>
*/
	CFLOGINSERVERURL interface{} `yaml:"CF_LOGIN_SERVER_URL,omitempty"`

	/*CFCLIENTSECRET - Descr: UAA client secret for Console application Default: <nil>
*/
	CFCLIENTSECRET interface{} `yaml:"CF_CLIENT_SECRET,omitempty"`

}