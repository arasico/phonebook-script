package routes

import (
	"phonebook-script/controllers/Auth"
	"phonebook-script/models"
)

var AuthController = &Auth.Controller{}

var routes = models.Routes{
	models.Route{
		"register",
		"POST",
		"/api/v1/auth/register",
		AuthController.Register,
	},
	//utils.Route{
	//	"TodoIndex",
	//	"GET",
	//	"/todos",
	//	TodoIndex,
	//},
	//utils.Route{
	//	"TodoCreate",
	//	"POST",
	//	"/todos",
	//	TodoCreate,
	//},
	//utils.Route{
	//	"TodoShow",
	//	"GET",
	//	"/todos/{todoId}",
	//	TodoShow,
	//},
}
