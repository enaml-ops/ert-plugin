package push_apps_manager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ssl struct {

	/*DisableHttpForLinks - Descr: When displaying external app route urls, prefix all of them with https Default: false
*/
	DisableHttpForLinks interface{} `yaml:"disable_http_for_links,omitempty"`

	/*SkipCertVerify - Descr: when connecting over TLS, don't verify certificates Default: false
*/
	SkipCertVerify interface{} `yaml:"skip_cert_verify,omitempty"`

}