package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getInitParams(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(question)
	scanner.Scan()

	return scanner.Text()
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file",

	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("token", getInitParams("Enter DocBase Access Token: "))
		viper.Set("author", getInitParams("Enter DocBase Account ID: "))
		viper.Set("domain", getInitParams("Enter Domain: "))
		viper.Set("scope", "group")

		viper.WriteConfigAs(viper.ConfigFileUsed())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
