package auctioneer 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AuctioneerJob struct {

	/*Diego - Descr: address where auctioneer listens for LRP and task start auction requests Default: 0.0.0.0:9016
*/
	Diego *Diego `yaml:"diego,omitempty"`

}