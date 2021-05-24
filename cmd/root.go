package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"netgrif.com/ncli/cmd/pathsanitize"
	"netgrif.com/ncli/cmd/version"
)

func NewRootCmd(input string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ncli",
		Short: "NCLI is best friend of Netgrif developer.",
		Long:  `Netgrif CLI tool is aggregation of commands and utility functions for better experience to work with Netgrif products and developing an application based on Netgrif Application Engine.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "Input string "+input)
			if err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(version.CreateCommand())
	rootCmd.AddCommand(pathsanitize.CreateCommand())

	return rootCmd
}
