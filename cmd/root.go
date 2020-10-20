package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	Author string
	Domain string
	Token  string
	Scope  string
}

var (
	cfgFile string
)

var config Config

var rootCmd = &cobra.Command{
	Use: "docbase",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("configFile: %s\nconfig: %#v\n", viper.ConfigFileUsed(), config)
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := os.OpenFile(home+"/.docbase.yaml", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		viper.AddConfigPath(home)
		viper.SetConfigName(".docbase")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
