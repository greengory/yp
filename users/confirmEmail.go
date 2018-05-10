package users

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/matcornic/hermes"
	"github.com/opiumated/shorpie/message"
	"github.com/opiumated/yellowpages/mongo"
	gomail "gopkg.in/gomail.v2"
	"gopkg.in/mgo.v2/bson"
)

func (u User) ConfirmAccount(w http.ResponseWriter, r *http.Request) {
	session := mongo.Get().Session.Copy()
	defer session.Close()
	//users/confirm?token=actualToken
	token := r.URL.Query().Get("token")

	if token == "" {
		log.Println("Token is needed to confirm an account.")
		message.NewAPIError(&message.APIError{Status: http.StatusBadRequest}, w)
		return
	}
	var user User
	collection := session.DB(database).C(mongo.USERCOLLECTION)
	if err := collection.Find(bson.M{"verified_token": token}).One(&user); err != nil {
		log.Println("Token doesn't exist")
		log.Println(err.Error())
		log.Println("Email: ", user.Email, " Token: ", token)
		message.NewAPIError(&message.APIError{Message: "Invalid token"}, w)
		return
	}
	user.VerifiedToken = "" //Reset VerifiedToken
	user.IsActive = true
	if err := Update(user.ID, user); err != nil {
		log.Println("Error Updating Information...")
		message.NewAPIError(&message.APIError{Message: "There was a problem while confirming your account. Please try again later"}, w)
		return
	}
	log.Println("Account Confirmed")
	message.NewAPIResponse(&message.APIResponse{Message: "Account Confirmed"}, w, http.StatusOK)
	return
}

func (u User) SendConfirmationEmail(user User) {
	h := hermes.Hermes{
		Theme: new(hermes.Default),
		Product: hermes.Product{
			Name: "Yellow Pages",
			Link: "http://localhost:3000/",
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: user.Name,
			Intros: []string{
				"Thank you for signing up with Yellow-pages, we are glad to have u onboard",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please confirm your account by clicking on the link below",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "http://localhost:8081/users/confirm?token=" + user.VerifiedToken,
					},
				},
			},
			Outros: []string{
				"If you need any help, please reply this email. We'd love to help",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Println("Error generating email")
		return
	}
	err = ioutil.WriteFile(user.Name+".html", []byte(emailBody), 0644)
	if err != nil {
		log.Println("You need permission to write to this folder..")
	}
	sendEmail(emailBody, user.Email)
}

func sendEmail(emailBody, recipient string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@yellowpages.com")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Email Confirmation")
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, "ofonimeusoro01@gmail.com", "phoenix01")
	if err := d.DialAndSend(m); err != nil {
		log.Println("Cannot connect to SMTP Server")
		return
	}
}
