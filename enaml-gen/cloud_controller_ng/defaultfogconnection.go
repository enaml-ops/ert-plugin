package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DefaultFogConnection struct {

	/*LocalRoot - Descr: Local root when fog provider is not overridden (should be an NFS mount if using more than one cloud controller) Default: /var/vcap/nfs/shared
*/
	LocalRoot interface{} `yaml:"local_root,omitempty"`

	/*Provider - Descr: Local fog provider (should always be 'Local'), used if fog_connection hash is not provided in the manifest Default: Local
*/
	Provider interface{} `yaml:"provider,omitempty"`

}