package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version of ncli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("‎ _   _        _                 _   __ \n| \\ | |  ___ | |_   __ _  _ __ (_) / _|\n|  \\| | / _ \\| __| / _` || '__|| || |_\t\n| |\\  ||  __/| |_ | (_| || |   | ||  _|\n|_| \\_| \\___| \\__| \\__, ||_|   |_||_|\t\n                   |___/\t\t\t\t")
			fmt.Println("Netgrif CLI v1.0.0 ")
		},
	}
}
