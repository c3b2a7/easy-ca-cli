package cli

import (
	"crypto/elliptic"
	"crypto/x509"
	"github.com/c3b2a7/easy-ca/ca"
)

type CertConfig struct {
	Subject   string
	ValidFrom string
	ValidFor  int

	IssuerCertPath       string
	IssuerPrivateKeyPath string
	CertOutputPath       string
	PrivateKeyOutputPath string

	RSA        bool
	ECDSA      bool
	ED25591    bool
	RSAKeySize int
	ECDSACurve string
}

func (c *CertConfig) IssuerCertificate() (*x509.Certificate, error) {
	blocks, err := LoadPem(c.IssuerCertPath)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificate(blocks.Bytes)
}

func (c *CertConfig) IssuerPrivateKey() (interface{}, error) {
	blocks, err := LoadPem(c.IssuerPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS8PrivateKey(blocks.Bytes)
}

func (c *CertConfig) CertificateOpts() ([]ca.CertificateOption, error) {
	var opts []ca.CertificateOption
	opts = append(opts, ca.WithSubject(c.Subject))

	if c.IssuerCertPath != "" && c.IssuerPrivateKeyPath != "" {
		issuer, err := c.IssuerCertificate()
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		privateKey, err := c.IssuerPrivateKey()
		if err != nil {
			return nil, err
		}
		opts = append(opts, ca.WithIssuer(issuer))
		opts = append(opts, ca.WithIssuerPrivateKey(privateKey))
	}

	return opts, nil
}

func (c *CertConfig) KeyOpts() []ca.KeyOption {
	var keyOpts []ca.KeyOption

	if c.RSA {
		keyOpts = append(keyOpts, ca.WithKeySize(c.RSAKeySize))
	} else if c.ECDSA {
		var curve elliptic.Curve
		switch c.ECDSACurve {
		case "P224":
			curve = elliptic.P224()
		case "P256":
			curve = elliptic.P256()
		case "P384":
			curve = elliptic.P384()
		case "P521":
			curve = elliptic.P521()
		}
		keyOpts = append(keyOpts, ca.WithCurve(curve))
	}

	return keyOpts
}

func (c *CertConfig) Algorithm() (algorithm string) {
	if c.RSA {
		algorithm = "RSA"
	} else if c.ECDSA {
		algorithm = "ECDSA"
	} else if c.ED25591 {
		algorithm = "ED25591"
	}
	return
}

func (c *CertConfig) GenKeyPair() (keyPair ca.KeyPair, err error) {
	var kpg ca.KeyPairGenerator
	if kpg, err = ca.GetKeyPairGenerator(c.Algorithm(), c.KeyOpts()...); err != nil {
		return
	}
	return kpg.GenerateKeyPair()
}
