package config_test

import (
	"io/ioutil"

	. "github.com/enaml-ops/ert-plugin/plugin/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetFingerPrint", func() {
	const controlFingerprint = "33:45:25:3c:73:9e:09:74:f8:e9:2b:96:1c:45:76:05"
	var err error
	var controlPublicKey []byte
	var fingerprint string

	BeforeEach(func() {
		controlPublicKey, err = ioutil.ReadFile("fixtures/fakepublic.key")
		fingerprint, err = GetFingerPrint(string(controlPublicKey))
		Ω(err).ShouldNot(HaveOccurred())
	})

	Context("when called with a valid public key string", func() {
		It("should create the valid md5 fingerprint for the public key", func() {
			Ω(fingerprint).Should(Equal(controlFingerprint), "ssh-keygen -lf fixtures/fakepublic.key -E md5 | cut -d ' ' -f2 | sed 's/MD5://g'")
		})
	})
})
