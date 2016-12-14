package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SyslogDaemonConfig struct {

	/*MaxMessageSize - Descr: maximum message size to be sent Default: 4k
*/
	MaxMessageSize interface{} `yaml:"max_message_size,omitempty"`

	/*Port - Descr: TCP port of syslog aggregator Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Enable - Descr: Enable or disable rsyslog configuration for forwarding syslog messages into metron Default: true
*/
	Enable interface{} `yaml:"enable,omitempty"`

	/*Transport - Descr: Transport to be used when forwarding logs (tcp|udp|relp). Default: tcp
*/
	Transport interface{} `yaml:"transport,omitempty"`

	/*CustomRule - Descr: Custom rule for syslog forward daemon Default: 
*/
	CustomRule interface{} `yaml:"custom_rule,omitempty"`

	/*FallbackAddresses - Descr: Addresses of fallback servers to be used if the primary syslog server is down. Only tcp or relp are supported. Each list entry should consist of "address", "transport" and "port" keys.  Default: []
*/
	FallbackAddresses interface{} `yaml:"fallback_addresses,omitempty"`

	/*Address - Descr: IP address for syslog aggregator Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

}