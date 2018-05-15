package locations

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/opiumated/yellowpages/message"
)

func (loc Location) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		log.Println("Invalid Payload")
		message.NewAPIError(&message.APIError{Success: false, Message: "Invalid Payload"}, w)
		return
	}
	location.ID = bson.NewObjectId()
	if loc.Exists(location.City, location.State) {
		log.Println("Location already exists")
		message.NewAPIError(&message.APIError{Success: false, Message: "location already exists"}, w)
		return
	}
	if err := Add(location); err != nil {
		message.NewAPIError(&message.APIError{Success: false, Message: "Cannot Create Location"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: location}, w, http.StatusOK)
}

func (loc Location) GetAll(w http.ResponseWriter, r *http.Request) {
	locations, err := GetAll()
	if err != nil {
		message.NewAPIError(&message.APIError{Success: false, Message: "Error retrieving locations"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: locations}, w, http.StatusOK)
}
