package auth

import (
	"fmt"
	"crypto/sha256"
)

func hashSHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	encrypted := hash.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}

type any interface{}
type JSONResponse struct {
	Message string 		`json:"message"`
	Data 	any 		`json:"data"`
}