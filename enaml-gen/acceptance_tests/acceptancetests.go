package acceptance_tests 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AcceptanceTests struct {

	/*IncludeHttpRoutes - Descr: When true, additionally runs the HTTP Routes test for Routing API Default: false
*/
	IncludeHttpRoutes interface{} `yaml:"include_http_routes,omitempty"`

	/*Verbose - Descr: Whether to pass the -v flag to router acceptance tests Default: false
*/
	Verbose interface{} `yaml:"verbose,omitempty"`

	/*Addresses - Descr: A list of addresses which will be checked for TCP connectivity and features Default: [10.244.14.2]
*/
	Addresses interface{} `yaml:"addresses,omitempty"`

	/*Nodes - Descr: The number of parallel test executors to spawn. The larger the number the higher the stress on the system. Default: 4
*/
	Nodes interface{} `yaml:"nodes,omitempty"`

	/*SkipSslValidation - Descr: When true, does not verify TLS certificates for any API calls made during the test run Default: false
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

	/*CloudController - Descr: URL of the Cloud Controller API Default: <nil>
*/
	CloudController *CloudController `yaml:"cloud_controller,omitempty"`

	/*SystemDomain - Descr: Domain for system components, e.g. bosh-lite.com Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

}