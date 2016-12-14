package blobstore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Blobstore struct {

	/*MaxUploadSize - Descr: Max allowed file size for upload Default: 5000m
*/
	MaxUploadSize interface{} `yaml:"max_upload_size,omitempty"`

	/*Port - Descr: TCP port on which the blobstore server (nginx) listens Default: 80
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Tls - Descr: The TCP port on which the internal blobstore server listens Default: 443
*/
	Tls *Tls `yaml:"tls,omitempty"`

	/*AdminUsers - Descr: List of Username and Password pairs that have admin access to the blobstore. Cloud Controller must use one of these to access the blobstore via HTTP Basic Auth.
Example:
  users:
  - username: user1
    password: password1
  - username: user2
    password: password2
 Default: <nil>
*/
	AdminUsers interface{} `yaml:"admin_users,omitempty"`

	/*SecureLink - Descr: The secret used for signing URLs Default: <nil>
*/
	SecureLink *SecureLink `yaml:"secure_link,omitempty"`

}