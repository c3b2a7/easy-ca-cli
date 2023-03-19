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
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/c3b2a7/easy-ca/ca"
	"os"
)

func Must(v interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return v
}

func Out(cfg *CertConfig, certificates []*x509.Certificate, keyPair ca.KeyPair) error {
	certFile, err := os.OpenFile(cfg.CertOutputPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer certFile.Close()
	err = ca.EncodeCertificateChain(certFile, certificates)
	if err != nil {
		return err
	}

	privateKeyFile, err := os.OpenFile(cfg.PrivateKeyOutputPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()
	return ca.EncodePKCS8PrivateKey(privateKeyFile, keyPair.PrivateKey)
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
