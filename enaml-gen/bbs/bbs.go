package bbs 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Bbs struct {

	/*Etcd - Descr: capacity of the tls client cache Default: <nil>
*/
	Etcd *Etcd `yaml:"etcd,omitempty"`

	/*DebugAddr - Descr: address at which to serve debug info Default: 127.0.0.1:17017
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*ListenAddr - Descr: address at which to serve API requests Default: 0.0.0.0:8889
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

	/*CaCert - Descr: PEM-encoded CA certificate Default: <nil>
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*HealthAddr - Descr: address at which to serve API requests Default: 127.0.0.1:8890
*/
	HealthAddr interface{} `yaml:"health_addr,omitempty"`

	/*ActiveKeyLabel - Descr: Label of the encryption key to be used when writing to the database Default: <nil>
*/
	ActiveKeyLabel interface{} `yaml:"active_key_label,omitempty"`

	/*Convergence - Descr: unclaimed tasks are marked as failed, after this duration in seconds Default: 1800
*/
	Convergence *Convergence `yaml:"convergence,omitempty"`

	/*ServerCert - Descr: PEM-encoded client certificate Default: <nil>
*/
	ServerCert interface{} `yaml:"server_cert,omitempty"`

	/*Sql - Descr: Whether to require SSL for BBS communication to the SQL backend Default: false
*/
	Sql *Sql `yaml:"sql,omitempty"`

	/*AdvertisementBaseHostname - Descr: Suffix for the BBS advertised hostname Default: bbs.service.cf.internal
*/
	AdvertisementBaseHostname interface{} `yaml:"advertisement_base_hostname,omitempty"`

	/*DesiredLrpCreationTimeout - Descr: expected maximum time to create all components of a desired LRP Default: <nil>
*/
	DesiredLrpCreationTimeout interface{} `yaml:"desired_lrp_creation_timeout,omitempty"`

	/*RequireSsl - Descr: require ssl for all communication the bbs Default: true
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

	/*EnableAccessLog - Descr: Enable access log, i.e. log every request made to the bbs Default: false
*/
	EnableAccessLog interface{} `yaml:"enable_access_log,omitempty"`

	/*ServerKey - Descr: PEM-encoded client key Default: <nil>
*/
	ServerKey interface{} `yaml:"server_key,omitempty"`

	/*EncryptionKeys - Descr: List of encryption keys to be used Default: []
*/
	EncryptionKeys interface{} `yaml:"encryption_keys,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*Auctioneer - Descr: Address of the auctioneer API Default: http://auctioneer.service.cf.internal:9016
*/
	Auctioneer *Auctioneer `yaml:"auctioneer,omitempty"`

}