package cloudfoundry

import (
	"github.com/enaml-ops/enaml"
	nfsmounterlib "github.com/enaml-ops/ert-plugin/enaml-gen/nfs_mounter"
	"github.com/enaml-ops/ert-plugin/plugin/plugin/config"
)

//CreateNFSMounterJob - Create the yaml job structure for NFSMounter
func CreateNFSMounterJob(config *config.Config) enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "nfs_mounter",
		Release: "cf",
		Properties: &nfsmounterlib.NfsMounterJob{
			NfsServer: &nfsmounterlib.NfsServer{
				Address: config.NFSIP,
				Share:   config.SharePath,
			},
		},
	}
}
