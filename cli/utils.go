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

func LoadPem(file string) (*pem.Block, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return block, nil
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
