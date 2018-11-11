package Contact

import (
	"github.com/gorilla/context"
	"net/http"
	"phonebook-script/utils"
)

//Controller ...
type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	result := c.Repository.getContact(context.Get(r, "id").(string))
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
	c.Repository.insertContact(r, photo, context.Get(r, "id").(string))
	utils.Respond(w, "success", http.StatusOK)
	return
}
