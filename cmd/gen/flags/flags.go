package flags

import (
	"github.com/c3b2a7/easy-ca-cli/cli"
	"github.com/c3b2a7/easy-ca/ca/constants"
	"github.com/spf13/cobra"
)

func ApplyCommonFlags(cmd *cobra.Command, cfg *cli.CertConfig) {
	cmd.Flags().StringVar(&cfg.Subject, "subject", "", "certificate subject formatted as "+constants.DefaultCASubject)
	cmd.Flags().StringVar(&cfg.ValidFrom, "start-date", "", "creation date formatted as Jan 1 15:04:05 2022")

	cmd.Flags().BoolVar(&cfg.RSA, "rsa", true, "rsa algorithm")
	cmd.Flags().IntVar(&cfg.RSAKeySize, "rsa-keysize", 2048, "rsa key size")
	cmd.Flags().BoolVar(&cfg.ECDSA, "ecdsa", false, "ecdsa algorithm")
	cmd.Flags().StringVar(&cfg.ECDSACurve, "ecdsa-curve", "P256", "ecdsa curve to use to generate a key. Valid values are P224, P256, P384, P521")
	cmd.Flags().BoolVar(&cfg.ED25591, "ed25591", false, "ed25591 algorithm")

	cmd.Flags().StringVar(&cfg.IssuerCertPath, "cert", "", "certificate file of issuer")
	cmd.Flags().StringVar(&cfg.IssuerPrivateKeyPath, "key", "", "private key file of issuer")
	cmd.Flags().StringVar(&cfg.CertOutputPath, "out-cert", "cert.pem", "certificate file output location")
	cmd.Flags().StringVar(&cfg.PrivateKeyOutputPath, "out-key", "key.pem", "private key file output location")

	cmd.MarkFlagRequired("subject")
}
