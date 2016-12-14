package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type NfsServer struct {

	/*Address - Descr: NFS server for droplets and apps (not used in an AWS deploy, use s3 instead) Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*SharePath - Descr: The location at which to mount the nfs share Default: /var/vcap/nfs
*/
	SharePath interface{} `yaml:"share_path,omitempty"`

}