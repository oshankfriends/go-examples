package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdNlu(token *string) *cobra.Command {
	var host string
	cmd := &cobra.Command{
		Use:   "nlu",
		Short: "query from NLU service",
		Long:  "get answer of queries from nlu",
	}
	cmd.PersistentFlags().StringVarP(&host, "host", "H", "https://gateway.dev.iamplus.xyz", "address of gateway")
	cmd.AddCommand(NewCmdQuery(token,&host))
	return cmd
}
