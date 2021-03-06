package destroy_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Autoscale struct {

	/*Cf - Descr: Username of the CF admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*Broker - Descr: Broker basic auth password Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*Organization - Descr: Organization where app is deployed Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

	/*Space - Descr: Space where app is deployed Default: <nil>
*/
	Space interface{} `yaml:"space,omitempty"`

}