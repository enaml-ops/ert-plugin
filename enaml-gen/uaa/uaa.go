package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaa struct {

	/*User - Descr: Contains a list of the default authorities/scopes assigned to a user Default: [openid scim.me cloud_controller.read cloud_controller.write cloud_controller_service_permissions.read password.write uaa.user approvals.me oauth.approvals notification_preferences.read notification_preferences.write profile roles user_attributes uaa.offline_token]
*/
	User *UaaUser `yaml:"user,omitempty"`

	/*RequireHttps - Descr: Request came in on a secure connection. Expect the load balancer/proxy to set the proper headers (x-forwarded-for, x-forwarded-proto) Default: true
*/
	RequireHttps interface{} `yaml:"require_https,omitempty"`

	/*Ssl - Descr: If this property Tomcat will listen to this port and expect https traffic. If null, tomcat will not listen to this port Default: 8443
*/
	Ssl *UaaSsl `yaml:"ssl,omitempty"`

	/*Clients - Descr: List of OAuth2 clients that the UAA will be bootstrapped with Default: <nil>
*/
	Clients interface{} `yaml:"clients,omitempty"`

	/*Proxy - Descr: Array of the router IPs acting as the first group of HTTP/TCP backends. These will be added to the proxy_ips_regex as exact matches. Default: []
*/
	Proxy *Proxy `yaml:"proxy,omitempty"`

	/*Jwt - Descr: The global refresh token validity for all zones if nothing is configured on the client Default: 2592000
*/
	Jwt *Jwt `yaml:"jwt,omitempty"`

	/*Ldap - Descr: Configures the UAA LDAP referral behavior. The following values are possible:
- follow -> Referrals are followed
- ignore -> Referrals are ignored and the partial result is returned
- throw  -> An error is thrown and the authentication is aborted
Reference: http://docs.oracle.com/javase/jndi/tutorial/ldap/referral/jndi.html
 Default: follow
*/
	Ldap *Ldap `yaml:"ldap,omitempty"`

	/*Scim - Descr: Enables the endpoint `/ids/Users` that allows consumers to translate user ids to name Default: true
*/
	Scim *Scim `yaml:"scim,omitempty"`

	/*CatalinaOpts - Descr: The options used to configure Tomcat Default: -Xmx768m -XX:MaxPermSize=256m
*/
	CatalinaOpts interface{} `yaml:"catalina_opts,omitempty"`

	/*SslCertificate - Descr: The server's ssl certificate. The default is a self-signed certificate and should always be replaced for production deployments Default: 
*/
	SslCertificate interface{} `yaml:"sslCertificate,omitempty"`

	/*ProxyIpsRegex - Descr: A pipe delimited set of regular expressions of IP addresses that are considered reverse proxies.
When a request from these IP addresses come in, the x-forwarded-for and x-forwarded-proto headers will be respected.
If the uaa.restricted_ips_regex is set, it will be appended to this list for backwards compatibility purposes.
 Default: 10\.\d{1,3}\.\d{1,3}\.\d{1,3}|192\.168\.\d{1,3}\.\d{1,3}|169\.254\.\d{1,3}\.\d{1,3}|127\.\d{1,3}\.\d{1,3}\.\d{1,3}|172\.1[6-9]{1}\.\d{1,3}\.\d{1,3}|172\.2[0-9]{1}\.\d{1,3}\.\d{1,3}|172\.3[0-1]{1}\.\d{1,3}\.\d{1,3}
*/
	ProxyIpsRegex interface{} `yaml:"proxy_ips_regex,omitempty"`

	/*Zones - Descr: A list of hostnames that are routed to the UAA, specifically the default zone in the UAA. The UAA will reject any Host headers that it doesn't recognize.
By default the UAA recognizes:
  The hostname from the property uaa.url
  The hostname from the property login.url
  localhost (in order to accept health checks)
Any hostnames added as a list are additive to the default hostnames allowed.
 Default: <nil>
*/
	Zones *Zones `yaml:"zones,omitempty"`

	/*Database - Descr: Timeout in seconds for the longest running queries. Take into DB migrations for this timeout as they may run during a long period of time. Default: 300
*/
	Database *Database `yaml:"database,omitempty"`

	/*LoggingLevel - Descr: Set UAA logging level.  (e.g. TRACE, DEBUG, INFO) Default: DEBUG
*/
	LoggingLevel interface{} `yaml:"logging_level,omitempty"`

	/*Servlet - Descr: Optional configuration of the UAA session cookie.
Defaults are the following key value pairs:
  secure: <(boolean)this value if set, otherwise require_https>
  http-only: <(boolean) - default to true. set HttpOnly flag on cookie.
  max-age: <(int) lifetime in seconds of cookie - default to 30 minutes)
  name: <(String) name of cookie, default is JSESSIONID>
  comment: <(String) optional comment in cookie>
  path: <(String) path for cookie, default is />
  domain: <(String) domain for cookie, default is incoming request domain>
 Default: <nil>
*/
	Servlet *Servlet `yaml:"servlet,omitempty"`

	/*Url - Descr: The base url of the UAA Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*Issuer - Descr: The url to use as the issuer URI Default: <nil>
*/
	Issuer interface{} `yaml:"issuer,omitempty"`

	/*LoggingUseRfc3339 - Descr: Sets the time format for log messages to be yyyy-MM-dd'T'HH:mm:ss.SSSXXX instead of yyyy-MM-dd HH:mm:ss.SSS Default: false
*/
	LoggingUseRfc3339 interface{} `yaml:"logging_use_rfc3339,omitempty"`

	/*SslPrivateKey - Descr: The server's ssl private key. Only passphrase-less keys are supported Default: 
*/
	SslPrivateKey interface{} `yaml:"sslPrivateKey,omitempty"`

	/*Password - Descr: Maximum number of characters required for password to be considered valid Default: 255
*/
	Password *UaaPassword `yaml:"password,omitempty"`

	/*Newrelic - Descr: To enable newrelic monitoring, the sub element of this property will be placed in
a configuration file called newrelic.yml in the jobs config directory.
The syntax that must adhere to documentation in https://docs.newrelic.com/docs/agents/java-agent/configuration/java-agent-configuration-config-file
The JVM option -javaagent:/path/to/newrelic.jar will be added to Apache Tomcat's startup script
The enablement of the NewRelic agent in the UAA is triggered by the property uaa.newrelic.common.license_key
The property uaa.newrelic.common.license_key must be set!
 Default: <nil>
*/
	Newrelic interface{} `yaml:"newrelic,omitempty"`

	/*Port - Descr: Port that uaa will accept connections on Default: 8080
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Admin - Descr: Secret of the admin client - a client named admin with uaa.admin as an authority Default: <nil>
*/
	Admin *Admin `yaml:"admin,omitempty"`

	/*Login - Descr: Default login client secret, if no login client is defined Default: <nil>
*/
	Login *UaaLogin `yaml:"login,omitempty"`

	/*DisableInternalUserManagement - Descr: Disables UI and API for internal user management Default: false
*/
	DisableInternalUserManagement interface{} `yaml:"disableInternalUserManagement,omitempty"`

	/*Authentication - Descr: Number of seconds to lock out an account when lockoutAfterFailures failures is exceeded Default: 300
*/
	Authentication *Authentication `yaml:"authentication,omitempty"`

	/*DisableInternalAuth - Descr: Disables internal user authentication Default: false
*/
	DisableInternalAuth interface{} `yaml:"disableInternalAuth,omitempty"`

}