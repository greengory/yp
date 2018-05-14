package business

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/opiumated/yellowpages/business/biztypes"
	"github.com/opiumated/yellowpages/message"
)

func (b Business) Create(w http.ResponseWriter, r *http.Request) {
	var business Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		log.Println("Invalid Payload")
		message.NewAPIError(&message.APIError{Message: "Invalid Payload"}, w)
		return
	}
	if !biztypes.Exists(business.Type) {
		biztypes.Add(business.Type)
	}
	business.IsVerified = false
	business.CreatedAt = time.Now()
	business.UpdatedAt = time.Now()

	if err := New(business); err != nil {
		log.Println("There was an error creating a new business")
		message.NewAPIError(&message.APIError{Message: "There was an error creating business"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true}, w, http.StatusOK)
	return

}
