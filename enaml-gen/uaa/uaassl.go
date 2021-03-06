package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type UaaSsl struct {

	/*ProtocolHeader - Descr: The header to look for to determine if ssl termination was performed by a front end load balancer. Default: x-forwarded-proto
*/
	ProtocolHeader interface{} `yaml:"protocol_header,omitempty"`

	/*Port - Descr: If this property Tomcat will listen to this port and expect https traffic. If null, tomcat will not listen to this port Default: 8443
*/
	Port interface{} `yaml:"port,omitempty"`

}