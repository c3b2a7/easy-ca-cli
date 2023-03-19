package gen

import (
	"github.com/c3b2a7/easy-ca-cli/cli"
	cmdCA "github.com/c3b2a7/easy-ca-cli/cmd/gen/ca"
	cmdTLS "github.com/c3b2a7/easy-ca-cli/cmd/gen/tls"
	"github.com/spf13/cobra"
)

func NewGenCmd(cfg *cli.CertConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen [<command>]",
		Short: "Generate certificates",
		Long:  "Generate a certificate with your specified params",
	}

	cmd.AddCommand(cmdCA.NewCmdGenCA(cfg))
	cmd.AddCommand(cmdTLS.NewCmdGenTLS(cfg))

	return cmd
}
