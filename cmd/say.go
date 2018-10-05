package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var cmdEcho = &cobra.Command{
	Use:   "say [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
		Echo echo's.
	`,
	Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
	fmt.Println(strings.Join(args, " "))
}

func init() {
	RootCmd.AddCommand(cmdEcho)
}
