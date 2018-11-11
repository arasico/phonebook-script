package routes

import (
	"phonebook-script/controllers/Auth"
	"phonebook-script/models"
)

var AuthController = &Auth.Controller{}

var routes = models.Routes{
	models.Route{
		"login",
		"Post",
		"/api/v1/auth/login",
		AuthController.Login,
	},
	models.Route{
		"register",
		"POST",
		"/api/v1/auth/register",
		AuthController.Register,
	},
}
