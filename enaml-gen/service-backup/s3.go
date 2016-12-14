package service_backup 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type S3 struct {

	/*EndpointUrl - Descr: URL of the blobstore endpoint. Override this to use an S3-compatible blobstore like RiakS2. Default: https://s3.amazonaws.com
*/
	EndpointUrl interface{} `yaml:"endpoint_url,omitempty"`

	/*BucketName - Descr: Bucket to which backups are uploaded Default: <nil>
*/
	BucketName interface{} `yaml:"bucket_name,omitempty"`

	/*BucketPath - Descr: Path in bucket to which backups are uploaded Default: <nil>
*/
	BucketPath interface{} `yaml:"bucket_path,omitempty"`

	/*SecretAccessKey - Descr: AWS Secret Access Key to access the blobstore Default: <nil>
*/
	SecretAccessKey interface{} `yaml:"secret_access_key,omitempty"`

	/*AccessKeyId - Descr: AWS Access Key to access the blobstore Default: <nil>
*/
	AccessKeyId interface{} `yaml:"access_key_id,omitempty"`

}