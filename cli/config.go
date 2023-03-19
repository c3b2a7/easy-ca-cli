/*
Copyright © 2023 c3b2a <c3b2a@qq.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
