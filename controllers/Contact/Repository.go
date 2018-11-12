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
	os.MkdirAll("/files/contact/"+userId+"/", os.ModePerm)
	photo := fmt.Sprint("/files/contact/", userId, "/", userId, time.Now().UnixNano(), handle.Filename)
	ioutil.WriteFile(photo, data, 0666)
	return photo
}

func (r Repository) insert(reg *http.Request, photo string, userId string) {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	session.DB(DBNAME).C(COLLECTION + userId).Insert(bson.M{"name": reg.FormValue("name"), "email": reg.FormValue("email"), "mobile": reg.FormValue("mobile"), "phoneNumber": reg.FormValue("phoneNumber"), "address": reg.FormValue("address"), "photo": photo, "created_at": time.Now().UnixNano()})
}

func (r Repository) getAll(userId string) interface{} {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	var result []interface{}
	session.DB(DBNAME).C(COLLECTION + userId).Find(nil).All(&result)
	return result
}
func (r Repository) getOne(userId string, id string) interface{} {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	var result interface{}
	session.DB(DBNAME).C(COLLECTION + userId).FindId(bson.ObjectIdHex(id)).One(&result)
	return result
}

func (r Repository) update(reg *http.Request, photo string, userId string, id string) {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	if photo == "" {
		var result struct {
			Photo string `json:"photo"`
		}
		session.DB(DBNAME).C(COLLECTION + userId).FindId(bson.ObjectIdHex(id)).One(&result)
		photo = result.Photo
	}
	session.DB(DBNAME).C(COLLECTION+userId).UpdateId(bson.ObjectIdHex(id), bson.M{"name": reg.FormValue("name"), "email": reg.FormValue("email"), "mobile": reg.FormValue("mobile"), "phoneNumber": reg.FormValue("phoneNumber"), "address": reg.FormValue("address"), "photo": photo, "update_at": time.Now().UnixNano()})
}

func (r Repository) delete(userId string, id string) bool {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	err = session.DB(DBNAME).C(COLLECTION + userId).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return false
	}
	return true
}
