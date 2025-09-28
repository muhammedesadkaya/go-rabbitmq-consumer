package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository/entity"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository/event"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/rabbitmq"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/response"
	"golang.org/x/net/context"
)

// GetUsers User godoc
// @Summary Get Users
// @Description Get Users
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Tags User
// @Router /v1/users [get]
func GetUsers(r *gin.Engine, userRepository repository.UserRepository) {
	r.GET("/v1/users", func(c *gin.Context) {

		ctx := context.Background()

		var (
			users []*entity.User
			err   error
		)

		if users, err = userRepository.GetAll(ctx); err != nil {
			c.JSON(response.BAD_REQUEST, nil)
		}

		c.JSON(response.OK, users)
	})
}

// InsertUser User godoc
// @Summary Insert User
// @Description Insert User
// @Param request body UserRequest true "request"
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Tags User
// @Router /v1/user [post]
func InsertUser(r *gin.Engine, messageBus rabbitmq.Client) {
	r.POST("/v1/user", func(c *gin.Context) {
		ctx := context.Background()

		var (
			user *entity.User
			err  error
		)

		if err = c.BindJSON(&user); err != nil {
			c.JSON(response.BAD_REQUEST, err)
		}

		message := event.UserInsertEvent{
			FullName:    user.FullName,
			Identity:    user.Identity,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Gender:      user.Gender,
		}

		if err = messageBus.Publish(ctx, "*", message); err != nil {
			fmt.Println("insertUserPublishError %r", err)
		}

		c.JSON(response.OK, nil)
	})
}
