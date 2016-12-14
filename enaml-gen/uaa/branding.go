package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Branding struct {

	/*FooterLinks - Descr: These links appear on the footer of all UAA pages. You may choose to add multiple urls for things like Support, Terms of Service etc. Default: <nil>
*/
	FooterLinks interface{} `yaml:"footer_links,omitempty"`

	/*FooterLegalText - Descr: This text appears on the footer of all UAA pages Default: <nil>
*/
	FooterLegalText interface{} `yaml:"footer_legal_text,omitempty"`

	/*CompanyName - Descr: This name is used on the UAA Pages and in account management related communication in UAA Default: <nil>
*/
	CompanyName interface{} `yaml:"company_name,omitempty"`

	/*SquareLogo - Descr: This is a base64 encoded PNG image which will be used as the favicon for the UAA pages Default: <nil>
*/
	SquareLogo interface{} `yaml:"square_logo,omitempty"`

	/*ProductLogo - Descr: This is a base64 encoded PNG image which will be used as the logo on all UAA pages like Login, Sign Up etc. Default: <nil>
*/
	ProductLogo interface{} `yaml:"product_logo,omitempty"`

}