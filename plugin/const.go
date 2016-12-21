package cloudfoundry

const (
	DefaultHM9000Port = 15487
	CFReleaseName     = "cf"
	CFReleaseVersion  = "239.0.40"

	PivotalERTVersion = "1.8.21"

	StemcellName    = "ubuntu-trusty"
	StemcellAlias   = "trusty"
	StemcellVersion = "3263.12"

	CFLinuxReleaseName    = "cflinuxfs2-rootfs"
	CFLinuxReleaseVersion = "1.40.0"

	GardenReleaseName    = "garden-runc"
	GardenReleaseVersion = "1.0.2"

	DiegoReleaseName    = "diego"
	DiegoReleaseVersion = "0.1485.8"

	CFMysqlReleaseName    = "cf-mysql"
	CFMysqlReleaseVersion = "26.8.0"

	EtcdReleaseName    = "etcd"
	EtcdReleaseVersion = "60.0.1"

	PushAppsReleaseName    = "push-apps-manager-release"
	PushAppsReleaseVersion = "652.0.0"

	NotificationsReleaseName    = "notifications"
	NotificationsReleaseVersion = "24.0.0"

	NotificationsUIReleaseName    = "notifications-ui"
	NotificationsUIReleaseVersion = "22.0.0"

	CFAutoscalingReleaseName    = "cf-autoscaling"
	CFAutoscalingReleaseVersion = "36.0.0"

	MySQLBackupReleaseName    = "mysql-backup"
	MySQLBackupReleaseVersion = "11.25.0"

	ServiceBackupReleaseVersion = "14.0.0"
)

const (
	DeploymentName = "cf"

	defaultBBSAPILocation = "bbs.service.cf.internal:8889"

	javaBuildpackName    = "java-offline-buildpack"
	javaBuildpackPackage = "buildpack_java_offline"
)

var flagsToInferFromCloudConfig = map[string][]string{
	"disktype": []string{
		"mysql-disk-type",
		"diego-db-disk-type",
		"diego-cell-disk-type",
		"diego-brain-disk-type",
		"etcd-disk-type",
		"nfs-disk-type",
	},
	"vmtype": []string{
		"diego-brain-vm-type",
		"errand-vm-type",
		"clock-global-vm-type",
		"doppler-vm-type",
		"uaa-vm-type",
		"diego-cell-vm-type",
		"diego-db-vm-type",
		"router-vm-type",
		"haproxy-vm-type",
		"nats-vm-type",
		"consul-vm-type",
		"etcd-vm-type",
		"nfs-vm-type",
		"mysql-vm-type",
		"mysql-proxy-vm-type",
		"cc-worker-vm-type",
		"cc-vm-type",
		"loggregator-traffic-controller-vmtype",
	},
	"az":      []string{"az"},
	"network": []string{"network"},
}
