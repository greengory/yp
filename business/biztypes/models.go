package biztypes

import (
	"errors"

	"github.com/opiumated/yellowpages/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Types struct {
	Name string `json:"business_type"`
}

func Add(name string) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	return collection.Insert(bson.M{"name": name})
}

func Remove(businessType Types) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	return collection.Remove(businessType)
}

func Update(name string) (Types, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var bizType Types
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	if err := collection.Find(bson.M{"business_type": name}).One(&bizType); err != nil {
		return bizType, errors.New("Business Type doesn't exist")
	}
	return bizType, nil
}

func FindByName(name string) (Types, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var bizType Types
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	err := collection.Find(bson.M{"name": name}).One(&bizType)
	if err != nil {
		return bizType, err
	}
	return bizType, nil
}

func Exists(name string) bool {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var bizType Types
	collection := session.DB(mongo.DATABASE).C(mongo.BUSINESS_TYPE)
	if err := collection.Find(bson.M{"name": name}).One(&bizType); err != nil {
		return false
	}
	return true
}
