package Auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

//Repository ...
type Repository struct{}

// DBNAME the name of the DB instance
var DBNAME = os.Getenv("MONGO_DB_NAME")

// COLLECTION is the name of the collection in DB
var COLLECTION = "users"

func (r Repository) checkEmail(email string) bool {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
	count, _ := c.Find(bson.M{"email": email}).Count()
	if count == 0 {
		return true
	}
	return false
}
func (r Repository) checkUsername(username string) bool {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
	count, _ := c.Find(bson.M{"username": username}).Count()
	if count == 0 {
		return true
	}
	return false
}

func (r Repository) insertUser(email string, password string, username string) (interface{}, bool) {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	i := bson.NewObjectId()
	session.DB(DBNAME).C(COLLECTION).Insert(bson.M{"_id": i, "email": email, "username": username, "password": string(hashedPassword)})
	user := make(map[string]interface{})
	if err != nil {
		log.Fatal(err)
		return user, false
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      i,
		"email":    email,
		"username": username,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}
	user["_id"] = i
	user["email"] = email
	user["username"] = username
	user["authorization"] = tokenString
	fmt.Println("Added New user => ", email)
	return user, true
}
