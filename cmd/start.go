package cmd

import (
	"fmt"
	"os"
	"toucan/internal/server"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the app",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.StartHTTP("0.0.0.0", 4444)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
