package main

import (
	"fmt"
	"github.com/oshankfriends/go-examples/gatewayCli/pkg/gatewayctl/cmd"
	"os"
)

func main() {
	command := cmd.NewGatewayctlCommand()
	if err := command.Execute(); err != nil {
		fmt.Errorf("%s\n", err)
		os.Exit(1)
	}
}
