package userapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository/event"
	"gitlab.com/go-rabbitmq-consumer-app/internal/user"
	"gitlab.com/go-rabbitmq-consumer-app/internal/user/docs"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/config"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/mongo"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/rabbitmq"
	"time"
)

func Init(cmd *cobra.Command, args []string) error {
	docs.Init()
	router := gin.New()

	var config config.ApplicationConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//region postgres example
	//type User struct {
	//	feature_id   string
	//	feature_name string
	//}
	//var postgreClient = postgresql.NewClient(config, "information_schema")
	//
	//var db *sql.DB
	//var esad error
	//if db, esad = postgreClient.OpenConnection(); esad != nil {
	//	fmt.Println(esad)
	//}
	//
	//defer postgreClient.CloseConnection(db)
	//
	//users := User{}
	//
	//err := db.QueryRow("SELECT feature_id, feature_name FROM information_schema.sql_features WHERE feature_id = $1", "B011").Scan(&users.feature_id, &users.feature_name)
	//if err != nil {
	//	log.Fatal("Failed to execute query: ", err)
	//}
	//endregion

	var mongoDBClient, err = mongo.NewClient(config.MongoDB.Address, 10*time.Second)
	if err != nil {
		return err
	}

	var messageBus = rabbitmq.NewRabbitMqClient(
		config.RabbitMQ.Host,
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		"",
		rabbitmq.RetryCount(0),
		rabbitmq.PrefetchCount(10))

	messageBus.AddPublisher("Test.Events.V1.User.Insert", rabbitmq.Topic, event.UserInsertEvent{})

	var userRepository = repository.NewUserRepository(mongoDBClient, "user")

	user.GetUsers(router, userRepository)
	user.InsertUser(router, messageBus)

	router.Run(":1453")

	return nil
}
