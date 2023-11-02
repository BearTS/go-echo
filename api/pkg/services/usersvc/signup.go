package usersvc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/BearTS/go-echo/api/pkg/api"
	"github.com/BearTS/go-echo/pkg/tables"
)

func (svc *UserSvcImpl) SignUp(c context.Context, req api.SignupRequest) (api.SignupResponse, int, error) {

	var message string

	// print the email and password
	if req.Email != nil {
		message = "Email is required"
		return api.SignupResponse{
			Message: &message,
		}, http.StatusBadRequest, nil
	}

	if req.Password != nil {
		message = "Password is required"
		return api.SignupResponse{
			Message: &message,
		}, http.StatusBadRequest, nil
	}

	// Create a new user
	var user tables.Users
	user.Email = string(*req.Email)
	// Hash the password


	if err := svc.DB.CreateUser(&user); err != nil {
		message = "Failed to create user"
		return api.SignupResponse{
			Message: &message,
		}, http.StatusInternalServerError, err
	}

	message = "Signup successful. Please check your email for the verification link."
	var signupResponse api.SignupResponse
	signupResponse.Message = &message

	var msg Message
	msg.From = "anujpflash@gmail.com"
	msg.To = []string{string(*req.Email)}
	msg.Subject = "Signup successful"
	msg.Body = "Please check your email for the verification link."
	msg.Type = "text"

	exchange := "" // Use an empty exchange for direct exchange (default)
	routingKey := "mail"

	// Publish the message to the queue
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message to JSON: %v", err)
	}
	err = svc.messageBroker.Publish(c, exchange, routingKey, body)
	if err != nil {
		return signupResponse, http.StatusInternalServerError, err
	}

	return signupResponse, http.StatusOK, nil
}

type Message struct {
	From         string
	To           []string
	Subject      string
	Body         string
	TemplateName string
	Data         interface{}
	Type         string // template or text or html
}
