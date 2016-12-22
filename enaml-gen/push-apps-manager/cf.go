package push_apps_manager

/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
 */
type Cf struct {

	/*SystemDomain - Descr: Cloud Foundry system domain, used for the Console application's URL Default: <nil>
	 */
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*ApiUrl - Descr: Full URL of Cloud Foundry API Default: <nil>
	 */
	ApiUrl interface{} `yaml:"api_url,omitempty"`

	/*AdminPassword - Descr: Password of the admin user Default: <nil>
	 */
	AdminPassword interface{} `yaml:"admin_password,omitempty"`

	/*AdminUsername - Descr: Username of the admin user Default: <nil>
	 */
	AdminUsername interface{} `yaml:"admin_username,omitempty"`

	/*AppsDomain - Descr: Cloud Foundry apps domain, used for default app domains Default: <nil>
	 */
	AppsDomain interface{} `yaml:"apps_domain,omitempty"`
}
