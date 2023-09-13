package usersvc

import (
	"fmt"

	"github.com/BearTS/go-echo/api/pkg/api"
	"github.com/labstack/echo/v4"
)

func (svc *UserSvcImpl) SignUp(c echo.Context, req api.SignupRequest) error {

	// print the email and password
	if req.Email != nil {
		fmt.Println(*req.Email)
	}

	if req.Password != nil {
		fmt.Println(*req.Password)
	}

	return nil
}
