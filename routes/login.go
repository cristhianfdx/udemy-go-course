package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cristhianforerod/udemy-go-course/db"
	"github.com/cristhianforerod/udemy-go-course/jwt"
	"github.com/cristhianforerod/udemy-go-course/models"
)

/*Login router*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid User or email: "+err.Error(), 417)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 417)
		return
	}

	document, isExists := db.TryLogin(user.Email, user.Password)
	if !isExists {
		http.Error(w, "Invalid User or email", 417)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error to generate Session Token: "+err.Error(), 417)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
