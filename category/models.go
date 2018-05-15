package category

import (
	"github.com/opiumated/yellowpages/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID       bson.ObjectId `json:"id"`
	Category string        `json:"category"`
	Slug     string        `json:"slug"`
}

type Categories []Category

func Add(category Category) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.CATEGORIESCOLLECTION)
	return collection.Insert(category)
}

func RetrieveAll() (Categories, error) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	var cats Categories

	collection := session.DB(mongo.DATABASE).C(mongo.CATEGORIESCOLLECTION)
	if err := collection.Find(nil).All(&cats); err != nil {
		return cats, err
	}
	return cats, nil
}

func Remove(categoryId bson.ObjectId) error {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	collection := session.DB(mongo.DATABASE).C(mongo.CATEGORIESCOLLECTION)
	return collection.RemoveId(categoryId)
}
