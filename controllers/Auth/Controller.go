package Auth

import (
	"net/http"
	"phonebook-script/utils"
	"strings"
)

//Controller ...
type Controller struct {
	Repository Repository
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.FormValue("email"), "@") {
		utils.Respond(w, "email address is required", 401)
		return
	}
	if len(r.FormValue("password")) < 6 {
		utils.Respond(w, "password is required", 401)
		return
	}
	checkLogin := c.Repository.checkLogin(r.FormValue("email"), r.FormValue("password"))
	utils.Respond(w, checkLogin, 200)
	return
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.FormValue("email"), "@") {
		utils.Respond(w, "email address is required", 401)
		return
	}
	if len(r.FormValue("password")) < 6 {
		utils.Respond(w, "password is required", 401)
		return
	}
	if r.FormValue("username") == "" {
		utils.Respond(w, "username is required", 401)
		return
	}
	err := c.Repository.checkEmail(r.FormValue("email"))
	if err == false {
		utils.Respond(w, "email address is exists", 401)
		return
	}
	err = c.Repository.checkUsername(r.FormValue("username"))
	if err == false {
		utils.Respond(w, "username is exists", 401)
		return
	}
	result, _ := c.Repository.insertUser(r.FormValue("email"), r.FormValue("password"), r.FormValue("username"))
	utils.Respond(w, result, 200)
	return
}
