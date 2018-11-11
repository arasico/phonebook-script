package middleware

import (
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
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
		context.Set(req, "id", hex.EncodeToString([]byte(user.Id)))
		context.Set(req, "email", user.Email)
		context.Set(req, "username", user.Username)
		next(w, req)
	})
}
