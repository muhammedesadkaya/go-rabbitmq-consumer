package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "go-rabbitmq-consumer-app",
	Short: "go-rabbitmq-consumer-app",
	Long:  ``,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.go-rabbitmq-consumer-app.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func initConfig() {

	if configFile == "" {
		os.Exit(1)
	}

	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
