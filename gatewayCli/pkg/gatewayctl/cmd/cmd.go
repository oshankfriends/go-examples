package cmd

import "github.com/spf13/cobra"

func NewGatewayctlCommand() *cobra.Command {
	var authToken string
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "gatewayctl is commandline tool to send query to gateway",
		Long:  `gatewayctl is commandline tool to send query to various omega skill via gateway`,
	}
	cmd.PersistentFlags().StringVar(&authToken, "token", "4a16c1af6e22eba2b1d7a30b08229519377b77ed6b20f82d4ad76fd29cf88756", "Authentication token")
	cmd.AddCommand(NewCmdNlu(&authToken))
	return cmd
}
