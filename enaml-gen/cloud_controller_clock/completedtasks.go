package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CompletedTasks struct {

	/*CutoffAgeInDays - Descr: How long a completed task will stay in cloud controller database before being cleaned up based on last updated time with success or failure. Default: 31
*/
	CutoffAgeInDays interface{} `yaml:"cutoff_age_in_days,omitempty"`

}