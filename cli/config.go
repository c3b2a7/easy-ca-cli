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
	"fmt"
	"github.com/c3b2a7/easy-ca/ca"
	"time"
)

type CertConfig struct {
	Subject   string
	StartDate string
	Days      int

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

func (c *CertConfig) IssuerCertificateChain() ([]*x509.Certificate, error) {
	blocks, err := LoadBlocks(c.IssuerCertPath)
	if err != nil {
		return nil, err
	}
	var chain []*x509.Certificate
	for _, block := range blocks {
		var cert *x509.Certificate
		cert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		chain = append(chain, cert)
	}
	return chain, nil
}

func (c *CertConfig) IssuerPrivateKey() (interface{}, error) {
	blocks, err := LoadBlock(c.IssuerPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	if blocks.Type == "EC PRIVATE KEY" {
		return x509.ParseECPrivateKey(blocks.Bytes)
	} else if blocks.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(blocks.Bytes)
	} else {
		priv, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
		if err == nil {
			return priv, err
		}
		return x509.ParsePKCS1PrivateKey(blocks.Bytes)
	}
}

func (c *CertConfig) CertificateOpts() ([]ca.CertificateOption, error) {
	var opts []ca.CertificateOption
	opts = append(opts, ca.WithSubject(c.Subject))

	if c.IssuerCertPath != "" && c.IssuerPrivateKeyPath != "" {
		issuerCertChain, err := c.IssuerCertificateChain()
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
		opts = append(opts, ca.WithIssuer(issuerCertChain[0]))
		opts = append(opts, ca.WithIssuerPrivateKey(privateKey))
	}

	var notBefore time.Time
	var err error
	if len(c.StartDate) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.ParseInLocation(time.DateTime, c.StartDate, time.Local)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse creation date: %v", err)
	}
	opts = append(opts, ca.WithNotBefore(notBefore))
	if c.Days != 0 {
		opts = append(opts, ca.WithNotAfter(notBefore.AddDate(0, 0, c.Days)))
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
