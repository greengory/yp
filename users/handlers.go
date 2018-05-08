package users

import (
	"log"
	"net/http"

	"github.com/opiumated/yellowpages/message"
)

func (u User) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := GetAll()
	if err != nil {
		log.Println("Error retrieving Users")
		message.NewAPIError(&message.APIError{Success: false}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Data: users}, w, http.StatusOK)
}
