package usersvc

import (
	"github.com/BearTS/go-echo/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserSvcImpl struct {
	gormDB *gorm.DB
	Logger logger.Logger
}

type Interface interface {
	SignUp(c echo.Context) error
}

func Handler(gormDB *gorm.DB) *UserSvcImpl {
	return &UserSvcImpl{
		gormDB: gormDB,
	}
}
