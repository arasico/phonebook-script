package Contact

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"phonebook-script/utils"
)

//Controller ...
type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	result := c.Repository.getAll(context.Get(r, "id").(string))
	utils.Respond(w, result, http.StatusOK)
	return
}
func (c *Controller) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := c.Repository.getOne(context.Get(r, "id").(string), id)
	utils.Respond(w, result, http.StatusOK)
	return
}

func (c *Controller) Store(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("name") == "" {
		utils.Respond(w, "name is required", http.StatusBadRequest)
		return
	}
	file, handle, err := r.FormFile("file")
	photo := ""
	if err == nil {
		defer file.Close()
		mimeType := handle.Header.Get("Content-Type")
		switch mimeType {
		case "image/jpeg":
			photo = c.Repository.saveFile(w, file, handle, context.Get(r, "id").(string))
		case "image/png":
			photo = c.Repository.saveFile(w, file, handle, context.Get(r, "id").(string))
		default:
			utils.Respond(w, "The format file is not valid.", http.StatusBadRequest)
			return
		}
	}
	c.Repository.insert(r, photo, context.Get(r, "id").(string))
	utils.Respond(w, "success", http.StatusOK)
	return
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if r.FormValue("name") == "" {
		utils.Respond(w, "name is required", http.StatusBadRequest)
		return
	}
	file, handle, err := r.FormFile("file")
	photo := ""
	if err == nil {
		defer file.Close()
		mimeType := handle.Header.Get("Content-Type")
		switch mimeType {
		case "image/jpeg":
			photo = c.Repository.saveFile(w, file, handle, context.Get(r, "id").(string))
		case "image/png":
			photo = c.Repository.saveFile(w, file, handle, context.Get(r, "id").(string))
		default:
			utils.Respond(w, "The format file is not valid.", http.StatusBadRequest)
			return
		}
	}
	c.Repository.update(r, photo, context.Get(r, "id").(string), id)
	utils.Respond(w, "success", http.StatusOK)
	return
}

func (c *Controller) Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := c.Repository.delete(context.Get(r, "id").(string), id)
	if !result {
		utils.Respond(w, "plz check your id", http.StatusBadRequest)
		return
	}
	utils.Respond(w, "success", http.StatusOK)
	return
}
