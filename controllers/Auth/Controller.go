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
		utils.Respond(w, "email address is required", http.StatusUnauthorized)
		return
	}
	if len(r.FormValue("password")) < 6 {
		utils.Respond(w, "password is required", http.StatusUnauthorized)
		return
	}
	checkLogin, check := c.Repository.checkLogin(r.FormValue("email"), r.FormValue("password"))
	if check == false {
		utils.Respond(w, "email and password is wrong", http.StatusUnauthorized)
		return
	}
	utils.Respond(w, checkLogin, http.StatusOK)
	return
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.FormValue("email"), "@") {
		utils.Respond(w, "email address is required", http.StatusUnauthorized)
		return
	}
	if len(r.FormValue("password")) < 6 {
		utils.Respond(w, "password is required", http.StatusUnauthorized)
		return
	}
	if r.FormValue("username") == "" {
		utils.Respond(w, "username is required", http.StatusUnauthorized)
		return
	}
	err := c.Repository.checkEmail(r.FormValue("email"))
	if err == false {
		utils.Respond(w, "email address is exists", http.StatusUnauthorized)
		return
	}
	err = c.Repository.checkUsername(r.FormValue("username"))
	if err == false {
		utils.Respond(w, "username is exists", http.StatusUnauthorized)
		return
	}
	result, _ := c.Repository.insertUser(r.FormValue("email"), r.FormValue("password"), r.FormValue("username"))
	utils.Respond(w, result, http.StatusOK)
	return
}
