package main

import (
	"log"

	"github.com/spf13/pflag"

	"github.com/alfranz/cardgen/cmd"
)

func init() {
	flags := pflag.NewFlagSet("cardgen", pflag.ExitOnError)
	pflag.CommandLine = flags
}

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
