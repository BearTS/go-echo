package usersvc

import (
	"context"

	"github.com/BearTS/go-echo/pkg/logger"
	"github.com/BearTS/go-echo/pkg/tables"
)

type UserSvcImpl struct {
	DB            UserDb
	logger        logger.Logger
	messageBroker MessageBroker
}

type UserDb interface {
	CreateUser(user *tables.Users) error
}

type MessageBroker interface {
	Publish(ctx context.Context, exchange, routingKey string, body []byte) error
}

func Handler(userDb UserDb, logger logger.Logger, messageBroker MessageBroker) *UserSvcImpl {
	return &UserSvcImpl{
		DB:            userDb,
		logger:        logger,
		messageBroker: messageBroker,
	}
}
