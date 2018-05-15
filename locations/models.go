package locations

import (
	"github.com/opiumated/yellowpages/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	ID      bson.ObjectId `json:"id"`
	City    string        `json:"city"`
	State   string        `json:"state"`
	Zip     string        `json:"zipcode"`
	Country string        `json:"country"`
}

func Add(location Location) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.LOCATIONSCOLLECTION)
	return collection.Insert(location)
}

func (location Location) Exists(city, state string) bool {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var geo Location
	collection := session.DB(mongo.DATABASE).C(mongo.LOCATIONSCOLLECTION)
	err := collection.Find(bson.M{"city": city}).One(&geo)
	if err != nil {
		return false
	}
	return true
}

func RemoveAll() error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.LOCATIONSCOLLECTION)
	_, err := collection.RemoveAll(nil)
	return err
}

func GetAll() ([]Location, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var locations []Location
	collection := session.DB(mongo.DATABASE).C(mongo.LOCATIONSCOLLECTION)
	if err := collection.Find(bson.M{}).All(&locations); err != nil {
		return locations, err
	}
	return locations, nil
}
