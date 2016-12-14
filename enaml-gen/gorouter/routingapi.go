package gorouter 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RoutingApi struct {

	/*AuthDisabled - Descr: Disables UAA authentication Default: false
*/
	AuthDisabled interface{} `yaml:"auth_disabled,omitempty"`

	/*Port - Descr: Port on which routing-api is running. Default: 3000
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Enabled - Descr: Enable the GoRouter to receive routes from the Routing API Default: false
*/
	Enabled interface{} `yaml:"enabled,omitempty"`

}