package tps 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Listener struct {

	/*DebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17014
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*ListenAddr - Descr: address at which to serve API requests Default: 0.0.0.0:1518
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

}