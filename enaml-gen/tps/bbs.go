package tps 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Bbs struct {

	/*ClientKey - Descr: PEM-encoded client key Default: <nil>
*/
	ClientKey interface{} `yaml:"client_key,omitempty"`

	/*ApiLocation - Descr: Address to the BBS Server Default: bbs.service.cf.internal:8889
*/
	ApiLocation interface{} `yaml:"api_location,omitempty"`

	/*ClientSessionCacheSize - Descr: capacity of the tls client cache Default: <nil>
*/
	ClientSessionCacheSize interface{} `yaml:"client_session_cache_size,omitempty"`

	/*RequireSsl - Descr: enable ssl for all communication with the bbs Default: true
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

	/*CaCert - Descr: PEM-encoded CA certificate Default: <nil>
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*ClientCert - Descr: PEM-encoded client certificate Default: <nil>
*/
	ClientCert interface{} `yaml:"client_cert,omitempty"`

	/*MaxIdleConnsPerHost - Descr: maximum number of idle http connections Default: <nil>
*/
	MaxIdleConnsPerHost interface{} `yaml:"max_idle_conns_per_host,omitempty"`

}