package config

type Config struct {
	IsCA      bool
	Subject   string
	Host      string
	ValidFrom string
	ValidFor  int

	RSA        bool
	ECDSA      bool
	ED25591    bool
	RSAKeySize int
	ECDSACurve string

	IssuerCertPath       string
	IssuerPrivateKeyPath string

	CertOutputPath string
	KeyOutputPath  string
}
