package auth

import (
	"gorm.io/gorm"
	"time"
	jwt "github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "User Authentication Service"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("Make the secret key here")

type User struct {
	gorm.Model
	
	ID 			uint 	`gorm:"primaryKey;autoIncrement"`
	Email 		string 	`gorm:"unique"`
	Name		string 
	Password 	string 	
}

type UserPublicData struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserCredentialClaim struct {
	jwt.StandardClaims
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func NewClaim(userData UserPublicData) UserCredentialClaim {
	return UserCredentialClaim{
		StandardClaims : jwt.StandardClaims {
			Issuer : APPLICATION_NAME,
			ExpiresAt : time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		ID : userData.ID,
		Name : userData.Name,
		Email : userData.Email,
	}
}
