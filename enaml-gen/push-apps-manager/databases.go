package push_apps_manager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Databases struct {

	/*AppUsageService - Descr: IP of database server for the app usage service Default: <nil>
*/
	AppUsageService *DatabasesAppUsageService `yaml:"app_usage_service,omitempty"`

	/*Console - Descr: Postgres database name for Console application Default: console
*/
	Console *Console `yaml:"console,omitempty"`

}