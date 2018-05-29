package cmd

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/oshankfriends/go-examples/gatewayCli/nluclient"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdQuery(token, host *string) *cobra.Command {
	queryReq := &nluclient.QueryRequest{}
	cmd := &cobra.Command{
		Use:   "query",
		Short: "ask query to NLU",
		Long: `command line tool for sending query to NLU.
This an alternative of calling rest API of gateway`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nluclient.DefaultClient.SetBase(host)
			_, resp, err := client.Query(*token, queryReq)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if resp != nil {
				fmt.Println()
				pretty.Printf("%# v\n", *resp)
			}
		},
	}
	cmd.Flags().StringVar(&queryReq.Input.Text, "input", "", "input text")
	cmd.Flags().StringVar(&queryReq.User.Id, "id", "iamplus.wrapper.test01@gmail.com", "user id")
	return cmd
}
