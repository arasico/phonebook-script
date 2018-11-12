package routes

import (
	"phonebook-script/controllers/Auth"
	"phonebook-script/controllers/Contact"
	"phonebook-script/middleware"
	"phonebook-script/models"
)

var AuthController = &Auth.Controller{}

var ContactController = &Contact.Controller{}

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
	models.Route{
		"storeContact",
		"POST",
		"/api/v1/contact",
		middleware.AuthenticationMiddleware(ContactController.Store),
	},
	models.Route{
		"indexContact",
		"Get",
		"/api/v1/contact",
		middleware.AuthenticationMiddleware(ContactController.Index),
	},
	models.Route{
		"indexContactOne",
		"GET",
		"/api/v1/contact/{id}",
		middleware.AuthenticationMiddleware(ContactController.Show),
	},
	models.Route{
		"updateContact",
		"POST",
		"/api/v1/contact/update/{id}",
		middleware.AuthenticationMiddleware(ContactController.Update),
	},
	models.Route{
		"deleteContact",
		"DELETE",
		"/api/v1/contact/{id}",
		middleware.AuthenticationMiddleware(ContactController.Destroy),
	},
}
