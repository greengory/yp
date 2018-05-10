package utils

import (
	"log"
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	// HashCost Cost for Bcrypt
	HashCost    = 14
	MaxByteSize = 32
)

//WriteJSON Writes a JSON response
func WriteJSON(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

//HashPassword Hashes a password string
//TODO: Add salt while hashing password
func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return hash, err
}

//CheckPassword Compares password and hash for equality
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken() string {
	byteGenerated, err := uuid.GenerateUUID()
	if err != nil {
		log.Println("Error Generating random bytes")
		return ""
	}
	return byteGenerated
}
