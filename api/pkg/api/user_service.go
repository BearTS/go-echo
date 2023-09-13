package api

import "github.com/labstack/echo/v4"

type UserService interface {
	SignUp(c echo.Context, req SignupRequest) error
}

// SignUp - Signup
// (POST /user/signup)
func (svc *Service) Signup(c echo.Context) error {
	svc.logger.Info("signup request received")
	return nil
}
