package test_restore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Mysql struct {

	/*AdminUsername - Descr: Database username Default: root
*/
	AdminUsername interface{} `yaml:"admin_username,omitempty"`

	/*AdminPassword - Descr: Database password Default: password
*/
	AdminPassword interface{} `yaml:"admin_password,omitempty"`

	/*Port - Descr: Database port for contacting mysql Default: 3306
*/
	Port interface{} `yaml:"port,omitempty"`

	/*ClusterIps - Descr: List of IP address of servers which can generate backups Default: <nil>
*/
	ClusterIps interface{} `yaml:"cluster_ips,omitempty"`

}