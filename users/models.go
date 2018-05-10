package users

import (
	"errors"
	"time"

	"github.com/opiumated/yellowpages/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	database = "yellowpages_test"
)

type User struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Password       string        `json:"-" bson:"-"`
	HashedPassword []byte        `json:"hash_password, omitempty"`
	VerifiedToken  string        `json:"verified_token" bson:"verified_token"`
	IsAdmin        bool          `json:"is_admin"`
	IsActive       bool          `json:"is_active"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

func Create(user User) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	return collection.Insert(user)
}

func GetAll() ([]User, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var users []User
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	if err := collection.Find(bson.M{}).All(&users); err != nil {
		return users, err
	}
	return users, nil
}

func Update(userId interface{}, newUser User) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	return collection.UpdateId(userId, newUser)
}

func DisableUser(userId string) error {
	if !bson.IsObjectIdHex(userId) {
		return errors.New("Invalid Format for User Id")
	}
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	return collection.UpdateId(userId, bson.M{"is_active": false})
}
