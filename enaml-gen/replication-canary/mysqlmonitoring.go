package replication_canary 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MysqlMonitoring struct {

	/*ReplicationCanary - Descr: Username of the UAA client used to create the notifications client Default: admin
*/
	ReplicationCanary *ReplicationCanary `yaml:"replication-canary,omitempty"`

	/*NotifyOnly - Descr: When true, replication failure will not shut traffic off at the proxy Default: false
*/
	NotifyOnly interface{} `yaml:"notify_only,omitempty"`

	/*RecipientEmail - Descr: The email address to send mysql monitoring notifications to Default: <nil>
*/
	RecipientEmail interface{} `yaml:"recipient_email,omitempty"`

}