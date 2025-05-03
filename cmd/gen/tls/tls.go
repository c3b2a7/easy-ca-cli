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
package tls

import (
	"crypto/x509"
	"github.com/c3b2a7/easy-ca-cli/cli"
	"github.com/c3b2a7/easy-ca-cli/cmd/gen/flags"
	"github.com/c3b2a7/easy-ca/ca"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

type tlsConfig struct {
	*cli.CertConfig
	Host string
}

func NewCmdGenTLS(cfg *cli.CertConfig) *cobra.Command {
	tlsCfg := &tlsConfig{CertConfig: cfg}

	cmd := &cobra.Command{
		Use:   "tls",
		Short: "Generate a tls certificate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenTLS(tlsCfg)
		},
	}

	flags.ApplyCommonFlags(cmd, cfg)
	cmd.Flags().IntVar(&cfg.Days, "days", 825, "days that certificate is valid for")
	cmd.Flags().StringVar(&tlsCfg.Host, "host", "", "comma-separated hostnames and IPs to generate a certificate for")

	return cmd
}

func runGenTLS(cfg *tlsConfig) error {
	certificateOpts, err := cfg.CertificateOpts()
	if err != nil {
		return err
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
		certificateOpts = append(certificateOpts, ca.WithIPs(ips), ca.WithDomains(domains))
	}
	certificateOpts = append(certificateOpts, ca.WithCA(false))

	keyPair, err := cfg.GenKeyPair()
	if err != nil {
		return err
	}

	var certificate *x509.Certificate
	certificate, err = ca.CreateGeneralCertificate(keyPair, certificateOpts...)
	if err != nil {
		return err
	}

	issuerCertificateChain := cli.Must(cfg.IssuerCertificateChain()).([]*x509.Certificate)
	return cli.Out(cfg.CertConfig, append([]*x509.Certificate{certificate}, issuerCertificateChain...), keyPair)
}
