package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ResourcePool struct {

	/*Cdn - Descr: URI for a CDN to used for resource pool downloads Default: 
*/
	Cdn *ResourcePoolCdn `yaml:"cdn,omitempty"`

	/*MinimumSize - Descr: Minimum size of a resource to add to the pool Default: 65536
*/
	MinimumSize interface{} `yaml:"minimum_size,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*ResourceDirectoryKey - Descr: Directory (bucket) used store app resources.  It does not have be pre-created. Default: cc-resources
*/
	ResourceDirectoryKey interface{} `yaml:"resource_directory_key,omitempty"`

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*MaximumSize - Descr: Maximum size of a resource to add to the pool Default: 536870912
*/
	MaximumSize interface{} `yaml:"maximum_size,omitempty"`

	/*WebdavConfig - Descr: The ca cert to use when communicating with webdav Default: 
*/
	WebdavConfig *ResourcePoolWebdavConfig `yaml:"webdav_config,omitempty"`

}