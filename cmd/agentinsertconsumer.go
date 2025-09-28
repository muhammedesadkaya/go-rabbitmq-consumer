package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository"
	"gitlab.com/go-rabbitmq-consumer-app/internal/user"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/config"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/mongo"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/rabbitmq"
	"time"
)

var userInsertConsumer = &cobra.Command{
	Use:   "userinsertconsumer",
	Short: "user insert consumer",
	Long:  `user insert consumer`,
	RunE:  runUserInsertConsumer,
}

func init() {
	rootCmd.AddCommand(userInsertConsumer)
}

func runUserInsertConsumer(cmd *cobra.Command, args []string) error {

	var config config.ApplicationConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	var messageBus = rabbitmq.NewRabbitMqClient(
		config.RabbitMQ.Host,
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		"",
		rabbitmq.RetryCount(0),
		rabbitmq.PrefetchCount(10))

	var mongoDBClient, err = mongo.NewClient(config.MongoDB.Address, 10*time.Second)

	if err != nil {
		return err
	}

	var userRepository = repository.NewUserRepository(mongoDBClient, "user")

	var consumer = user.NewUserInsertConsumer(config, messageBus, userRepository)

	consumer.Construct()

	return messageBus.RunConsumers()
}
