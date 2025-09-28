package user

import (
	"encoding/json"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository/entity"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/config"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/rabbitmq"
	"golang.org/x/net/context"
)

type UserInsertConsumer struct {
	config         config.ApplicationConfig
	messageBus     rabbitmq.Client
	userRepository repository.UserRepository
}

func NewUserInsertConsumer(config config.ApplicationConfig, messageBus rabbitmq.Client, userRepository repository.UserRepository) *UserInsertConsumer {

	return &UserInsertConsumer{
		config:         config,
		messageBus:     messageBus,
		userRepository: userRepository,
	}
}

func (self *UserInsertConsumer) Construct() {
	self.messageBus.AddConsumer("In.User.Insert").
		SubscriberExchange("*", rabbitmq.Topic, "Test.Events.V1.User.Insert").
		HandleConsumer(self.userInsertConsumer())
}

func (self *UserInsertConsumer) userInsertConsumer() func(message rabbitmq.Message) error {

	return func(message rabbitmq.Message) error {

		ctx := context.Background()

		var (
			user *entity.User
			err  error
		)

		if err = json.Unmarshal(message.Payload, &user); err != nil {
			return err
		}

		newUser := entity.NewUser(user.FullName, user.Identity, user.PhoneNumber, user.Email, user.Gender)

		if err = self.userRepository.Insert(ctx, newUser); err != nil {
			//TODO: Logger yazılacak.
		}

		return nil
	}
}
