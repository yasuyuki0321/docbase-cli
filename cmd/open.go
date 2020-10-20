package cmd

import (
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open DocBase top page",

	Run: func(cmd *cobra.Command, args []string) {
		url := "https://iret.docbase.io/"
		OpenUrl(url)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
