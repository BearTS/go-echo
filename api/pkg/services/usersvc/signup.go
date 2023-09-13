package usersvc

import (
	"github.com/BearTS/go-echo-template/api/pkg/api"
	"github.com/labstack/echo/v4"
)

func (svc *UserSvcImpl) SignUp(c echo.Context, req api.SignupRequest) error {
	svc.Logger.Info("signup request received")
	return nil
}
