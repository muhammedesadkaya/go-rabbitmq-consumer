package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/go-rabbitmq-consumer-app/api/userapi"
)

var userApi = &cobra.Command{
	Use:   "user-api",
	Short: "user api",
	Long:  `user api`,
	RunE:  userapi.Init,
}

func init() {
	rootCmd.AddCommand(userApi)
}
