package acceptance_tests 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cf struct {

	/*AppDomains - Descr: List of shared domains for pushed apps (tests will use the first domain) Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*AdminPassword - Descr: Password of the admin user Default: <nil>
*/
	AdminPassword interface{} `yaml:"admin_password,omitempty"`

	/*SmokeTests - Descr: Runs smoke test errand as an existing org. Creates a new org if false Default: false
*/
	SmokeTests *SmokeTests `yaml:"smoke_tests,omitempty"`

	/*ApiUrl - Descr: Full URL of Cloud Foundry API Default: <nil>
*/
	ApiUrl interface{} `yaml:"api_url,omitempty"`

	/*AdminUsername - Descr: Username of the admin user Default: <nil>
*/
	AdminUsername interface{} `yaml:"admin_username,omitempty"`

	/*SkipSslValidation - Descr: Whether to add --skip-ssl-validation for cf cli Default: false
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

}