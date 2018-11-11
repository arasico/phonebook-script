package middleware

import (
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"phonebook-script/utils"
)

type User struct {
	Id       bson.ObjectId `json:"_id"`
	Email    string        `json:"email"`
	Username string        `json:"username"`
	jwt.StandardClaims
}

// DBNAME the name of the DB instance
var DBNAME = os.Getenv("MONGO_DB_NAME")

// COLLECTION is the name of the collection in DB
var COLLECTION = "users"

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader == "" {
			utils.Respond(w, "An authorization header is required", http.StatusUnauthorized)
			return
		}
		user := User{}
		token, _ := jwt.ParseWithClaims(authorizationHeader, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if token.Valid != true {
			utils.Respond(w, "An authorization header is required", http.StatusUnauthorized)
			return
		}
		if checkLogin(hex.EncodeToString([]byte(user.Id))) == false {
			utils.Respond(w, "plz check your authorization", http.StatusUnauthorized)
			return
		}
		context.Set(req, "id", hex.EncodeToString([]byte(user.Id)))
		context.Set(req, "email", user.Email)
		context.Set(req, "username", user.Username)
		next(w, req)
	})
}

func checkLogin(userId string) bool {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
	count, _ := c.FindId(bson.ObjectIdHex(userId)).Count()
	if count == 0 {
		return false
	}
	return true

}
