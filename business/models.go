package business

import (
	"errors"
	"time"

	"github.com/opiumated/yellowpages/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zipcode"`
	Country string `json:"country"`
}

type Business struct {
	ID               bson.ObjectId `json:"id"`
	Name             string        `json:"name"`
	Location         Location      `json:"location"`
	OwnersName       string        `json:"owner"`
	OwnersPhone      string        `json:"owner_phone"`
	BusinessEmail    string        `json:"email"`
	Type             string        `json:"business_type"`
	IsVerified       bool          `json:"is_verified"`
	Website          string        `json:"website"`
	BusinessCategory string        `json:"category"`
	Description      string        `json:"description"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

func New(business Business) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	if business.AlreadyExist(business.Name) {
		return errors.New("Business with that name already exists")
	}
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS)
	return collection.Insert(business)
}

func (b Business) AlreadyExist(businessName string) bool {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var biz Business
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS)
	if err := collection.Find(bson.M{"name": businessName}).One(&biz); err != nil {
		return false
	}
	return true
}

func SearchBusiness(query interface{}) ([]Business, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var results []Business
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	if err := collection.Find(query).Sort().Limit(20).All(&results); err != nil {
		return results, err
	}
	return results, nil
}
