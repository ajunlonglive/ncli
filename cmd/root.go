package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{
	Use:   "ncli",
	Short: "NCLI is best friend of Netgrif developer.",
	Long:  `Netgrif CLI tool is aggregation of commands and utility functions for better experience to work with Netgrif products and developing an application based on Netgrif Application Engine.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
