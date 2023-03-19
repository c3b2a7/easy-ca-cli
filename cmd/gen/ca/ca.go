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
package ca

import (
	"crypto/x509"
	"github.com/c3b2a7/easy-ca-cli/cli"
	"github.com/c3b2a7/easy-ca-cli/cmd/gen/flags"
	"github.com/c3b2a7/easy-ca/ca"
	"github.com/spf13/cobra"
)

func NewCmdGenCA(cfg *cli.CertConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "Generate a certificate authority",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCA(cfg)
		},
	}

	flags.ApplyCommonFlags(cmd, cfg)
	cmd.Flags().IntVar(&cfg.ValidFor, "duration", 365*20, "duration that certificate is valid for")

	return cmd
}

func runGenCA(cfg *cli.CertConfig) error {
	certificateOpts, err := cfg.CertificateOpts()
	if err != nil {
		return err
	}
	certificateOpts = append(certificateOpts, ca.WithCA(true))

	keyPair, err := cfg.GenKeyPair()
	if err != nil {
		return err
	}

	var certificate *x509.Certificate
	if cfg.IssuerCertPath != "" && cfg.IssuerPrivateKeyPath != "" {
		certificate, err = ca.CreateMiddleRootCertificate(keyPair, certificateOpts...)
	} else {
		certificate, err = ca.CreateSelfSignedRootCertificate(keyPair, certificateOpts...)
	}
	if err != nil {
		return err
	}

	return cli.Out(cfg, []*x509.Certificate{certificate}, keyPair)
}
