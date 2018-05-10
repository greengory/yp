package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/opiumated/yellowpages/message"
	"github.com/opiumated/yellowpages/mongo"
	"github.com/opiumated/yellowpages/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type JwtToken struct {
	Token string `json:"token"`
}

func (u User) Exists(email string) bool {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var user User
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	return collection.Find(bson.M{"email": email}).One(&user) == nil
}

func (u User) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := GetAll()
	if err != nil {
		log.Println("Error retrieving Users")
		message.NewAPIError(&message.APIError{Success: false}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: users}, w, http.StatusOK)
}

func (u User) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Println("Invalid Payload")
		message.NewAPIError(&message.APIError{Success: false, Message: "Invalid Payload"}, w)
		return
	}
	newUser.ID = bson.NewObjectId()
	newUser.IsActive = false
	newUser.HashedPassword, _ = utils.HashPassword(newUser.Password)
	newUser.VerifiedToken = utils.GenerateToken()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	//Check if the email combo already exist
	if u.Exists(newUser.Email) {
		message.NewAPIError(&message.APIError{Message: "User with that e-mail already exists", Status: http.StatusConflict}, w)
		return
	}

	if err := Create(newUser); err != nil {
		log.Println("Error creating a new user")
		log.Println(err)
		message.NewAPIError(&message.APIError{Success: false, Message: "Error while creating the user"}, w)
		return
	}
	log.Println("User successfully created...")
	u.SendConfirmationEmail(newUser)
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: newUser}, w, http.StatusOK)
	return
}

func (u User) Login(w http.ResponseWriter, r *http.Request) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	defer r.Body.Close()

	type Credential struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userCredential Credential
	var user User
	if err := json.NewDecoder(r.Body).Decode(&userCredential); err != nil {
		log.Println("Invalid Payload")
		message.NewAPIError(&message.APIError{Success: false, Message: "Invalid Payload"}, w)
		return
	}
	//Check if the email and password hash matches
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	if err := collection.Find(bson.M{"email": userCredential.Email}).One(&user); err != nil {
		log.Println("Unauthorized: Invalid Email or Password")
		message.NewAPIError(&message.APIError{Status: http.StatusUnauthorized}, w)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(userCredential.Password)); err == nil {
		log.Println("Account exist... Creating token")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":   userCredential.Email, //Don't store password in JWT (it's reversible)
			"expires": time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("thisWillBeMovedToADedicatedStruct"))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	} else {
		log.Println("Password and Hashed password doesn't match")
		message.NewAPIError(&message.APIError{Status: http.StatusUnauthorized}, w)
		return
	}
}

func (u User) RemoveAll(w http.ResponseWriter, r *http.Request) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	_, err := collection.RemoveAll(nil)
	if err != nil {
		message.NewAPIError(&message.APIError{Success: false}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "All user data has been removed"}, w, http.StatusOK)
	return
}
