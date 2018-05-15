package category

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/metal3d/go-slugify"

	"gopkg.in/mgo.v2/bson"

	"github.com/opiumated/yellowpages/message"
)

func (c Category) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var category Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		log.Println("Invalid Payload")
		message.NewAPIError(&message.APIError{Success: false, Message: "Invalid Payload"}, w)
		return
	}
	category.ID = bson.NewObjectId()
	category.Slug = slugify.Marshal(category.Category)

	if err := Add(category); err != nil {
		log.Println("Error Adding Category")
		log.Println(err)
		message.NewAPIError(&message.APIError{Success: false, Message: "Error Adding Category"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "Category added"}, w, http.StatusOK)
}

func (c Category) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	allCategories, err := RetrieveAll()
	if err != nil {
		log.Println("Error Retrieving Categories")
		message.NewAPIError(&message.APIError{Success: false, Message: "Error Retrieving Categories"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: allCategories}, w, http.StatusOK)
}
