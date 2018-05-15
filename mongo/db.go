package mongo

import (
	"log"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

var (
	mongo Mongo
)

const (
	DATABASE             = "yellowpages_test"
	USERCOLLECTION       = "users"
	LOCATIONSCOLLECTION  = "locations"
	RESTAURANTCOLLECTION = "restaurants"
	CATEGORIESCOLLECTION = "categories"
	BUSINESS             = "business"
	BUSINESS_TYPE        = "business_types"
)

type Mongo struct {
	Session *mgo.Session
}

func Init() {
	mongo = Mongo{}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Database: viper.Get("database.name").(string),
	})
	if err != nil {
		log.Fatal("Error Connecting to database...")
	}
	log.Println("Connected to database.")
	mongo.Session = session
}

func Get() *Mongo {
	return &mongo
}
