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
package flags

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/c3b2a7/easy-ca-cli/cli"
)

func ApplyCommonFlags(cmd *cobra.Command, cfg *cli.CertConfig) {
	cmd.Flags().StringVar(&cfg.Subject, "subject", "", `subject of certificate, formatted as /C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Root`)
	cmd.Flags().StringVar(&cfg.StartDate, "start-date", "", "starting date of certificate, formatted as "+time.DateTime)

	cmd.Flags().BoolVar(&cfg.RSA, "rsa", false, "rsa algorithm")
	cmd.Flags().IntVar(&cfg.RSAKeySize, "rsa-keysize", 3072, "rsa key size, valid values are 2048, 3072, 4096")
	cmd.Flags().BoolVar(&cfg.ECDSA, "ecdsa", false, "ecdsa algorithm")
	cmd.Flags().StringVar(&cfg.ECDSACurve, "ecdsa-curve", "P384", "ecdsa curve to use to generate a key, valid values are P224, P256, P384, P521")
	cmd.Flags().BoolVar(&cfg.ED25519, "ed25519", false, "ed25519 algorithm")

	cmd.Flags().StringVar(&cfg.IssuerCertPath, "issuer-cert", "", "certificate file of issuer")
	cmd.Flags().StringVar(&cfg.IssuerPrivateKeyPath, "issuer-key", "", "private key file of issuer")
	cmd.Flags().StringVar(&cfg.CertOutputPath, "out-cert", "", "certificate file output location")
	cmd.Flags().StringVar(&cfg.PrivateKeyOutputPath, "out-key", "", "private key file output location")
	cmd.Flags().StringVar(&cfg.PKCS12OutputPath, "out-pkcs12", "", "pkcs12 file output location")
	cmd.Flags().StringVar(&cfg.PKCS12Password, "pkcs12-password", "", "password for pkcs12 file")

	cli.Must(cmd.MarkFlagFilename("issuer-cert"))
	cli.Must(cmd.MarkFlagFilename("issuer-key"))
	cli.Must(cmd.MarkFlagFilename("out-cert"))
	cli.Must(cmd.MarkFlagFilename("out-key"))

	cli.Must(cmd.MarkFlagRequired("subject"))
	cmd.MarkFlagsRequiredTogether("out-pkcs12", "pkcs12-password")
}
