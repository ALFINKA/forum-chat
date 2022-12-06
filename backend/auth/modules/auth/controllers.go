package auth

import (
	"fmt"
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	jwt "github.com/golang-jwt/jwt/v4"
)

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("database").(*gorm.DB) // Load database from context
	if r.Method != "POST" {
		http.Error(w, "HTTP Method not accepted", http.StatusMethodNotAllowed)
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	var user User
	res := db.Model(&User{}).Where("email = ?", email).First(&user)

	if res.Error != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return		
	}

	if user.Password != hashSHA256(password) {
		http.Error(w, "Password is invalid", http.StatusBadRequest)
		return
	}

	userdata := UserPublicData{
		ID : user.ID,
		Name : user.Name,
		Email : user.Email,
	}

	claims := NewClaim(userdata)
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResp, err := json.Marshal(JSONResponse{
		Message : "New Token Generated",
		Data : signedToken,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("database").(*gorm.DB) // Load database from context

	switch r.Method {
	case "POST":
		user := &User{
			Name : r.PostFormValue("name"),
			Email	 : r.PostFormValue("email"),
			Password : hashSHA256(r.PostFormValue("password")),
		}
		
		result := db.Create(user)

		if result.Error != nil {
			http.Error(w, "Failed to create new user", http.StatusBadRequest)
			return
		}

		jsonResp, err := json.Marshal(JSONResponse{
			Message	 : "New user has been registered",
			Data	 : "",
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		w.Write(jsonResp)
		return
		
	case "GET":
		var users []UserPublicData
		db.Model(&User{}).Select([]string{"ID", "Name", "Email"}).Find(&users)

		jsonResp, err := json.Marshal(JSONResponse{
			Message : "Retrieved all User",
			Data  : users,
		})
		if(err != nil) {
			http.Error(w, "Something went wrong", http.StatusMethodNotAllowed)
			return
		}
		w.Write(jsonResp)
		return
	default:
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}
}

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	badResponse := JSONResponse{}
	if r.Method != "POST" {
		badResponse.Message = "HTTP method is not accepted"
		jsonResp, _ := json.Marshal(badResponse)
		http.Error(w, string(jsonResp), http.StatusMethodNotAllowed)
		return
	}
	
	err := r.ParseForm()
	if err != nil {
		badResponse.Message = err.Error()
		jsonResp, _ := json.Marshal(badResponse)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	tokenString := r.PostFormValue("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		badResponse.Message = err.Error()
		jsonResp, _ := json.Marshal(badResponse)
		http.Error(w, string(jsonResp), http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		badResponse.Message = err.Error()
		jsonResp, _ := json.Marshal(badResponse)
		http.Error(w, string(jsonResp), http.StatusBadRequest)
		return
	}

	jsonResp, _ := json.Marshal(JSONResponse{
		Message : "Token verified",
		Data : claims,
	})

	w.Write(jsonResp)
	return
}