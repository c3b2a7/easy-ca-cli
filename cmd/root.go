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
package cmd

import (
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"
	"github.com/c3b2a7/easy-ca-cli/config"
	"github.com/c3b2a7/easy-ca/ca"
	"github.com/c3b2a7/easy-ca/ca/constants"
	"github.com/spf13/cobra"
	"net"
	"os"
	"strings"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "easy-ca-cli",
	Short: "easy-ca-cli is a very simple certificate generator",
	Long: `A simple and easy to use certificate generator built with love by c3b2a in Go.
Complete documentation is available at https://github.com/c3b2a7/easy-ca-cli#usage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var keyOpts []ca.KeyOption
		var certOpts = []ca.CertificateOption{
			ca.WithCA(cfg.IsCA),
			ca.WithSubject(cfg.Subject),
		}

		if cfg.Host != "" {
			var ips []string
			var domains []string
			for _, h := range strings.Split(cfg.Host, ",") {
				if ip := net.ParseIP(h); ip != nil {
					ips = append(ips, h)
				} else {
					domains = append(domains, h)
				}
			}
			certOpts = append(certOpts, ca.WithIPs(ips), ca.WithDomains(domains))
		}

		var algorithm string
		if cfg.RSA {
			algorithm = "RSA"
			keyOpts = append(keyOpts, ca.WithKeySize(cfg.RSAKeySize))
		} else if cfg.ECDSA {
			algorithm = "ECDSA"
			var curve elliptic.Curve
			switch cfg.ECDSACurve {
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
		} else if cfg.ED25591 {
			algorithm = "ED25591"
		}

		keyPairGenerator, err := ca.GetKeyPairGenerator(algorithm, keyOpts...)
		if err != nil {
			return err
		}
		keyPair, err := keyPairGenerator.GenerateKeyPair()
		if err != nil {
			return err
		}

		var certificate *x509.Certificate
		if cfg.IsCA {
			if cfg.IssuerCertPath != "" && cfg.IssuerPrivateKeyPath != "" {
				certBlock, err := loadPem(cfg.IssuerCertPath)
				if err != nil {
					return err
				}
				issuer, err := x509.ParseCertificate(certBlock.Bytes)
				if err != nil {
					return err
				}
				issuerPrivBlock, err := loadPem(cfg.IssuerPrivateKeyPath)
				if err != nil {
					return err
				}
				privateKey, err := x509.ParsePKCS8PrivateKey(issuerPrivBlock.Bytes)
				if err != nil {
					return err
				}
				certOpts = append(certOpts, ca.WithIssuer(issuer))
				certOpts = append(certOpts, ca.WithIssuerPrivateKey(privateKey))
				certificate, err = ca.CreateMiddleRootCertificate(keyPair, certOpts...)
			} else {
				certificate, err = ca.CreateSelfSignedRootCertificate(keyPair, certOpts...)
			}
		} else {
			certificate, err = ca.CreateGeneralCertificate(keyPair, certOpts...)
			if err != nil {
				return err
			}
		}

		certFile, err := os.OpenFile(cfg.CertOutputPath, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer certFile.Close()
		err = ca.EncodeCertificateChain(certFile, []*x509.Certificate{certificate})
		if err != nil {
			return err
		}

		privateKeyFile, err := os.OpenFile(cfg.KeyOutputPath, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer privateKeyFile.Close()
		err = ca.EncodePKCS8PrivateKey(privateKeyFile, keyPair.PrivateKey)
		if err != nil {
			return err
		}
		return nil
	},
}

func loadPem(file string) (*pem.Block, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return block, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&cfg.IsCA, "ca", true, "Whether this cert should be its own Certificate Authority")
	rootCmd.Flags().StringVar(&cfg.Subject, "subject", constants.DefaultCASubject, "Certificate subject formatted as "+constants.DefaultCASubject)
	rootCmd.Flags().StringVar(&cfg.ValidFrom, "start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
	rootCmd.Flags().IntVar(&cfg.ValidFor, "duration", 825, "Duration that certificate is valid for")
	rootCmd.Flags().StringVar(&cfg.Host, "host", "", "Comma-separated hostnames and IPs to generate a certificate for")

	rootCmd.Flags().BoolVar(&cfg.RSA, "rsa", true, "RSA algorithm")
	rootCmd.Flags().IntVar(&cfg.RSAKeySize, "rsa-keysize", 2048, "RSA key size")
	rootCmd.Flags().BoolVar(&cfg.ECDSA, "ecdsa", false, "ECDSA algorithm")
	rootCmd.Flags().StringVar(&cfg.ECDSACurve, "ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256, P384, P521")
	rootCmd.Flags().BoolVar(&cfg.ED25591, "ed25591", false, "ED25591 algorithm")

	rootCmd.Flags().StringVar(&cfg.KeyOutputPath, "key-output", "key.pem", "private key file output location")
	rootCmd.Flags().StringVar(&cfg.CertOutputPath, "cert-output", "cert.pem", "certificate file output location")
}
