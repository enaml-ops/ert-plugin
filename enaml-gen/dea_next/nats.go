package dea_next 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Password - Descr: password for NATS login Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Machines - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

	/*User - Descr: user name for NATS login Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Port - Descr: TCP port of NATS server Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

}