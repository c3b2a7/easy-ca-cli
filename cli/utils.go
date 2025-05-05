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
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/c3b2a7/easy-ca/ca"
	"os"
	"software.sslmate.com/src/go-pkcs12"
)

func Must(v interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return v
}

func Out(cfg *CertConfig, certificates []*x509.Certificate, keyPair ca.KeyPair) error {
	privateKeyFilePath, certFilePath := checkAndGenerateOutputPath(cfg, certificates)
	certFile, err := os.OpenFile(certFilePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer certFile.Close()

	if err = ca.EncodeCertificateChain(certFile, certificates); err != nil {
		return err
	}

	privateKeyFile, err := os.OpenFile(privateKeyFilePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	if err = ca.EncodePKCS8PrivateKey(privateKeyFile, keyPair.PrivateKey); err != nil {
		return err
	}

	if cfg.PKCS12OutputPath != "" {
		if err = writePKCS12File(cfg, certificates, keyPair); err != nil {
			return err
		}
	}
	return nil
}

func writePKCS12File(cfg *CertConfig, certificates []*x509.Certificate, keyPair ca.KeyPair) error {
	cert := certificates[0]
	var caCerts []*x509.Certificate
	if len(certificates) > 1 {
		caCerts = certificates[1:]
	}

	var pfxData []byte
	var err error
	if pfxData, err = pkcs12.Legacy.Encode(keyPair.PrivateKey, cert, caCerts, cfg.PKCS12Password); err != nil {
		return err
	}
	pfxFile, err := os.OpenFile(cfg.PKCS12OutputPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer pfxFile.Close()
	_, err = pfxFile.Write(pfxData)
	return err
}

func checkAndGenerateOutputPath(cfg *CertConfig, certificates []*x509.Certificate) (string, string) {
	if len(certificates) == 0 {
		panic("certificates must not be empty")
	}

	privateKeyFile, certFile := cfg.PrivateKeyOutputPath, cfg.CertOutputPath
	if privateKeyFile != "" && certFile != "" {
		return privateKeyFile, certFile
	}

	cert := certificates[0]
	commonName := cert.Subject.CommonName
	if cert.Subject.CommonName == "" {
		commonName = DetermineCertificateKind(cert).String()
	}

	if privateKeyFile == "" {
		privateKeyFile = commonName + ".key.pem"
	}
	if certFile == "" {
		certFile = commonName + ".cert.pem"
	}

	return privateKeyFile, certFile
}

type CertificateKind int8

const (
	TLS CertificateKind = 1 << iota
	IntermediateCA
	RootCA
)

func (k CertificateKind) String() string {
	switch k {
	case TLS:
		return "tls"
	case IntermediateCA:
		return "intermediate_ca"
	case RootCA:
		return "ca"
	default:
		return "unknown"
	}
}

func DetermineCertificateKind(cert *x509.Certificate) CertificateKind {
	if cert == nil {
		panic("certificate is nil")
	}
	if cert.IsCA {
		if bytes.Equal(cert.RawSubject, cert.RawIssuer) {
			return RootCA
		}
		return IntermediateCA
	} else {
		return TLS
	}
}

func LoadBlock(file string) (*pem.Block, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return block, nil
}

func LoadBlocks(file string) ([]*pem.Block, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var blocks []*pem.Block
	for len(b) > 0 {
		block, rest := pem.Decode(b)
		blocks = append(blocks, block)
		if bytes.Equal(b, rest) {
			return nil, fmt.Errorf("invalid PEM block in file %s", file)
		}
		b = rest
	}
	return blocks, nil
}

func PublicKey(priv any) any {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}
