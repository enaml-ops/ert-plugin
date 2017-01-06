package config

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/xchapter7x/lo"
	"golang.org/x/crypto/ssh"
	"gopkg.in/urfave/cli.v2"

	"github.com/pkg/errors"
)

func generateCfSshKeypair() (privatekey string, publickey string, err error) {

	pk, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", errors.Wrap(err, "rsa.generateKey failed")
	}

	pkBuffer := bytes.NewBufferString("")
	pkPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}
	if err := pem.Encode(pkBuffer, pkPEM); err != nil {
		return "", "", errors.Wrap(err, "pem.Encode failed")
	}

	pub, err := ssh.NewPublicKey(&pk.PublicKey)
	if err != nil {
		return "", "", errors.Wrap(err, "ssh.NewPublicKey failed")
	}

	public := ssh.MarshalAuthorizedKey(pub)
	return pkBuffer.String(), string(public), nil
}

func NewConfig(c *cli.Context) (*Config, error) {
	var err error
	config := &Config{}
	config.DiegoSSHPrivateKey, config.DiegoSSHPublicKey, err = generateCfSshKeypair()

	if err != nil {
		lo.G.Error("couldn't generate private key for SSH proxy")
		return nil, err
	}
	config.DiegoSSHHostFingerPrint, err = GetFingerPrint(config.DiegoSSHPublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "GetFingerPrint failed")
	}

	err = pcli.UnmarshalFlags(config, c)
	if err != nil {
		return nil, errors.Wrap(err, "UnmarshalFlags failed")
	}

	secret, err := NewSecret(c)
	if err != nil {
		return nil, errors.Wrap(err, "NewSecret failed")
	}
	config.Secret = *secret
	config.User = NewUser(c)

	if config.MySQLProxyExternalHost == "" {
		config.MySQLProxyExternalHost = config.MySQLProxyIPs[0]
	}

	certs, err := NewCerts(c)
	if err != nil {
		return nil, errors.Wrap(err, "NewCerts failed")
	}
	config.Certs = certs
	return config, nil
}

type Config struct {
	DiegoSSHHostFingerPrint       string   `omg:"-"`
	DiegoSSHPublicKey             string   `omg:"-"`
	DiegoSSHPrivateKey            string   `omg:"-"`
	DeploymentName                string   `omg:"deployment-name,optional"`
	AZs                           []string `omg:"az"`
	StemcellName                  string
	NetworkName                   string `omg:"network"`
	SystemDomain                  string
	AppDomains                    []string `omg:"app-domain"`
	AllowSSHAccess                bool     `omg:"allow-app-ssh-access"`
	SkipSSLCertVerify             bool     `omg:"skip-cert-verify"`
	NATSPort                      int      `omg:"nats-port"`
	UAALoginProtocol              string   `omg:"uaa-login-protocol"`
	SyslogAddress                 string   `omg:"syslog-address,optional"`
	SyslogPort                    int      `omg:"syslog-port,optional"`
	SyslogTransport               string   `omg:"syslog-transport,optional"`
	DopplerZone                   string   `omg:"doppler-zone,optional"`
	DopplerMessageDrainBufferSize int      `omg:"doppler-drain-buffer-size"`
	BBSRequireSSL                 bool     `omg:"bbs-require-ssl"`
	CCUploaderJobPollInterval     int      `omg:"cc-uploader-poll-interval"`
	CCExternalPort                int      `omg:"cc-external-port"`
	SelfServiceLinksEnabled       bool     `omg:"uaa-enable-selfservice-links"`
	SignupsEnabled                bool     `omg:"uaa-signups-enabled"`
	CompanyName                   string   `omg:"uaa-company-name"`
	ProductLogo                   string   `omg:"uaa-product-logo"`
	SquareLogo                    string   `omg:"uaa-square-logo"`
	FooterLegalText               string   `omg:"uaa-footer-legal-txt"`
	LDAPUrl                       string   `omg:"uaa-ldap-url,optional"`
	LDAPUserDN                    string   `omg:"uaa-ldap-user-dn,optional"`
	LDAPSearchBase                string   `omg:"uaa-ldap-search-base,optional"`
	LDAPSearchFilter              string   `omg:"uaa-ldap-search-filter,optional"`
	LDAPMailAttributeName         string   `omg:"uaa-ldap-mail-attributename,optional"`
	LDAPEnabled                   bool     `omg:"uaa-ldap-enabled,optional"`
	SharePath                     string   `omg:"nfs-share-path"`
	SupportAddress                string   `omg:"support-address,optional"`
	MinCliVersion                 string   `omg:"min-cli-version"`
	HAProxySkip                   bool     `omg:"skip-haproxy"`
	MySQLProxyExternalHost        string   `omg:"mysql-proxy-external-host,optional"`
	RouterEnableSSL               bool     `omg:"router-enable-ssl"`
	NFSAllowFromNetworkCIDR       []string `omg:"nfs-allow-from-network-cidr"`
	LoggregatorPort               int
	*Certs                        `omg:"-"` // certs are a special case that we handle manually for now
	IP
	VMType
	Disk
	Secret
	InstanceCount
	User
}

func (c *Config) MySQLProxyHost() string {
	return c.MySQLProxyIPs[0]
}
