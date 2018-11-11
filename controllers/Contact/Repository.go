package Contact

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

//Repository ...
type Repository struct{}

// DBNAME the name of the DB instance
var DBNAME = os.Getenv("MONGO_DB_NAME")

// COLLECTION is the name of the collection in DB
var COLLECTION = "contacts"

func (r Repository) saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader, userId string) string {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return ""
	}
	os.MkdirAll("./files/contact/"+userId+"/", os.ModePerm)
	photo := fmt.Sprint(userId, time.Now().Unix(), handle.Filename)
	ioutil.WriteFile("./files/contact/"+userId+"/"+photo, data, 0666)
	return photo
}

func (r Repository) insertContact(reg *http.Request, photo string) {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	session.DB(DBNAME).C(COLLECTION).Insert(bson.M{"name": reg.FormValue("name"), "email": reg.FormValue("email"), "mobile": reg.FormValue("mobile"), "phoneNumber": reg.FormValue("phoneNumber"), "address": reg.FormValue("address"), "photo": photo, "created_at": time.Now().Unix()})
}
