package test_autoscaling 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type TestAutoscalingJob struct {

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Domain - Descr: The CF top-level domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Autoscale - Descr: Password of the CF admin user Default: <nil>
*/
	Autoscale *Autoscale `yaml:"autoscale,omitempty"`

}