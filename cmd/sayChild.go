package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var subEcho = &cobra.Command{
	Use:   "child [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
		Echo echo's
	`,
	Run: subEchoRun,
}

func subEchoRun(cmd *cobra.Command, args []string) {
	fmt.Println("From Child", strings.Join(args, " "))
}

func init() {
	cmdEcho.AddCommand(subEcho)
}
